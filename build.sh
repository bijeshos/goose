#!/bin/bash
echo "performing goose build process"
echo "creating binary for windows:amd64"
env GOOS=windows GOARCH=amd64 go build -o goose.exe
echo "done"
echo "creating binary for linux:amd64"
env GOOS=linux GOARCH=amd64 go build -o goose
echo "done"
echo "finishing build process"