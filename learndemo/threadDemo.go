package main

import (
	"fmt"
	"time"
)

func goFun(i int) {
	fmt.Println("goroutine ", i, "*****")
}
func ThreadMain() {
	for i := 0; i < 1000; i++ {
		goFun(i) //开启一个并发协程
	}
	time.Sleep(time.Second * 2)
}
func main() {
	ThreadMain()
}
