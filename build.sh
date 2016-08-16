#!/bin/bash

VERSION="v0.3.1"
DOCKER_REGISTRY="docker.io"

rm -rf bin

GOOS=darwin  GOARCH=amd64 go build -o bin/macos/envexp
GOOS=linux   GOARCH=amd64 go build -o bin/linux/envexp
GOOS=windows GOARCH=amd64 go build -o bin/windows/envexp.exe

#tar czf bin/envexp-$VERSION-macos.tar.gz --directory=bin/macos envexp
#tar czf bin/envexp-$VERSION-linux.tar.gz --directory=bin/linux envexp
#zip     bin/envexp-$VERSION-windows.zip -j bin/windows/envexp.exe

mkdir bin/docker
cp Dockerfile bin/docker/Dockerfile
cp bin/linux/envexp bin/docker/envexp

docker build --no-cache -t "$DOCKER_REGISTRY/maxmanuylov/envexp:$VERSION" bin/docker
docker tag "$DOCKER_REGISTRY/maxmanuylov/envexp:$VERSION" "$DOCKER_REGISTRY/maxmanuylov/envexp:latest"

#docker push "$DOCKER_REGISTRY/maxmanuylov/envexp:$VERSION"
#docker push "$DOCKER_REGISTRY/maxmanuylov/envexp:latest"
