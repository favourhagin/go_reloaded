package main

import (
	"fmt"
	"regexp"
)

func fixPunctuation(text string) string {
	// remove spaces before .,!?;:
	re := regexp.MustCompile(`\s+([.,!?;:])`)
	text = re.ReplaceAllString(text, "$1")
	return text
}
func 