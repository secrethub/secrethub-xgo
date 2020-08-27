#!/bin/bash

mkdir app
go build -o secrethub.a -buildmode=c-archive ./secrethub_wrapper.go ./error_handling.go
make swig
make
dotnet publish build.csproj -c Debug -o app
make clean