package main

import "fmt"

func main() {
	/**
	变量声明方式
	**/
	/**
	   注：
	  	声明全局变量只能采用1、2、3方法
	  	:=方法只能声明函数体内的变量
	  **/

	//	第一种：声明一个变量
	var a int
	fmt.Println("a =", a)
	fmt.Printf("a type = %T\n", a)

	//	第二种：声明一个变量并初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("b type =  %T\n", b)

	//	第三种：声明一个变量省略数据类型，通过值匹配当前变量的数据类型
	var bb = 10010
	fmt.Println("bb=", bb)
	fmt.Printf("bb type = %T \n", bb)

	//	第四种：省去var关键字，直接自动匹配(常用)
	c := 188
	fmt.Println("c=", c)
	fmt.Printf("c type = %T \n", c)

	e := "string"
	fmt.Println("e=", e)
	fmt.Printf("e type = %T \n", e)

	f := 1.89
	fmt.Println("f=", f)
	fmt.Printf("f type = %T \n", f)

	//	声明多个变量
	var x, y int = 123, 234
	fmt.Println("x:", x, "y:", y)

	var xx, yy = 123, "234"
	fmt.Println("xx:", xx, "yy:", yy)
	fmt.Printf("xx:%T ", xx)
	fmt.Printf(" yy:%T\n", yy)
	//	多行多变量声明
	var (
		i int  = 123
		j bool = false
	)
	fmt.Println("i: ", i, "j: ", j)
}
