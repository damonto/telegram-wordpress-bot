#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -o build/linux-amd64-telegram-wordpress-bot *.go
chmod +x build/linux-amd64-telegram-wordpress-bot

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -trimpath -o build/linux-arm64-telegram-wordpress-bot *.go
chmod +x build/linux-arm64-telegram-wordpress-bot

CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -trimpath -o build/macos-arm64-telegram-wordpress-bot *.go
chmod +x build/macos-arm64-telegram-wordpress-bot
