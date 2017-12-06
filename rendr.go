package main

//go:generate protoc -I rpc/ rpc/rendr.proto --go_out plugins=grpc:rpc

import (
	"github.com/kbence/rendr/cmd"
)

func main() {
	cmd.Execute()
}
