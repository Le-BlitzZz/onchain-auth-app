FROM authonchain/develop:alpine3.21

# Set default working directory.
WORKDIR "/go/src/github.com/Le-BlitzZz/onchain-auth-app"

# Copy source to image.
COPY . .
