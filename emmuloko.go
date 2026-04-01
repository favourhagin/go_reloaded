package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hexToDecimal(hexstr string) string {
	e := strings.Fields(hexstr)
	for i := 0; i < len(e)-1; i++ {

		if e[i] == "(hex)" {
			output, _ := strconv.ParseInt(e[i-1], 16, 64)
			e[i-1] = strconv.FormatInt(output, 10)
			e = append(e[:i], e[i+1:]...)

		}
	}

	return strings.Join(e, " ")
}

func main() {
	fmt.Println(hexToDecimal("1E (hex) files were added"))
}
