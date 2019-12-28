package main

import (
	"github.com/pojntfx/gon2n/pkg"
)

func main() {
	edge := pkg.Edge{
		DeviceName:    "edge1",
		NetworkName:   "dev",
		SecretKey:     "asdf",
		MyMacAddress:  "DE:AD:BE:EF:01:10",
		MyIPv4Addr:    "10.0.0.2",
		Supernode:     "192.168.178.35:1234",
		KeepOnRunning: true,
	}

	edge.Start()
}
