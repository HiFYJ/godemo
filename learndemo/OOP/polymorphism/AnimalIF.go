package main

//本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}
