#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o bin/coffeeorders-amd64-linux main.go
./bin/coffeeorders-amd64-linux