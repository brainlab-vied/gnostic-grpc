package main

import "log"
import "github.com/laurenz-eschwey-bl/gnostic-grpc/plugin"

//go:generate sh COMPILE-PROTOS.sh

func main() {
	log.Print("main")
	plugin.Main()
}
