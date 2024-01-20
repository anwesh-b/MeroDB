#!/bin/bash

export PATH="$PATH:$(go env GOPATH)/bin"

rm -rf ./server/pb
rm -rf ./client/pb
protoc ./pb/*.proto --go_out=./server --go-grpc_out=./server --go_opt=paths=source_relative 
protoc ./pb/*.proto --go_out=./client --go-grpc_out=./client --go_opt=paths=source_relative 
