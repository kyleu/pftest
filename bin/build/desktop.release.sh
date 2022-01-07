#!/bin/bash

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

TGT=$1
[ "$TGT" ] || TGT="v0.0.0"

docker build -f tools/desktop/Dockerfile -t pftest .

rm -rf tmp/release
mkdir -p tmp/release

cd "tmp/release"

id=$(docker create pftest)
docker cp $id:/dist - > ./desktop.tar
docker rm -v $id
tar -xvf "desktop.tar"
rm "desktop.tar"

mv dist/darwin_amd64/pftest "pftest.macos"
mv dist/darwin_arm64/pftest "pftest.macos.arm64"
mv dist/linux_amd64/pftest "pftest"
mv dist/windows_amd64/pftest "pftest.exe"
rm -rf "dist"

# macOS x86_64
cp -R "../../tools/desktop/template" .

mkdir -p "./Test Project.app/Contents/Resources"
mkdir -p "./Test Project.app/Contents/MacOS"

cp -R "./template/macos/Info.plist" "./Test Project.app/Contents/Info.plist"
cp -R "./template/macOS/icons.icns" "./Test Project.app/Contents/Resources/icons.icns"

cp "pftest.macos" "./Test Project.app/Contents/MacOS/pftest"

echo "signing amd64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Test Project.app/Contents/MacOS/pftest"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Test Project.app"

cp "./template/macos/appdmg.config.json" "./appdmg.config.json"

echo "building macOS amd64 DMG..."
appdmg "appdmg.config.json" "./pftest_${TGT}_macos_x86_64_desktop.dmg"
zip -r "pftest_${TGT}_macos_x86_64_desktop.zip" "./Test Project.app"

# macOS arm64
cp "pftest.macos.arm64" "./Test Project.app/Contents/MacOS/pftest"

echo "signing arm64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Test Project.app/Contents/MacOS/pftest"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Test Project.app"

echo "building macOS arm64 DMG..."
appdmg "appdmg.config.json" "./pftest_${TGT}_macos_arm64_desktop.dmg"
zip -r "pftest_${TGT}_macos_arm64_desktop.zip" "./Test Project.app"

# macOS universal
rm "./Test Project.app/Contents/MacOS/pftest"
lipo -create -output "./Test Project.app/Contents/MacOS/pftest" pftest.macos pftest.macos.arm64

echo "signing universal desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Test Project.app/Contents/MacOS/pftest"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Test Project.app"

echo "building macOS universal DMG..."
appdmg "appdmg.config.json" "./pftest_${TGT}_macos_all_desktop.dmg"
zip -r "pftest_${TGT}_macos_all_desktop.zip" "./Test Project.app"

# linux
echo "building Linux zip..."
zip "pftest_${TGT}_linux_x86_64_desktop.zip" "./pftest"

#windows
echo "building Windows zip..."
curl -L -o webview.dll https://github.com/webview/webview/raw/master/dll/x64/webview.dll
curl -L -o WebView2Loader.dll https://github.com/webview/webview/raw/master/dll/x64/WebView2Loader.dll
zip "pftest_${TGT}_windows_x86_64_desktop.zip" "./pftest.exe" "./webview.dll" "./WebView2Loader.dll"

mkdir -p "../../build/dist"
mv "./pftest_${TGT}_macos_x86_64_desktop.dmg" "../../build/dist"
mv "./pftest_${TGT}_macos_x86_64_desktop.zip" "../../build/dist"
mv "./pftest_${TGT}_macos_arm64_desktop.dmg" "../../build/dist"
mv "./pftest_${TGT}_macos_arm64_desktop.zip" "../../build/dist"
mv "./pftest_${TGT}_macos_all_desktop.dmg" "../../build/dist"
mv "./pftest_${TGT}_macos_all_desktop.zip" "../../build/dist"
mv "./pftest_${TGT}_linux_x86_64_desktop.zip" "../../build/dist"
mv "./pftest_${TGT}_windows_x86_64_desktop.zip" "../../build/dist"
