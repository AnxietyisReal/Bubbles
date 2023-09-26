#!/bin/bash
# Builds the container image with optional push to the container registry.

tag=$1

if [ "$tag" == "" ]; then
    echo "Usage: buildContainer.sh <tag>"
    exit 1
fi

echo "Building image toast/bubbles:$tag from docker/bot/Dockerfile"

docker build -t git.toast-server.net/toast/bubbles:$tag -f docker/bot/Dockerfile .
echo "Image has been built"

read -p "Do you want to push the image to the registry? (y|n) " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]
then
    docker push git.toast-server.net/toast/bubbles:$tag
    echo "Image has been pushed to the registry (https://git.toast-server.net/toast/-/packages/container/bubbles/$tag)"
    echo "Remoting into the server to pull the image"
    bash -c "./deployContainer.sh"
fi

###EOF###