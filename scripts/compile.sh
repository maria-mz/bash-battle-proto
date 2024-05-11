#!/bin/bash
# Compile the *.proto files and update generated protobuf code

echo "Starting script..."

echo "Cleaning up directory..."

DIR_NAME="proto"

rm -rf $DIR_NAME
mkdir $DIR_NAME

GOPATH=$(echo `go env GOPATH`)

export PATH="$PATH:$GOPATH/bin"

echo "Compiling..."

protoc --go_out=$DIR_NAME --go_opt=paths=source_relative --go-grpc_out=$DIR_NAME --go-grpc_opt=paths=source_relative *.proto

echo "Done!"
