package main

import (
	"fmt"
	"reflect"
)

/**
使用reflect两个方法
	1、TypeOf
	2、ValueOf
**/
func main() {
	//num:=3.1415926
	//getReflectInfo(num)
	//
	//dog := interfaceDemo.Dog{}
	//getReflectInfo(dog)

	user := User{111, "ss", 18}

	ReflectMethod(user)
}

func getReflectInfo(arg interface{}) {
	fmt.Println("type is ", reflect.TypeOf(arg))
	fmt.Println("value is ", reflect.ValueOf(arg))
}

//
//type User struct {
//	Id int
//	Name string
//	Age int
//}
//
//func (this User)CallMethod()  {
//	fmt.Println("user is call")
//	fmt.Printf("%v\n",this)
//}

func ReflectMethod(input interface{}) {
	//	获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("input type is ", inputType.Name())

	//	获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("value is ", inputValue)

	//	通过type获取里面的字段
	//1、获取interface的reflect.Type,通过Type得到NumField，进行遍历
	//2、得到每个field，数据类型
	//3、通过filed有一个Interface()方法得到对应value
	for i := 0; i < inputType.NumField(); i++ {
		filed := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s:%v=%v\n", filed.Name, filed.Type, value)
	}
	//	通过type获取里面的方法调用
	for i := 0; i < inputType.NumMethod(); i++ {
		me := inputType.Method(i)
		fmt.Printf("%s : %v\n", me.Name, me.Type)
	}
}
