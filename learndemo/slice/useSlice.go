package main

import "fmt"

/**
slice切片关键字：
	1、数组：长度固定；声明方式指定长度；传参时严格匹配数组类型；传参时是值拷贝
	2、slice切片：长度动态；传参时是引用传递；
**/
func main() {
	//testArr() //测试固定长度数组
	testArr2() //测试动态数组slice切片

	//userSlice()
	//sliceContent() //测试切片操作
}

/**
切片容量操作追加和截取
**/
func sliceContent() {

	var numbers = make([]int, 3, 5)                                                                     //number length:3,cap:5()尾指针指向的位置
	fmt.Printf("number length:%d,cap:%d,value:%v \n", len(numbers), cap(numbers), numbers)              //number length:3,cap:5,value:[0 0 0]
	numbers = append(numbers, 1)                                                                        //给切片追加一个元素1
	fmt.Printf("after append number length:%d,cap:%d,value:%v \n", len(numbers), cap(numbers), numbers) //after append number length:4,cap:5,value:[0 0 0 1]

	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	//超过cap容量之后，cap容量扩容为2倍
	fmt.Printf("after append number length:%d,cap:%d,value:%v \n", len(numbers), cap(numbers), numbers) //after append number length:6,cap:10,value:[0 0 0 1 2 3]

	var number2 = []int{1, 2, 3, 4, 5}
	child := number2[1:4] //左闭右开 [1,4) 应该截取2,3,4
	for _, value := range child {
		fmt.Print("value:", value, "\t") //value:2 value:3 value:4
	}
	child[0] = 100                     //虽然是截取，但是指向的是一个内存空间，所以修改会改变原数组的值
	fmt.Println("\nnumber2:", number2) //number2: [1 100 3 4 5]

	//	解决上述会修改原数组方式
	var childCp = make([]int, 3)
	copy(childCp, child)
	childCp[0] = 200                   //改变不会影响原数组，深拷贝
	fmt.Println("\nchildCp:", childCp) //childCp: [200 3 4]
	fmt.Println("\nnumber2:", number2) //number2: [1 100 3 4 5]
}

/**
slice声明方式
**/
func userSlice() {
	//	第一种
	slice1 := []int{1, 2, 3, 4}
	fmt.Printf("slice1 length: %d,type:%T,value:%v\n", len(slice1), slice1, slice1) //slice1 length: 4,type:[]int,value:[1 2 3 4]
	//	第二种 声明但是不分配空间
	var slice2 []int                                                                            //不指定长度
	fmt.Printf("before make slice2 length: %d,type:%T,value:%v\n", len(slice2), slice2, slice2) //before make slice2 length: 0,type:[]int,value:[]
	slice2 = make([]int, 3)                                                                     //分配空间
	fmt.Printf("after make slice2 length: %d,type:%T,value:%v\n", len(slice2), slice2, slice2)  //after make slice2 length: 3,type:[]int,value:[0 0 0]
	//	第三种 声明的同时分配空间
	var slice3 = make([]int, 3)
	fmt.Printf("slice3 length: %d,type:%T,value:%v\n", len(slice3), slice3, slice3) //slice3 length: 3,type:[]int,value:[0 0 0]
	//	第四种 第三种的简写
	slice4 := make([]int, 3)
	fmt.Printf("slice4 length: %d,type:%T,value:%v\n", len(slice4), slice4, slice4) //slice4 length: 3,type:[]int,value:[0 0 0]

	//	注：判断是否为空
	if slice4 == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("not is nill")
	}
}

/**
使用slice时是引用传递，修改的值会影响原始数组
**/
func printSlice(arr []int) {
	for _, v := range arr {
		fmt.Println("v:", v)
	}
	arr[1] = 100
}

//动态长度数组 切片slice
func testArr2() {
	myarr := []int{1, 2, 3, 4, 5}
	fmt.Printf("myarr type=%T \n", myarr) //这里输出类型没有长度 myarr type=[]int
	printSlice(myarr)
	fmt.Println("-----------")
	for index, value := range myarr {
		fmt.Println("index:", index, "value:", value)
	}
}

/**
这里printArr的参数实际还是 值拷贝 ，对于原数组的修改没有影响
**/
func printArr(arr [4]int) {
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}
	arr[1] = 100
}

//固定长度数组
func testArr() {
	var myarr1 [10]int
	myarr2 := [10]int{1, 2, 3, 4, 5}
	myarr3 := [4]int{1, 2, 3, 4}

	for i := 0; i < len(myarr1); i++ {
		fmt.Println(myarr1[i])
	}

	for index, value := range myarr2 {
		fmt.Println("index:", index, "value:", value)
	}

	fmt.Printf("myarr1 type = %T \n", myarr1)
	fmt.Printf("myarr2 type = %T \n", myarr2)
	fmt.Printf("myarr3 type = %T \n", myarr3)
	fmt.Println("--------------")
	printArr(myarr3)
}
