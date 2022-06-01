package main

import (
	"fmt"
	"time"
)

/**
有缓冲的channel
特点：当channel满了时，写数据的go程会阻塞；channel为空时，读数据的go程也会阻塞
**/
func main() {
	//创建有缓存的channel 空间为3
	c := make(chan int, 3)
	fmt.Println("channel len(c) = ", len(c), ",cap(c) is ", cap(c))

	go func() {
		defer fmt.Println("子程执行结束")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子程运行中,新增元素", i, ",len(c) = ", len(c), ",cap(c) =", cap(c))
		}
	}()

	time.Sleep(time.Second * 2)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}
	fmt.Println("main end")
	/** output：
	channel len(c) =  0 ,cap(c) is  3
	子程运行中,新增元素 0 ,len(c) =  1 ,cap(c) = 3
	子程运行中,新增元素 1 ,len(c) =  2 ,cap(c) = 3
	子程运行中,新增元素 2 ,len(c) =  3 ,cap(c) = 3
	num =  0
	子程运行中,新增元素 3 ,len(c) =  3 ,cap(c) = 3
	子程执行结束
	num =  1
	num =  2
	num =  3
	main end
	**/
}
