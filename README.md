# Sia-Dashboard

A Simple Linux Monitoring Dashboard.

![preview01](https://raw.githubusercontent.com/openrhc/Sia-Dashboard/main/doc/imgs/preview01.jpg)

## Installation

```bash
wget $release_url
gzip -d sia-dashboard.gz
chmod +x sia-dashboard
mv sia-dashboard /usr/local/bin/
```

## Usage

```bash
sia-dashboard -h

Usage of sia-dashboard:
  -a string
        address (default "0.0.0.0")
  -p string
        port (default "8000")

Sia-Dashboard openrhc-20220820. Custom go1.19 linux/arm64
A Simple Linux Monitoring Dashboard.
```

## Systemd

```
/etc/systemd/system/sia-dashboard.service
[Unit]
Description=Sia-Dashboard Daemon
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/sia-dashboard
Restart=on-failure
RestartSec=5s
StandardOutput=append:/dev/null
StandardError=append:/dev/null

[Install]
WantedBy=multi-user.target
```

## Compilation

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $appname  -trimpath -ldflags "-s -w -X main.Version=$version -buildid=" main.go
```

## Thanks

[Rudolf-Barbu/Ward](https://github.com/Rudolf-Barbu/Ward)
