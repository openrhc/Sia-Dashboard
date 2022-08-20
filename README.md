# Sia-Dashboard

A Simple Linux Monitoring Dashboard.

![preview01](https://raw.githubusercontent.com/openrhc/Sia-Dashboard/main/doc/imgs/preview01.jpg)

## Installation

```bash
wget -O /usr/local/bin/sia-dashboard $release_url
chmod +x /usr/local/bin/sia-dashboard
```

## Usage

```bash
sia-dashboard -p 8000
```

## Compilation

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $appname  -trimpath -ldflags "-s -w -X main.Version=$version -buildid=" main.go
```

## Thanks

[Rudolf-Barbu/Ward](https://github.com/Rudolf-Barbu/Ward)
