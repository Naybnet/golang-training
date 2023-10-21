#!/bin/zsh

VERSION=$1
FILENAME=go${VERSION}.linux-amd64.tar.gz

sudo rm -rf /usr/local/go
echo https://dl.google.com/go/${FILENAME}
wget https://dl.google.com/go/${FILENAME}
sudo tar -C /usr/local -xzf ./${FILENAME}

rm -rf ./${FILENAME}
