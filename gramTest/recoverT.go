package main

import "fmt"

func main() {
	testa()
	testb(20) //当值是1的时候，就不会越界，值是20的时候，就会越界报错。
	testc()
}

func testa() {
	fmt.Println("aaaaaaaaaaaaaaaaa")
}

func testb(x int) {

	var a [10]int
	a[x] = 111 //当x为20时候，导致数组越界，产生一个panic，导致程序崩溃
}

func testc() {
	fmt.Println("cccccccccccccccccc")
}
