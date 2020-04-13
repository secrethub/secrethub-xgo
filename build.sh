#!/bin/bash

set -e

protoc --java_out lite:clients/java/src/main/java \
    --go_out bridge \
    --python_out clients/python/secrethub \
    secrethub.proto

NAME=secrethub
TARGET_OS=linux-amd64
EXT=".so"
if [[ "${OSTYPE}" == "darwin"* ]]; then
    EXT=".dylib"
    TARGET_OS="darwin"
fi

mkdir -p build/${TARGET_OS}    
OUTPUT="build/${TARGET_OS}/lib${NAME}${EXT}"
echo "Building ${OUTPUT}"
go build -buildmode=c-shared -o ${OUTPUT} secrethub.go

rm -rf clients/java/src/main/resources/darwin/
rm -rf clients/java/src/main/resources/linux-amd64/

cp -R build/* clients/java/src/main/resources/
cp -R build/* clients/python/secrethub/resources/
