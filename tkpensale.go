package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CapN(text string) string {
	word := strings.Fields(text)

	for i := len(word) - 1; i > 0; i-- {
		switch {
		case i > 0 && word[i] == "(cap)":
			for j := i - 1; j < i; j++ {
				word[j] = strings.Title(word[j])
			}
			word = append(word[:i], word[i+1:]...)
		case i > 0 && word[i] == "(cap,":
			text := strings.TrimSuffix(word[i+1], ")")
			n, err := strconv.Atoi(text)
			if err != nil {
				n = 1
			}
			for j := i - n; j < i; j++ {
				word[j] = strings.Title(word[j])
			}
			word = append(word[:i], word[i+2:]...)
		case i > 0 && strings.Contains(word[i], ","):
			n := extractNum(word[i])
			for j := i - n; j < i; j++ {
				word[j] = strings.Title(word[j])
			}
			word = append(word[:i], word[i+1:]...)
		}

	}
	return strings.Join(word, " ")
}
func extractNum(cmd string) int {
	val := strings.Split(cmd, ",")
	if len(val) != 2 {
		return 1
	}
	n, err := strconv.Atoi(val[1])
	if err != nil {
		return 1
	}
	return n
}

func main() {
	fmt.Println(CapN("Welcome to the  Brooklyn bridge (cap, )"))
	fmt.Println(CapN("This is so exciting (cap,2)"))
	fmt.Println(CapN("it (cap) was the best of times"))
	fmt.Println(CapN("it is just (cap, 2) all about fun"))
	fmt.Println(CapN("chatgbt is not no 1 in (cap, 2)"))
}
