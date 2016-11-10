# Happy

[![License Apache 2][badge-license]](LICENSE)

A clone of [Pepper](https://github.com/jessfraz/pepper) for your Gitlab projects.

## Installation

You can download the binaries :

* Architecture i386 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_linux_386) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_openbsd_386) / [windows](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/happy-0.1.0_netbsd_arm) ]


## Usage

```bash
$ happy -h
happy - 0.1.0
  -d    run in debug mode
  -dry-run
        do not change branch settings just print the changes that would occur
  -gitlabURL string
        Gitlab URL (default "https://gitlab.com")
  -token string
        Gitlab API token
  -version
        print version and exit
```

```bash
$ happy -token 2468 -dry-run
[UPDATE] TbcGit:master will be changed to protected
[OK] road-to-modern-devops:master is already protected
[UPDATE] eudyptula:master will be changed to protected
[UPDATE] arcaishon:master will be changed to protected
[OK] Scame:master is already protected
```

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat
