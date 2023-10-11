#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.

## Builds the app, installing all prerequisites

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

./bin/bootstrap.sh
./bin/templates.sh
go mod download

./bin/build/client.sh

make build-release
mkdir -p ./tmp
mv "./build/release/pftest" "./tmp/pftest"
