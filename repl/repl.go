package repl

import (
	"strings"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)

	strSlice := strings.Fields(lowered)

	return strSlice
}
