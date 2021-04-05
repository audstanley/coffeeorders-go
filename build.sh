#!/bin/bash
GOOS=windows GOARCH=amd64 go build -o bin/coffeeorders-amd64.exe main.go
GOOS=windows GOARCH=386 go build -o bin/coffeeorders-386.exe main.go
GOOS=darwin GOARCH=amd64 go build -o bin/coffeeorders-amd64-macos main.go
GOOS=linux GOARCH=amd64 go build -o bin/coffeeorders-amd64-linux main.go