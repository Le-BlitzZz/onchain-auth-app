#!/usr/bin/env bash

if [[ -z $1 ]]; then
  echo "Usage: ${0##*/} [filename]" 1>&2
  exit 1
fi

set -e

BUILD_BIN=${1:-authonchain}

echo "Building AuthOnchain..."

BUILD_CMD=(go build -gcflags="all=-N -l" -o "${BUILD_BIN}" cmd/authonchain/app.go)

# Build app binary.
echo "=> compiling \"$BUILD_BIN\""
echo "=> ${BUILD_CMD[*]}"
"${BUILD_CMD[@]}"

# Display binary size.
du -h "${BUILD_BIN}"

echo "Done."
