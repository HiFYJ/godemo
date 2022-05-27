package main

import "fmt"

type Human struct {
	Name string
	Sex  string
}

func (this *Human) Print() {
	fmt.Println("*****输出人human属性：")
	fmt.Println("name:", this.Name)
	fmt.Println("sex:", this.Sex)
}
func (this *SuperMan) print() {
	fmt.Println("*****输出人superMan属性：")
	fmt.Println("name:", this.Name)
	fmt.Println("sex:", this.Sex)
	fmt.Println("level:", this.level)
}
func (this *Human) Walk() {
	fmt.Println("***human walk()")
}
func (this *Human) Eat() {
	fmt.Println("***human eat()")
}
func (this *SuperMan) Walk() {
	fmt.Println("***SuperMan walk()")
}
func (this *SuperMan) Talk() {
	fmt.Println("***SuperMan talk()")
}

type SuperMan struct {
	Human
	level int
}

func main() {
	/* output
	***human walk()
	***human eat()
	**************
	***human eat()
	***SuperMan walk()
	***SuperMan talk()
	 */
	var hu = Human{"zhang3", "wuman"}
	hu.Walk()
	hu.Eat()
	hu.Print()
	fmt.Println("**************")

	//声明并初始化子类方式一
	/*su:= SuperMan{
		Human{
			"钢铁侠","man",
		},100,
	}*/
	//方式二
	var su SuperMan
	su.Name = "绿巨人"
	su.Sex = "男人"
	su.level = 80

	su.Eat()  //子类没有重写，调用父类的
	su.Walk() //子类重写，调用子类
	su.Talk() //子类的方法

	su.Print()
}
