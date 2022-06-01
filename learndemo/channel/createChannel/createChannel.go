package main

import (
	"fmt"
	"time"
)

/**
channel的创建使用
**/
func main() {
	//定义一个channel
	c := make(chan int)

	fmt.Println("start main")
	go func() {
		defer fmt.Println(" goroutine end")
		fmt.Println("执行goroutine")
		time.Sleep(1 * time.Second)
		c <- 233 //将233发送给channel c
	}()

	/**
	如果子程始终不设置c值，main获取c值时会阻塞；
	如果main始终不获取c值，子程会阻塞在设置时
	**/
	num := <-c //从channel中接收数据 这里会阻塞
	fmt.Println("main num is ", num)
	fmt.Println("end")

	/*output：
	start main
	执行goroutine
	 goroutine end
	main num is  233
	end
	*/
}
