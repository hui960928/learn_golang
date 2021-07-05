package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	a := "ttt?yyy"
	index := strings.LastIndex(a, "?")
	fmt.Println(a[index+1:])
	ff := time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Println(ff)
}
