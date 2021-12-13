package main

import (
	"demo/interfaceDemo"
	"fmt"
)

func main() {
	//testHelloWorld("hello world")
	//structDemo.StructDemo3()

	//user:=structDemo.NewUser("张三",2,"北京")
	//user.Introduce()
	//structDemo.JsonDemo()
	/*var num1 int8 = 8
	var num2 int8 = 18

	var re = calc.Add(num1,num2)
	fmt.Printf("res:%d",re)*/


	//interfaceDemo.Demo()
	/*var dog interfaceDemo.Dog
	dog.Say()*/

	/*var wang interfaceDemo.Wang
	wang.Move()*/
	// 值接收者
	var x interfaceDemo.Mover
	var dog = interfaceDemo.Wang{}
	x = dog
/*	x.Move()
	x.Say()*/
	//dog指针wangcai内部会自动求值*wangcai
	var wangcai = &dog
	x=wangcai
	x.Say()
	x.Move()
}

func testHelloWorld(str string) {
	fmt.Println(str)
}