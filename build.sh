#!/bin/sh
appname="sia-dashboard"

author="openrhc-"

version=""

if [ -f "VERSION" ]; then
    version=`cat VERSION`
fi

if [[ -z $version ]]; then
    if [ -d ".git" ]; then
        version=`git symbolic-ref HEAD | cut -b 12-`-`git rev-parse HEAD`
    else
        version=$author`date +"%Y%m%d"`
    fi
fi

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $appname  -trimpath -ldflags "-s -w -X main.Version=$version -buildid=" main.go
