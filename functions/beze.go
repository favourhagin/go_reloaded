package main

import (
	"strconv"
)

func binToDecimal(binstr string) (int64, error) {

	return strconv.ParseInt(binstr, 2, 64)
}
