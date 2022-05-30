package main

import "fmt"

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("****cat sleep")
}
func (this *Cat) GetColor() string {
	fmt.Println("****cat color is ", this.color)
	return this.color
}
func (this *Cat) GetType() string {
	return "cat"
}
