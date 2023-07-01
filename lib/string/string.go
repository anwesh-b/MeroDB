package string

import (
	"strings"
)

func TrimString(str string) string {
	return strings.Trim(str, " \t\n")
}
