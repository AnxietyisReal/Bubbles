#!/bin/bash
# Builds the container image with optional push to the container registry.

user=$1
repo=$2
tag=$3
dir=$4

if [ "$user" == "" ] || [ "$repo" == "" ] || [ "$tag" == "" ] || [ "$dir" == "" ]; then
    echo "Usage: buildContainer.sh <user> <repo> <tag> <dir>"
    exit 1
fi

echo "Building image $user/$repo:$tag from $dir"

docker build -t git.toast-server.net/$user/$repo:$tag -f $dir .
echo "Image has been built"

read -p "Do you want to push the image to the registry? (y/n) " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]
then
    docker push git.toast-server.net/$user/$repo:$tag
    echo "Image has been pushed to the registry (https://git.toast-server.net/$user/-/packages/container/$repo/$tag)"
fi

###EOF###