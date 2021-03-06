#!/bin/sh

# https://download.docker.com/linux/ubuntu/dists/xenial/pool/stable/amd64/docker-ce_17.06.0~ce-0~ubuntu_amd64.deb

DOCKER_VERSION={{ .DockerVersion }}
UBUNTU_RELEASE={{ .ReleaseVersion }}
ARCH={{ .DockerArch }}
OUT_DIR=/tmp
URL="https://download.docker.com/linux/ubuntu/dists/${UBUNTU_RELEASE}/pool/stable/${ARCH}/docker-ce_${DOCKER_VERSION}~ce-0~ubuntu_${ARCH}.deb"

wget -O $OUT_DIR/$(basename $URL) $URL
sudo apt install -y $OUT_DIR/$(basename $URL)
rm $OUT_DIR/$(basename $URL)
