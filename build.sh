#!/bin/bash

# Project name and output directory
PROJECT_NAME="GpioController"
OUTPUT_DIR="bin"

rm go.mod
rm go.sum

go mod init GpioController

go mod tidy

# Create the output directory if it doesn't exist
mkdir -p $OUTPUT_DIR

# # Update fyne and dependencies
# go get -u fyne.io/fyne/v2
# go get -u github.com/go-gl/gl/v3.1/gles2

# # Clean the Go module cache to avoid any old packages
# go clean -modcache

# Build the project with specific tags for Fyne and logging enabled
echo "Building binary..."
# go build -v -tags="gles2 no_native_menus" -o $OUTPUT_DIR/$PROJECT_NAME cmd/main.go
env GOOS=linux GOARCH=arm GOARM=6 go build -o $OUTPUT_DIR/$PROJECT_NAME cmd/main.go

# Check if the build succeeded
if [ $? -eq 0 ]; then
    echo "Build succeeded. The binary is located at $OUTPUT_DIR/$PROJECT_NAME"
else
    echo "Build failed."
fi
