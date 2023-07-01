package parser

import (
	cli "github.com/anwesh-b/MeroDB/lib/cli" 
	"strings"
)

func EvaluateInput(str string) {
	if strings.HasPrefix(str, "insert"){
		cli.CLog("Inserting data")
	} else if strings.HasPrefix(str, "select"){
		cli.CLog("Selecting data")
	} else if strings.HasPrefix(str, "update"){
		cli.CLog("Updating data")
	} else if strings.HasPrefix(str, "delete"){
		cli.CLog("Deleting data")
	} else {
		cli.CLog("Invalid command")
	}	
	cli.CLog("")
}
