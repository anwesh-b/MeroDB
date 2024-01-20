package main

import (
	cli "github.com/anwesh-b/MeroDB/lib/cli"
	netgrpc "github.com/anwesh-b/MeroDB/server/src/netgrpc"
	reader "github.com/anwesh-b/MeroDB/server/src/reader"
)

const port = "8080"
const portC = ":" + port

func main() {
	cli.CLog("Starting MeroDB server")
	cli.CLog("--------------------")

	netgrpc.InitServer()
	cli.CLog("new server instance created")

	cli.CLog("Welcome to MeroDB server. Type .exit to exit.")
	cli.CLog("")
	cli.CLog("")
	cli.CLog("")
	reader.InjectReader()
	cli.CLog("Thanks for using MeroDB.")
}
