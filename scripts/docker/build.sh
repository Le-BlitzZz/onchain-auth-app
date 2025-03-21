#!/usr/bin/env bash

set -e

if [[ -z $1 ]]; then
  echo "Usage: ${0##*/} [develop|authonchain]" 1>&2
    exit 1
fi

echo "Building image 'authonchain/$1' from docker/$1/Dockerfile...";

docker build \
  --no-cache \
  -t authonchain/$1:1.24.1-alpine3.21 \
  -f docker/$1/Dockerfile .


echo "Done."
