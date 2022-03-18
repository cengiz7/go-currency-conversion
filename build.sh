#!/usr/bin/bash

main="main.go"

###### LINUX binaries
# 64-bit
env GOOS=linux GOARCH=amd64 go build -o bin/linux/app-amd64-linux $main
# 32-bit
env GOOS=linux GOARCH=386 go build -o bin/linux/app-386-linux $main


###### MacOS binaries
# 64-bit
env GOOS=darwin GOARCH=amd64 go build -o bin/macos/app-amd64-darwin $main
# 32-bit
env GOOS=darwin GOARCH=386 go build -o bin/macos/app-386-darwin $main


###### Windows binaries
# 64-bit
env GOOS=windows GOARCH=amd64 go build -o bin/windows/app-amd64.exe $main
# 32-bit
env GOOS=windows GOARCH=386 go build -o bin/windows/app-386.exe $main