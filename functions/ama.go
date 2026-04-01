package main

import (
	"strconv"
	"strings"
)

func lowN(text string) string {
	word := strings.Fields(text)

	for i := len(word) - 1; i > 0; i-- {
		if word[i] == "(low," {
			numstr := strings.TrimSuffix(word[i+1], ")")
			n, _ := strconv.Atoi(numstr)
			for j := i - n; j < i; j++ {
				word[j] = strings.ToLower(word[j])
			}
			word = append(word[:i], word[i+2:]...)
		}
	}
	return strings.Join(word, " ")
}
