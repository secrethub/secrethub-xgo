#!/bin/bash

set -e

NAME=secrethub
EXT=".so"
if [[ "${OSTYPE}" == "darwin"* ]]; then
    EXT=".dylib"
fi

mkdir -p lib    
OUTPUT="lib/lib${NAME}${EXT}"
echo "Building ${OUTPUT}"
go build -buildmode=c-shared -o ${OUTPUT} secrethub.go

