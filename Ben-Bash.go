package main

import (
	"fmt"
	"strings"
)

func fixArticles(text string) string {
	word := strings.Fields(text)

	for i := 0; i < len(word)-1; i++ {

		if word[i] == "An" && !strings.ContainsAny(word[i+1][:1], "aeriouAERIOUHh") {
			word[i] = "A"
		}
		if word[i] == "A" && strings.ContainsAny(word[i+1][:1], "aeriouAERIOUHh") {
			word[i] = "An"
		}
	}
	return strings.Join(word, " ")
}
func main() {
	fmt.Println(fixArticles("the it was man. A amazing rock. A honest man. A book"))
}
