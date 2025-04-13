FROM golang:1.24.1-alpine3.21 AS go-build
FROM node:22-alpine3.21 AS node-build

FROM alpine:3.21

ENV NODE_ENV="production" \
    GOPATH="/go" \
    GOBIN="/usr/local/bin"

ENV PATH="$GOPATH/bin:/usr/local/go/bin:$PATH"

# Copy Go runtime
COPY --from=go-build /usr/local/go /usr/local/go
COPY --from=go-build /go /go
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 1777 "$GOPATH"

# Copy Node.js runtime
COPY --from=node-build /usr/local /usr/local
COPY --from=node-build /opt /opt

# Copy scripts.
COPY --chmod=755 /scripts/ /scripts/

# Add dependencies
RUN apk add --no-cache bash make libstdc++ libgcc && \
    echo "Installing Go debugger..." && \
    go install github.com/go-delve/delve/cmd/dlv@latest

# Set default working directory.
WORKDIR "/go/src/github.com/Le-BlitzZz/onchain-auth-app"

# Expose HTTP port.
EXPOSE 8080

# Set the default command.
CMD ["/scripts/cmd.sh", "tail", "-f", "/dev/null"]
