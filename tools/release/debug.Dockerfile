# Content managed by Project Forge, see [projectforge.md] for details.
FROM golang:alpine

LABEL "org.opencontainers.image.authors"="Kyle U"
LABEL "org.opencontainers.image.source"="https://github.com/kyleu/pftest"
LABEL "org.opencontainers.image.vendor"="kyleu"
LABEL "org.opencontainers.image.title"="Test Project"
LABEL "org.opencontainers.image.description"="A Test application, built with Project Forge"

RUN apk add --update --no-cache ca-certificates tzdata bash curl htop libc6-compat

RUN apk add --no-cache ca-certificates dpkg gcc git musl-dev \
    && mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && go get github.com/go-delve/delve/cmd/dlv

SHELL ["/bin/bash", "-c"]

# main http port
EXPOSE 41000
# marketing port
EXPOSE 41001

WORKDIR /

ENTRYPOINT ["/pftest", "-a", "0.0.0.0"]

COPY pftest /
