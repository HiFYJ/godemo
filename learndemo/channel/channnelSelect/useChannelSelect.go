package main

import "fmt"

/**
channel 的select的使用
select具备多路channel的监控状态功能
**/
func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacii(c, quit)
}

func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			//如果c可以写
			x = y
			y = x + y
		case <-quit:
			//如果quit可以读
			fmt.Println("quit")
			return
		}
	}
}
