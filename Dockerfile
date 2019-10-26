FROM golang:1.13-alpine

ARG PEER=716a2292c6d60b140e8af18575e4fb7bc1ae3f99@tcp:192.168.1.199:26656
ARG MONIKER=$HOSTNAME

ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python

WORKDIR /go/src/github.com/cratos-network/cratos-sdk
COPY . .

# Install minimum necessary dependencies, build Cratos SDK, remove packages
RUN apk add --no-cache $PACKAGES && \
    make tools && \
    make install

# Final image
FROM alpine:edge

ENV GOPATH=/go
ENV GOBIN=$GOPATH/bin
ENV GO111MODULE=on

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/cratosd /usr/bin/cratosd
COPY --from=build-env /go/bin/cratosd /usr/bin/cratoscli

RUN cratosd unsafe-reset-all
RUN curl --silent -XGET https://raw.githubusercontent.com/cratos-network/testnets/master/latest/genesis.json > $HOME/.cratosd/config/genesis.json
# Set the moniker has
RUN sed -i -E 's/moniker[^"]*"([^"]*)"/moniker = "${MONIKER}"/g' ~/.cratosd/config/config.toml
RUN sed -E 's/persistent_peers.*"()"/persistent_peers = "${PEER}"/g' ~/.cratosd/config/config.toml

# Run gaiad by default, omit entrypoint to ease using container with gaiacli
CMD ["cratosd", "start"]    