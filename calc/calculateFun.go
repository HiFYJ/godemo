package calc

import "fmt"

func init() {
	fmt.Println("calculate init")
}
func Add(a int8, b int8) int8 {
	return a+b
}

func Del(a int8, b int8) int8 {
	return a-b
}

func Mul(a int8, b int8) int8 {
	return a*b
}

func Dev(a int8, b int8) int8 {
	return a/b
}