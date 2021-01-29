#!/bin/bash
env GOOS=windows go build -ldflags "-X main.os=windows -H=windowsgui -s -w" -o ./builds/windows.exe
env GOOS=darwin go build -ldflags "-X main.os=darwin -s -w" -o ./builds/macos
env GOOS=linux go build -ldflags "-X main.os=linux -s -w" -o ./builds/linux