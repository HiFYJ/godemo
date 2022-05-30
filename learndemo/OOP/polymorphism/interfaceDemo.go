package main

import "fmt"

/**
 多态概念
 接口interface
 基本要素：
	有一个父类（接口）
	有子类（实现父类的所有接口）
	父类类型的变量（指针）指向（引用）子类的具体数据变量
**/
func main() {
	//方式一
	/*var animal AnimalIF
	animal = &Cat{"blue"}
	animal.GetColor()
	animal.GetType()
	animal.Sleep()

	animal = &Dog{"black"}
	animal.GetColor()
	animal.GetType()
	animal.Sleep()*/

	//	方式二
	cat := Cat{"yellow"}
	dog := Dog{"red"}

	showAnimal(&cat)
	showAnimal(&dog)

}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("type =", animal.GetType())
	fmt.Println("color =", animal.GetColor())
}

//本质是一个指针
//type AnimalIF interface {
//	Sleep()
//	GetColor() string
//	GetType() string
//}

//具体的类 猫
//type Cat struct {
//	color string
//}
//func (this *Cat) Sleep() {
//	fmt.Println("****cat sleep")
//}
//func (this *Cat) GetColor()string {
//	fmt.Println("****cat color is ",this.color)
//	return this.color
//}
//func (this *Cat) GetType()string {
//	return "cat"
//}

//具体的类 狗
//type Dog struct {
//	color string
//}
//func (this *Dog) Sleep() {
//	fmt.Println("****dog sleep")
//}
//func (this *Dog) GetColor()string {
//	fmt.Println("****dog color is ",this.color)
//	return this.color
//}
//func (this *Dog) GetType()string {
//	return "dog"
//}
