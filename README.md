CLI utility that shows current OS information, 
like OS name, version, kernel version, 
installed packages, etc.

## Usage

```bash
$ go run /cmd/client/main.go

```

output
```json
{
  "hostname": "AAA-MacBook-Pro.local",
  "ip": "1.1.1.1",
  "os": "macos",
  "osName": "macOS",
  "osVersion": {
    "version": "14.2.1",
    "major": 14,
    "minor": 2,
    "patch": 1
  },
  "kernelVersion": {
    "version": "23.2.0",
    "major": 23,
    "minor": 2,
    "patch": 0
  },
  "arch": "arm64",
  "apps": [
    {
      "name": "Wireless Diagnostics",
      "version": "11.0",
      "major": 11,
      "minor": 0,
      "patch": 0
    },
    {
      "name": "iOS App Installer",
      "version": "1.0",
      "major": 1,
      "minor": 0,
      "patch": 0
    },
    {
      "name": "Finder",
      "version": "14.2",
      "major": 14,
      "minor": 2,
      "patch": 0
    }
    # etc...
  ]
}
```

send POST request to server with flag -u

```bash
$ go run /cmd/server/main.go -u https://mysite.com/post

```