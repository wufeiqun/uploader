#!/bin/sh
# GOOS=linux GOARCH=amd64 go build -o uploader-linux-amd64 main.go
GOOS=darwin GOARCH=arm64 go build -o uploader main.go