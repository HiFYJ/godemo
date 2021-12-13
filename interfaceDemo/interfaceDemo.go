package interfaceDemo

import "fmt"

type Sayer interface {
	say()
}

type Dog struct {}
type Cat struct {}

func (d Dog) Say()  {
	fmt.Println("一条狗,汪汪汪")
}

func (c Cat) Say()  {
	fmt.Println("一只猫,喵喵喵")
}

func Demo() {
	var d Dog
	d.Say()
	var c Cat
	c.Say()
}