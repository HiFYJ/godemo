package main

import (
	"fmt"
	"time"
)

/**
创建goroutine
**/
//主goroutine
func main() {

	//go newTask()//创建一个go程，去执行newTask流程

	/*	i:=0
		for{
			i++
			fmt.Println("main goroutine i=",i)
			time.Sleep(1*time.Second)
		}*/
	//output：
	//main goroutine i= 1
	//new task goroutine i= 1
	//new task goroutine i= 2
	//main goroutine i= 2

	/*如果main此时结束，线程结束*/
	//fmt.Println("main 执行结束")

	//	匿名函数创建子程(一个形参为空，返回值为空的函数)
	/**
	这里不能使用return返回
	**/
	/*go func() {
		defer fmt.Println("a defer")
		go func(){
			defer fmt.Println("b defer")
			//退出当前goroutine
			runtime.Goexit()
			fmt.Println("B")
		}() //这里的()表示调用函数
		fmt.Println("A")
	}()*/

	/**
	 这里虽然子协程有返回值，因为main程和子程是异步的，所以这里无法接收子程的返回值
	需要接收的话，用到后面的channel
	**/
	go func(a int, b int) bool {
		fmt.Println("a is ", a, ", b is ", b)
		return true
	}(10, 20)
	//a is  10 , b is  20
	i := 0
	for {
		i++
		time.Sleep(1 * time.Second)
	}
}

//子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Println("new task goroutine i=", i)
		time.Sleep(1 * time.Second)
	}
}
