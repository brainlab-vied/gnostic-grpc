package main

import "github.com/42crunch/gnostic-grpc/plugin"

//go:generate sh COMPILE-PROTOS.sh

func main() {
	plugin.Main()
}
