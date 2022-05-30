package main

import "fmt"

/**
万能数据类型interface使用
**/

//interface {}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("called myFunc****")
	fmt.Println("arg is :", arg)

	//	区分数据类型 给interface提供了“类型断言”的机制
	value, ok := arg.(string)

	fmt.Println("arg value is "+value, ",ok is string:", ok)

}

type Book struct {
	auth string
}

func main() {
	book := Book{auth: "墨菲定律"}

	myFunc(book)

	myFunc("true") //arg value is true ,ok is string: true
	myFunc(10086)
}
