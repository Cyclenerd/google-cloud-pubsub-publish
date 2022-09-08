#!/usr/bin/env bash

# Linux
echo "Linux"                                                            && \
GOOS=linux GOARCH=amd64 go build -o pubsub-publish-linux-x86_64 main.go && \
GOOS=linux GOARCH=arm64 go build -o pubsub-publish-linux-arm64  main.go && \

# macOS
echo "macOS"                                                             && \
GOOS=darwin GOARCH=amd64 go build -o pubsub-publish-macos-x86_64 main.go && \
GOOS=darwin GOARCH=arm64 go build -o pubsub-publish-macos-arm64  main.go && \

# Windows
echo "Windows"                                                                  && \
GOOS=windows GOARCH=amd64 go build -o pubsub-publish-windows-x86_64.exe main.go && \
GOOS=windows GOARCH=arm64 go build -o pubsub-publish-windows-arm64.exe  main.go && \

echo "✅ DONE"
