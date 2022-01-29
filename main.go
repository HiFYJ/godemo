package main

import (
	"fmt"
)

func main() {
	//testHelloWorld("hello world")
	//structDemo.StructDemo3()

	//user:=structDemo.NewUser("张三",2,"北京")
	//user.Introduce()
	//structDemo.JsonDemo()
	//var num1 int8 = 8
	//var num2 int8 = 18
	//
	//var re = calc.Add(num1,num2)
	//fmt.Printf("res:%d",re)


	//interfaceDemo.Demo()
	//var dog interfaceDemo.Dog
	//dog.Say()

	//var wang interfaceDemo.Wang
	//wang.Move()
	/*值接收者*/
	//var x interfaceDemo.Mover
	//var dog = interfaceDemo.Wang{}
	//x = dog
	/*
	dog指针wangcai内部会自动求值*wangcai
	*/
	//var wangcai = &dog //fc就是*Wang类型
	//x=wangcai
	//x.Move()

	//var y interfaceDemo.Say
	/*
	var d = interfaceDemo.Wang{} 这里声明的d为Wang类型
	y = d y无法接收Wang类型
	*/
	//var fc = &interfaceDemo.Wang{}  //fc就是*Wang类型
	//y = fc
	//y.Say()

	//var xyj = interfaceDemo.Haier{}
	//xyj.Dry()
	//xyj.Wash()

	//reflectDemo.Input()
	//reflectDemo.StructInPut()
}

func testHelloWorld(str string) {
	fmt.Println(str)
}