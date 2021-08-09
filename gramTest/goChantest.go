package main

import (
	"fmt"
	"sync"
	"time"
)

type grep struct {
	Name string
	Age  int
}

func main() {

	var wg sync.WaitGroup
	bb := []grep{}
	bb = append(bb, grep{Age: 18, Name: "fan"})
	bb = append(bb, grep{Age: 19, Name: "ya"})
	bb = append(bb, grep{Age: 20, Name: "hui"})
	cc := make(chan []grep, len(bb))

	for _, v := range bb {
		wg.Add(1)
		fmt.Println("123")
		go wgTest1(&wg, cc, &v)
	}
	wg.Wait()
	close(cc)
	var vd []grep
	for bb := range cc {
		vd = append(vd, bb...)
	}
	fmt.Println(vd)
}

func wgTest1(wg *sync.WaitGroup, ch chan<- []grep, gr *grep) {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	gb := []grep{}
	gp := grep{
		Name: gr.Name,
		Age:  gr.Age,
	}

	ch <- append(gb, gp)

}
