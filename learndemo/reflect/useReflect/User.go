package main

import "fmt"

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) CallMethod() {
	fmt.Println("user is call")
	fmt.Printf("%v\n", this)
}
