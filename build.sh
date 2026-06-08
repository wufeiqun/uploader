#!/bin/bash
set -e

APP_NAME="uploader"
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(date '+%Y-%m-%d %H:%M:%S')
LDFLAGS="-s -w -X 'main.Version=${VERSION}' -X 'main.BuildTime=${BUILD_TIME}'"

echo "Building ${APP_NAME} ${VERSION} ..."

# Linux amd64（用于服务器部署）
GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o "${APP_NAME}-linux-amd64" main.go
echo "  ✓ ${APP_NAME}-linux-amd64"

# macOS Apple Silicon（本地开发）
GOOS=darwin GOARCH=arm64 go build -ldflags "${LDFLAGS}" -o "${APP_NAME}-darwin-arm64" main.go
echo "  ✓ ${APP_NAME}-darwin-arm64"

echo "Done."
