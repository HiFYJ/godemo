package main

import "fmt"

/**
 面向对象思想之对象类的表示与封装
	类名、属性名、方法名首字母大写表示对外(其他包)可以访问，否则只能够在本包内访问
**/
func main() {
	user := User{}
	user.SetUserName("张三")
	user.Tel = "12313164581320"
	user.Age = 18
	user.PrintUser()
}

//User大写表示其他包可以引用
type User struct {
	//属性名大写表示该属性是对外能够访问的
	UserName string
	Age      int
	Tel      string
}

func (this *User) GetUserName() string {
	return this.UserName
}

//
//func (this User)setUserName(name string)  {
//	//操作的都是副本（拷贝），无法改变原值
//	this.userName = name
//}
func (this *User) SetUserName(name string) {
	//通过指针传递，可以改变原值
	this.UserName = name
}
func (this *User) PrintUser() {
	fmt.Println("userName:", this.UserName)
	fmt.Println("age:", this.Age)
	fmt.Println("tel:", this.Tel)
}
