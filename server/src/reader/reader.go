package reader

import (
	cli "github.com/anwesh-b/MeroDB/lib/cli" 
	string "github.com/anwesh-b/MeroDB/lib/string" 
	parser "github.com/anwesh-b/MeroDB/server/src/parser" 
	"bufio"
	"os"
)


func InjectReader(){
	for {
		cli.CLog("MeroDB > ")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		text = string.TrimString(text)

		if text == ".exit" {
			break
		}

		parser.EvaluateInput(text)
	}
}


