package main

import "fmt"

//具体的类 狗
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("****dog sleep")
}
func (this *Dog) GetColor() string {
	fmt.Println("****dog color is ", this.color)
	return this.color
}
func (this *Dog) GetType() string {
	return "dog"
}
