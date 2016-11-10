// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

const (
	// BANNER is what is printed for help/info output.
	BANNER = "happy - %s\n"
	// VERSION is the binary version.
	VERSION = "0.1.0"
)

var (
	debug     bool
	dryrun    bool
	version   bool
	token     string
	gitlabURL string
)

func init() {
	// parse flags
	flag.StringVar(&token, "token", "", "Gitlab API token")
	flag.StringVar(&gitlabURL, "gitlabURL", "", "Gitlab URL")
	flag.BoolVar(&dryrun, "dry-run", false, "do not change branch settings just print the changes that would occur")
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&debug, "d", false, "run in debug mode")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if version {
		fmt.Printf("%s", VERSION)
		os.Exit(0)
	}

	// set log level
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if token == "" {
		usageAndExit("Gitlab token cannot be empty.", 1)
	}
}

func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}

func main() {
	// On ^C, or SIGTERM handle exit.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for sig := range c {
			logrus.Infof("Received %s, exiting.", sig.String())
			os.Exit(0)
		}
	}()

	client := gitlab.NewClient(nil, token)
	if gitlabURL != "" {
		cfg := &tls.Config{
			InsecureSkipVerify: true,
		}
		http.DefaultClient.Transport = &http.Transport{
			TLSClientConfig: cfg,
		}
		client.SetBaseURL(gitlabURL)
	}
	logrus.Debugf("URL: %s", client.BaseURL())

	// Get the current user
	user, _, err := client.Users.CurrentUser()
	if err != nil {
		logrus.Fatal(err)
	}
	username := user.Username
	logrus.Debugf("Current user: %s", username)

	if err := getProjects(client, username); err != nil {
		logrus.Fatal(err)
	}
}

func getProjects(client *gitlab.Client, username string) error {
	opt := &gitlab.ListProjectsOptions{Search: gitlab.String(username)}
	projects, _, err := client.Projects.ListProjects(opt)
	if err != nil {
		return err
	}
	for _, project := range projects {
		logrus.Debugf("Project: %s %s", project.Name, project.DefaultBranch)
		if err := handleProject(client, project); err != nil {
			logrus.Warn(err)
		}
	}
	return nil
}

func handleProject(client *gitlab.Client, project *gitlab.Project) error {
	branches, _, err := client.Branches.ListBranches(project.ID)
	if err != nil {
		return err
	}
	for _, branch := range branches {
		if branch.Name == "master" {
			if branch.Protected {
				fmt.Printf("[OK] %s:%s is already protected\n", project.Name, branch.Name)
				return nil
			}
			fmt.Printf("[UPDATE] %s:%s will be changed to protected\n", project.Name, branch.Name)
			if dryrun {
				return nil
			}

			if _, _, err := client.Branches.ProtectBranch(project.ID, branch.Name); err != nil {
				return err
			}
		}
	}
	return nil
}
