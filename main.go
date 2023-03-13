package main

import "github.com/laurenz-eschwey-bl/gnostic-grpc/plugin"

//go:generate sh COMPILE-PROTOS.sh

func main() {
	plugin.Main()
}
