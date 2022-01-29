package interfaceDemo

import "fmt"

/**
多个类型实现同一接口
并且一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现
**/

type WashingMachine interface {
	Wash()
	Dry()
}
//烘干机
type Dryer struct {}

func (d Dryer) Dry() {
	fmt.Println("烘干了")
}
//洗衣机
type Haier struct {
	Dryer //嵌入烘干机
}

func (h Haier) Wash() {
	fmt.Println("洗干净了")
}
