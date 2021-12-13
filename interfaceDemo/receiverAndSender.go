package interfaceDemo

import "fmt"

type Mover interface {
	Move()
	Say()
}

type Wang struct {}

func (wang Wang) Move(){
	fmt.Println("wangwang move")
}
func (wang Wang) Say() {
	fmt.Println("wangwang say")
}