#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.

## Starts the notebook in dev mode

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/.."

if [ -f ".env" ]; then
	export "$(cat .env | grep -v "#" | xargs)"
fi

npm run dev -- --port 41010
