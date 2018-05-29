#!/bin/sh
make
scp -i ~/.ssh/vultr -r dist/linux_amd64/* root@207.246.80.69:/root/torrent
