#!/bin/bash



pwd

go build logFile.com/log-file-go

[ ! -d "./go-build" ] && mkdir ./go-build
[ ! -f "./go-build/log-file-go/log-file-go" ] && rm -rf ./go-build/log-file-go/log-file-go 

mv ./log-file-go ./go-build

chmod 755 ./go-build/log-file-go

