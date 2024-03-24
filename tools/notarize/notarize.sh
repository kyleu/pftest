#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

if [ "$PUBLISH_TEST" != "true" ]
then
  xcrun notarytool submit --apple-id $APPLE_EMAIL --team-id $APPLE_TEAM_ID --password $APPLE_PASSWORD ./build/dist/pftest_0.0.0_darwin_amd64_desktop.dmg
  xcrun notarytool submit --apple-id $APPLE_EMAIL --team-id $APPLE_TEAM_ID --password $APPLE_PASSWORD ./build/dist/pftest_0.0.0_darwin_amd64.zip

  xcrun notarytool submit --apple-id $APPLE_EMAIL --team-id $APPLE_TEAM_ID --password $APPLE_PASSWORD ./build/dist/pftest_0.0.0_darwin_arm64_desktop.dmg
  xcrun notarytool submit --apple-id $APPLE_EMAIL --team-id $APPLE_TEAM_ID --password $APPLE_PASSWORD ./build/dist/pftest_0.0.0_darwin_arm64.zip

  xcrun notarytool submit --apple-id $APPLE_EMAIL --team-id $APPLE_TEAM_ID --password $APPLE_PASSWORD ./build/dist/pftest_0.0.0_darwin_all_desktop.dmg
  xcrun notarytool submit --apple-id $APPLE_EMAIL --team-id $APPLE_TEAM_ID --password $APPLE_PASSWORD ./build/dist/pftest_0.0.0_darwin_all.zip
fi
