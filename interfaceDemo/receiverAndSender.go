package interfaceDemo

import "fmt"

type Mover interface {
	Move()
}
type Say interface {
	Say()
}

type Wang struct {}

/*值接收*/
func (wang Wang) Move(){
	fmt.Println("wangwang move")
}
/*指针接收*/
func (wang *Wang) Say() {
	fmt.Println("wangwang say")
}