#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.
set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir

echo "running [build-machine]...";
tart run --dir=data:./data --dir=pftest:../.. pftest-build-machine

