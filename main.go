package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := int64(10)
	b := strconv.FormatInt(a, 10)
	fmt.Println(b)
}
