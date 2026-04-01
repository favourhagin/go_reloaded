package main

import (
	"strings"
)

func fixQuotes(input string) string {
	inQuotes := false
	quotes := ""
	var start int
	result := ""

	for i, char := range input {
		if !inQuotes && (char == '\'' || char == '"') {
			quotes = string(char)
			inQuotes = true
			result += string(char)
			start = i + 1

		} else if inQuotes && string(char) == quotes {
			inside_text := input[start:i]
			inside_text = strings.TrimSpace(inside_text)
			result = result + inside_text + string(char)

			inQuotes = false

		} else if !inQuotes {
			result += string(char)
		}
	}
	return result
}
