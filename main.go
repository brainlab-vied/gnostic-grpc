package main

import "github.com/brainlab-vied/gnostic-grpc/plugin"

//go:generate sh COMPILE-PROTOS.sh

func main() {
	plugin.Main()
}
