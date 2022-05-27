package main

import "fmt"

/**
map 使用
**/
func main() {
	//createMap()
	userMap()
}

/**
 map使用
**/
func userMap() {
	mymap := make(map[string]string)
	//	添加
	mymap["user"] = "李四"
	mymap["age"] = "26"
	mymap["sex"] = "man"
	mymap["address"] = "nanjing"
	//	遍历
	printMap(mymap)
	//	删除
	delete(mymap, "address")
	//	修改
	mymap["sex"] = "男"
	fmt.Println("-------------")
	printMap(mymap)
}
func printMap(mymap map[string]string) {
	//这里的传参是引用传递
	for k, v := range mymap {
		fmt.Println("key:", k, "\tvalue:", v)
	}
	mymap["age"] = "88" //第二次调用会显示修改后的值
}

/**
 map声明方式
**/
func createMap() {
	//	第一种方式 先声明 使用前分配空间
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("初始map为空")
	}
	myMap = make(map[string]string, 5)
	myMap["one"] = "hello"
	myMap["two"] = "world"
	myMap["three"] = "!"
	fmt.Println("map:", myMap)
	//	第二种方式 声明同时分配空间
	mymap2 := make(map[int]string)
	mymap2[0] = "你好"
	mymap2[1] = "中国"
	mymap2[2] = "boy"
	fmt.Println("map2:", mymap2)
	//第三种方式
	mymap3 := map[string]string{
		"user": "张三",
		"age":  "26",
		"sex":  "男",
	}
	fmt.Println("map3:", mymap3)
	fmt.Println("---------------")
	//	遍历
	for k, v := range mymap3 {
		fmt.Println("key:", k, "\tvalue:", v)
	}
}
