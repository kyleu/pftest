#!/bin/bash

## Builds the Android library and application

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

TGT=$1
[ "$TGT" ] || TGT="0.0.0"

echo "building gomobile for Android..."
mkdir -p build/dist/mobile_android_arm64
GOARCH=arm64 time gomobile bind -o build/dist/mobile_android_arm64/pftest.aar -target=android -androidapi 21 github.com/kyleu/pftest/app/cmd
echo "gomobile for Android completed successfully, building distribution..."
cd "build/dist/mobile_android_arm64"
zip -r "../pftest_${TGT}_android_aar.zip" .

echo "creating Android project..."
cd "$dir/../.."
mkdir -p build/dist/mobile_android_app_arm64
cp -R tools/android/* build/dist/mobile_android_app_arm64

echo "building Android project..."
cd build/dist/mobile_android_app_arm64
rm -rf ./app/libs
mkdir -p ./app/libs
cp ../mobile_android_arm64/pftest.aar ./app/libs/
cp ../mobile_android_arm64/pftest-sources.jar ./app/libs/

gradle assembleDebug
cd app/build/outputs/apk/debug
zip -r "$dir/../../build/dist/pftest_${TGT}_android_apk.zip" .
