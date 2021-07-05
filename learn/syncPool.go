package main

import (
	"encoding/json"
	"sync"
)

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

func main() {
	var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})
	stu := studentPool.Get().(*Student)
	json.Unmarshal(buf, stu)
	studentPool.Put(stu)
}
