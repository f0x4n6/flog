# flog
[![Go Reference](https://pkg.go.dev/badge/github.com/f0x4n6/flog.svg)](https://pkg.go.dev/github.com/f0x4n6/flog)
[![Go Report Card](https://goreportcard.com/badge/github.com/f0x4n6/flog?style=flat-square)](https://goreportcard.com/report/github.com/f0x4n6/flog)
[![Release](https://img.shields.io/github/release/f0x4n6/flog.svg?style=flat-square)](https://github.com/f0x4n6/flog/releases/latest)

Log forensic artifacts as JSON in [ECS](https://www.elastic.co/guide/en/ecs/current/index.html) format.

```console
go install github.com/f0x4n6/flog@latest
```

## Usage
```console
$ flog [-pqhv] [-D DIRECTORY] [FILE ...]
```

Available options:

- `-D` Log directory
- `-p` Pretty JSON
- `-q` Quiet mode
- `-h` Show usage
- `-v` Show version

Required system commands:

- [dotnet](https://dotnet.microsoft.com/en-us/download/dotnet/9.0)

> Use `scripts/eztools.sh` to install [Eric Zimmerman's Tools](https://ericzimmerman.github.io/#!index.md).

## Artifacts
Supported artifacts for Windows 7+ systems:

- [System Event Logs](https://forensics.wiki/windows_event_log_%28evt%29/)
- [User JumpLists](https://forensics.wiki/jump_lists/)
- [User ShellBags](https://forensics.wiki/shell_item/)
- [User Browser Histories](https://forensics.wiki/google_chrome/)

## License
Released under the [MIT License](LICENSE.md).