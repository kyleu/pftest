#!/bin/bash

## Builds the Docker image and runs it

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

echo "Building [linux amd64]..."
GOOS=linux GOARCH=amd64 make build
mv ./build/debug/pftest .
docker build -t=pftest -f=./tools/release/Dockerfile.release .
rm ./pftest
docker run -it pftest
