package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dir, err := os.Getwd() // 获取当前路径
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
	a := strings.Index(dir, "\\")
	b := strings.Index(dir, "learn_golang")
	fmt.Println(dir[a:b])
}
