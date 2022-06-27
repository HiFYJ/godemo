package main

import "fmt"

/**
 结构体使用
	1、定义
	2、赋值
	3、参数(副本)
**/
func main() {
	//userType()
	var temp = Book{
		title: "1111",
		auth:  "2222",
	}
	fmt.Printf("值：%v  ----类型：%T\n", temp, temp)

	temp2 := Book{
		title: "123",
		auth:  "231",
	}
	fmt.Printf("值：%v  ----类型：%T\n", temp2, temp2)

	temp3 := &Book{
		title: "1234",
		auth:  "23123",
	}
	fmt.Printf("值：%v  ----类型：%T\n", temp3, temp3)
}

func userType() {

	//	声明一种行的数据类型myint，是int的一个别名
	//	type myint int
	//	var a myint = 10
	//	fmt.Printf("value:%d,type:%T\n",a,a)

	var mybook Book
	mybook.title = "活着"
	mybook.auth = "余华"
	fmt.Println(mybook) //{活着 余华}
	changeBook(mybook)
	fmt.Println(mybook) //{活着 余华}
	changeBook2(&mybook)
	fmt.Println(mybook) //{活着 村上春树}
}

type Book struct {
	title string
	auth  string
}

func changeBook(mybook Book) {
	//参数是作为一个副本传递的，不会影响原来的值
	mybook.auth = "村上春树"
}
func changeBook2(mybook *Book) {
	//指针传递，会影响原来的值
	mybook.auth = "村上春树"
}
