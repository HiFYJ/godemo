package structDemo

import (
	"encoding/json"
	"fmt"
)

type User struct {
	name string
	sex int8
	city string
}
func StructDemo() {
	var me User
	me.city = "南京"
	me.name= "Jack"
	me.sex =2
	fmt.Printf("me:%v\n",me)
}

/*匿名结构体*/
func StructDemo2() {
	var me struct{
		name string
		sex int8
		city string
	}
	me.city = "南京"
	me.name= "Jack love"
	me.sex =1
	fmt.Printf("me:%v\n",me)
}

/*创建指针类型结构体*/
func StructDemo3() {
	var me = new (User)
	me.city = "南京"
	me.name= "Jack ma"
	me.sex =1
	fmt.Printf("me:%#v\n",me)
}

/*构造方法*/
func NewUser(name string,sex int8,city string) *User  {
	return &User{
		name: name,
		sex: sex,
		city: city,
	}
}

func (user User) Introduce()  {
	fmt.Printf("I'm :%v\n",user)
}


type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func JsonDemo() {
	s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "沙河娜扎",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"}
}