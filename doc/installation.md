# Installation

## Pre-built binaries
Download any package from the [release page](https://github.com/kyleu/pftest/releases).

### Homebrew
```shell
brew install kyleu/kyleu/pftest
```

### deb, rpm and apk packages
Download the .deb, .rpm or .apk packages from the [release page](https://github.com/kyleu/pftest/releases) and install them with the appropriate tools.

## Running with Docker
```shell
docker run -p 41000:41000 ghcr.io/kyleu/pftest:latest
docker run -p 41000:41000 ghcr.io/kyleu/pftest:latest-debug
```

## Built from source

### go install
```shell
go install github.com/kyleu/pftest@latest
```

### Source code

If you want to contribute to the project, please follow the steps on our [contributing guide](contributing).

If you just want to build from source for whatever reason, follow these steps:

```shell
git clone https://github.com/kyleu/pftest
cd pftest
make build
./build/debug/pftest --help
```

A script has been provided at `./bin/dev.sh` that will auto-reload when the project changes.

Note that you may need to run `./bin/bootstrap.sh` before building the project for the first time.
