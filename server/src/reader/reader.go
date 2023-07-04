package reader

import (
	"bufio"
	"fmt"
	"os"

	string "github.com/anwesh-b/MeroDB/lib/string"
	parser "github.com/anwesh-b/MeroDB/server/src/parser"
)

func InjectReader() {
	for {
		fmt.Print("MeroDB > ")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		text = string.TrimString(text)

		if text == ".exit" {
			break
		}

		parser.EvaluateInput(text)
	}
}
