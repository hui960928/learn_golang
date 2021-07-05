package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type str struct {
	Name string
	Age  int64
}

type strs struct {
	Name   string
	Age    int64
	Gender string
}

func main() {
	var s1 str
	s1.Name = "范"
	s1.Age = 12

	var s2 strs
	copier.Copy(&s2, &s1)
	fmt.Println(s2)

}
