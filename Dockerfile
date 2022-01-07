# Build image
FROM golang:alpine AS builder

ENV GOFLAGS="-mod=readonly"

RUN apk add --update --no-cache bash ca-certificates make git curl build-base
RUN go get -u github.com/valyala/quicktemplate/qtc

RUN mkdir /app
WORKDIR /app

ADD ./go.mod        /app/go.mod
ADD ./go.sum        /app/go.sum

RUN go mod download

ADD ./app           /app/app
ADD ./assets        /app/assets
ADD ./bin           /app/bin
ADD ./main.go       /app/main.go
ADD ./Makefile      /app/Makefile
ADD ./queries       /app/queries
ADD ./views         /app/views

# $PF_SECTION_START(dockerbuild)$
# $PF_SECTION_END(dockerbuild)$

RUN go mod download
RUN set -xe && bash -c 'make build-release'
RUN mv build/release /build

# Final image
FROM alpine

COPY --from=builder /build/* /

# $PF_SECTION_START(dockerimage)$
# $PF_SECTION_END(dockerimage)$

# main http port
EXPOSE 41000
# marketing port
EXPOSE 41001

ENTRYPOINT ["/pftest", "-a", "0.0.0.0"]
