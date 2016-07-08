#!/bin/bash

VERSION="0.1"

rm -rf bin

GOOS=darwin  GOARCH=amd64 go build -o bin/macos/envexp
GOOS=linux   GOARCH=amd64 go build -o bin/linux/envexp
GOOS=windows GOARCH=amd64 go build -o bin/windows/envexp.exe

tar czf bin/envexp-$VERSION-macos.tar.gz --directory=bin/macos envexp
tar czf bin/envexp-$VERSION-linux.tar.gz --directory=bin/linux envexp
zip     bin/envexp-$VERSION-windows.zip -j bin/windows/envexp.exe
