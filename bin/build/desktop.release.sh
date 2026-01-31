#!/bin/bash

## Builds desktop release artifacts for macOS, Linux, and Windows.
##
## Usage:
##   ./bin/build/desktop.release.sh [version] [-y|--yes]
##
## Arguments:
##   version  Version tag for output filenames (default: v0.0.0).
##
## Requires:
##   - Docker
##   - appdmg, codesign, lipo (macOS)
##   - APPLE_SIGNING_IDENTITY set for codesign
##   - curl and zip
##
## Outputs:
##   - build/dist/pftest_<version>_darwin_*_desktop.(dmg|zip)
##   - build/dist/pftest_<version>_linux_amd64_desktop.zip
##   - build/dist/pftest_<version>_windows_amd64_desktop.zip

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

require_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "error: required command '$1' not found${2:+ ($2)}" >&2
    exit 1
  fi
}

require_env() {
  if [[ -z "${!1:-}" ]]; then
    echo "error: required environment variable '$1' is not set" >&2
    exit 1
  fi
}

require_cmd docker "install Docker Desktop from https://www.docker.com/products/docker-desktop/"
require_cmd zip "install zip from your package manager"
require_cmd curl "install curl from your package manager"
require_cmd appdmg "install via npm i -g appdmg"
require_cmd codesign "requires macOS codesign tool"
require_cmd lipo "requires macOS Xcode command line tools"
require_env APPLE_SIGNING_IDENTITY

TGT=""

while [[ $# -gt 0 ]]; do
  case "$1" in
    --) shift; break;;
    -*)
      echo "unknown option: $1" >&2
      exit 1
      ;;
    *)
      if [[ -z "$TGT" ]]; then
        TGT="$1"
        shift
      else
        echo "unexpected argument: $1" >&2
        exit 1
      fi
      ;;
  esac
done

TGT=${TGT:-v0.0.0}

if command -v retry &> /dev/null
then
  retry -t 4 -- docker build -f tools/desktop/Dockerfile.desktop -t pftest .
else
  docker build -f tools/desktop/Dockerfile.desktop -t pftest .
fi


rm -rf tmp/release
mkdir -p tmp/release

cd "tmp/release"

id=$(docker create pftest)
docker cp "$id":/dist - > ./desktop.tar
docker rm -v "$id"
tar -xvf "desktop.tar"
rm "desktop.tar"

mv dist/darwin_amd64/pftest "pftest.darwin"
mv dist/darwin_arm64/pftest "pftest.darwin.arm64"
mv dist/linux_amd64/pftest "pftest"
mv dist/windows_amd64/pftest "pftest.exe"
rm -rf "dist"

# darwin amd64
cp -R "../../tools/desktop/template" .

mkdir -p "./Test Project.app/Contents/Resources"
mkdir -p "./Test Project.app/Contents/MacOS"

cp -R "./template/darwin/Info.plist" "./Test Project.app/Contents/Info.plist"
cp -R "./template/darwin/icons.icns" "./Test Project.app/Contents/Resources/icons.icns"

cp "pftest.darwin" "./Test Project.app/Contents/MacOS/pftest"

echo "signing amd64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s "${APPLE_SIGNING_IDENTITY}" "./Test Project.app/Contents/MacOS/pftest"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s "${APPLE_SIGNING_IDENTITY}" "./Test Project.app"

cp "./template/darwin/appdmg.config.json" "./appdmg.config.json"

echo "building macOS amd64 DMG..."
appdmg "appdmg.config.json" "./pftest_${TGT}_darwin_amd64_desktop.dmg"
zip -r "pftest_${TGT}_darwin_amd64_desktop.zip" "./Test Project.app"

# darwin arm64
cp "pftest.darwin.arm64" "./Test Project.app/Contents/MacOS/pftest"

echo "signing arm64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s "${APPLE_SIGNING_IDENTITY}" "./Test Project.app/Contents/MacOS/pftest"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s "${APPLE_SIGNING_IDENTITY}" "./Test Project.app"

echo "building macOS arm64 DMG..."
appdmg "appdmg.config.json" "./pftest_${TGT}_darwin_arm64_desktop.dmg"
zip -r "pftest_${TGT}_darwin_arm64_desktop.zip" "./Test Project.app"

# macOS universal
rm "./Test Project.app/Contents/MacOS/pftest"
lipo -create -output "./Test Project.app/Contents/MacOS/pftest" pftest.darwin pftest.darwin.arm64

echo "signing universal desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s "${APPLE_SIGNING_IDENTITY}" "./Test Project.app/Contents/MacOS/pftest"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s "${APPLE_SIGNING_IDENTITY}" "./Test Project.app"

echo "building macOS universal DMG..."
appdmg "appdmg.config.json" "./pftest_${TGT}_darwin_all_desktop.dmg"
zip -r "pftest_${TGT}_darwin_all_desktop.zip" "./Test Project.app"

# linux
echo "building Linux zip..."
zip "pftest_${TGT}_linux_amd64_desktop.zip" "./pftest"

#windows
echo "building Windows zip..."
curl -L -o webview.dll https://github.com/webview/webview/raw/master/dll/x64/webview.dll
curl -L -o WebView2Loader.dll https://github.com/webview/webview/raw/master/dll/x64/WebView2Loader.dll
zip "pftest_${TGT}_windows_amd64_desktop.zip" "./pftest.exe" "./webview.dll" "./WebView2Loader.dll"

mkdir -p "../../build/dist"
mv "./pftest_${TGT}_darwin_amd64_desktop.dmg" "../../build/dist"
mv "./pftest_${TGT}_darwin_amd64_desktop.zip" "../../build/dist"
mv "./pftest_${TGT}_darwin_arm64_desktop.dmg" "../../build/dist"
mv "./pftest_${TGT}_darwin_arm64_desktop.zip" "../../build/dist"
mv "./pftest_${TGT}_darwin_all_desktop.dmg" "../../build/dist"
mv "./pftest_${TGT}_darwin_all_desktop.zip" "../../build/dist"
mv "./pftest_${TGT}_linux_amd64_desktop.zip" "../../build/dist"
mv "./pftest_${TGT}_windows_amd64_desktop.zip" "../../build/dist"

cd "$dir/../.."
echo "Builds written to ./build/dist (pftest_${TGT}_*_desktop.*)"
