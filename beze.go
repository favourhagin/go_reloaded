package main

import (
	"fmt"
	"strconv"

)

func binToDecimal(binstr string)(int64,error) {

return strconv.ParseInt(binstr,2,64)
}

func main() {
	fmt.Println(binToDecimal("10"))
}



