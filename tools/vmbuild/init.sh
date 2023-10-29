#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.
set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir

echo "initializing tart virtual machine...";
tart clone ghcr.io/cirruslabs/macos-sonoma-base:latest pftest-build-machine

tart set pftest-build-machine --cpu 6 --memory 4096
