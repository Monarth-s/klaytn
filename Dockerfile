# Global ARGs
ARG DOCKER_BASE_IMAGE=klaytn/build_base:latest
ARG PKG_DIR=/klaytn-docker-pkg
ARG SRC_DIR=/go/src/github.com/Monarth-s/klaytn

FROM ${DOCKER_BASE_IMAGE} AS builder
LABEL maintainer="Tony Lee <tony.jm@krustuniverse.com>"
ARG SRC_DIR
ARG PKG_DIR

ARG VINITN_RACE_DETECT=0
ENV VINITN_RACE_DETECT=$VINITN_RACE_DETECT

ARG VINITN_STATIC_LINK=0
ENV VINITN_STATIC_LINK=$VINITN_STATIC_LINK

ARG VINITN_DISABLE_SYMBOL=0
ENV VINITN_DISABLE_SYMBOL=$VINITN_DISABLE_SYMBOL

WORKDIR $SRC_DIR
# Cache default $GOMODCACHE
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download -x

# Cache default $GOCACHE
# First 'make kcn' to populate build cache and then 'make all' in parallel
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    make kcn && \
    make all -j

FROM --platform=linux/amd64 ubuntu:20.04
ARG SRC_DIR
ARG PKG_DIR

RUN apt update && \
            apt install -yq musl-dev ca-certificates && \
            update-ca-certificates && \
            mkdir -p $PKG_DIR/conf $PKG_DIR/bin

# Startup scripts and binaries must be in the same location
COPY --from=builder $SRC_DIR/build/bin/* $PKG_DIR/bin/
COPY --from=builder $SRC_DIR/build/packaging/linux/bin/* $PKG_DIR/bin/
COPY --from=builder $SRC_DIR/build/packaging/linux/conf/* $PKG_DIR/conf/

ENV PATH=$PKG_DIR/bin:$PATH

EXPOSE 8551 8552 32323 61001 32323/udp
