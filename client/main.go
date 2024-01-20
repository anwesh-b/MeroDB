package main

import (
	"github.com/anwesh-b/MeroDB/client/src/netgrpc"
	cli "github.com/anwesh-b/MeroDB/lib/cli"
)

func main() {
	// Conection to server
	cli.CLog("Connected to server")
	cli.CLog("MeroDB server")
	cli.CLog("--------------------")
	cli.CLog("Welcome to MeroDB server. Type .exit to exit.")
	cli.CLog("")
	cli.CLog("")
	cli.CLog("")
	netgrpc.InitClient()
}
