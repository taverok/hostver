CLI utility that shows current OS information, 
like OS name, version, kernel version, 
installed packages, etc.

## Usage

```bash
$ go run /cmd/client/main.go

```

output Linux
```json
{
  "hostname": "aaa",
  "ip": "1.1.1.1",
  "os": "linux",
  "osName": "Ubuntu",
  "osVersion": {
    "version": "22.04 LTS",
    "major": 22,
    "minor": 4,
    "patch": 0
  },
  "kernelVersion": {
    "version": "6.6.6-76060606-generic",
    "major": 6,
    "minor": 6,
    "patch": 6
  },
  "arch": "amd64",
  "apps": [
    {
      "name": "accountsservice",
      "version": "22.07.5-2ubuntu1.3pop1~1675994188~22.04~45e75a8",
      "major": 22,
      "minor": 7,
      "patch": 5
    },
    {
      "name": "acl",
      "version": "2.3.1-1",
      "major": 2,
      "minor": 3,
      "patch": 1
    },
    {
      "name": "acpi-support",
      "version": "0.144",
      "major": 0,
      "minor": 144,
      "patch": 0
    }
    // etc...
  ]
}
```

output macOS
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
    // etc...
  ]
}
```

send POST request to server with flag -u

```bash
$ go run /cmd/server/main.go -u https://mysite.com/post

```