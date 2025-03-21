FROM authonchain/develop:1.24.1-alpine3.21

# Set default working directory.
WORKDIR "/go/src/github.com/Le-BlitzZz/onchain-auth-app"

# Copy source to image.
COPY . .
