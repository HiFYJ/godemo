package main

import "fmt"

/**
 1、defer 关键字作用类似于finally，在执行最后执行（文件操作关闭流等）
 2、defer可以有多个，采用栈的形式保存，先入后出
 3、defer和return先返回return，后执行defer（defer是在括号执行结束之后，也就是return执行完成）
**/
func main() {
	//测试defer执行顺序
	//defer fmt.Println("main finished")
	//defer fmt.Println("main finished2")
	//defer fmt.Println("main finished3")
	//
	//fmt.Println("content output")
	//fmt.Println("content output2")

	//测试defer执行方法的顺序
	//defer fun1()
	//defer fun2()
	//defer fun3()

	testDeferAndReturn()
}
func testDeferAndReturn() int {
	//测试defer和return执行顺序
	defer deferFun() //defer会在{}执行结束后执行
	return returnFun()
}
func returnFun() int {
	fmt.Println("return function call****")
	return 0
}

func deferFun() int {
	fmt.Println("defer function call****")
	return 0
}

func fun1() {
	fmt.Println("function A")
}
func fun2() {
	fmt.Println("function B")
}
func fun3() {
	fmt.Println("function C")
}
