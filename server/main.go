package main

import (
	cli "github.com/anwesh-b/MeroDB/lib/cli" 
	reader "github.com/anwesh-b/MeroDB/server/src/reader"
)

func main() {
	cli.CLog("MeroDB server")
	cli.CLog("--------------------")
	cli.CLog("Welcome to MeroDB server. Type .exit to exit.")
	cli.CLog("")
	cli.CLog("")
	cli.CLog("")
	reader.InjectReader()
	cli.CLog("Thanks for using MeroDB.")
}
