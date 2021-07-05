package main

import (
	"fmt"
	"learn_golang/learn/model/yk_base"
)

func main()  {
	var user yk_base.User
	data, err := yk_base.FindUserById(&user, "5ec76168aa838b333828d67a")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)

}