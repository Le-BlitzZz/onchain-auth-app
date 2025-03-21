# Build stage.
FROM authonchain/develop:1.24.1-alpine3.21 AS build

# Copy source to image.
WORKDIR "/go/src/github.com/Le-BlitzZz/onchain-auth-app"
COPY . .

# Build app.
RUN make install

# Production stage.
FROM authonchain/develop:1.24.1-alpine3.21

# Set default working directory.
WORKDIR /authonchain

# Export HTTP port.
EXPOSE 8080

# Copy app files.
COPY --from=build --chmod=755 /opt/authonchain/ /opt/authonchain

# Start app.
CMD ["/scripts/cmd.sh", "/opt/authonchain/bin/authonchain", "start"]
