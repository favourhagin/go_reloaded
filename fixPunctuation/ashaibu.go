package main

import (
	"fmt"
	"regexp"
)

func fixPunctuation(text string) string {
	// remove spaces before .,!'?;:
	re := regexp.MustCompile(`\s+([.,!'?;:])`)
	text = re.ReplaceAllString(text, "$1")
	return text
}
func main() {
	fmt.Println(fixPunctuation("I was thinking ... You were right!"))
	fmt.Println(fixPunctuation("I was exactly how they described me'AWESOME'"))
	fmt.Println(fixPunctuation(" I am coming to you '.     "))
	fmt.Println(fixPunctuation("I was just , on my own    "))
	fmt.Println(fixPunctuation("Dont you think you are crazy   ?"))

}
