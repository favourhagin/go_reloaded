package main

import "strings"

func fixArticle(text string) string {
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
