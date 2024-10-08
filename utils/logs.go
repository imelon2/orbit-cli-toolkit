package utils

import (
	"fmt"
	"log"

	"github.com/hokaccha/go-prettyjson"
)

func PrintPrettyJson(v interface{}) {
	formatter := prettyjson.NewFormatter()
	formatter.Indent = 2

	coloredJson, err := formatter.Marshal(v)
	if err != nil {
		log.Fatalf("Failed to Marshal calldata: %v", err)
	}
	fmt.Println(string(coloredJson))
}

func BoldString(str string) string {
	return "\033[1m" + str + "\033[0m"
}

func BoldGreenString(str string) string {
	return "\x1b[32;1m" + str + "\033[0m"
}
