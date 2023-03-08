#!/bin/bash

export PATH=$PATH:~/.local/go/bin
export PATH=$PATH:~/.local/protoc/bin
export PATH="$PATH:$(go env GOPATH)/bin"

gobin="$(go env GOPATH)/bin"
./COMPILE-PROTOS.sh
go build
mv gnostic-grpc $gobin

#gopath=$(go env GOPATH)
#./COMPILE-PROTOS.sh
#cd plugin
#echo ${gopath}/bin/gnostic-grpc
#go build -o ${gopath}/bin/gnostic-grpc 
