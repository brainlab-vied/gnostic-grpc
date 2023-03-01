#!/bin/bash
gobin="$(go env GOPATH)/bin"
./COMPILE-PROTOS.sh
go build
mv gnostic-grpc $gobin
