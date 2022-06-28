package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine *xorm.Engine
var err error

func init() {
	Engine, err = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/xorm?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接失败：", err)
		return
	}
	//连接测试
	if err := Engine.Ping(); err != nil {
		fmt.Println("测试数据库连接失败：", err)
		return
	}
	//defer Engine.Close()
	fmt.Println("数据库连接成功。。。。")
	Engine.ShowSQL(true)
	Engine.SetMaxIdleConns(20) //连接池空闲数大小
	Engine.SetMaxOpenConns(20) //最大连接数

}

func main() {
	//Engine.Query("select * from user;")
	SaveToDB("活着", "这是一本书，作者余华")
}

func SaveToDB(title, content string) {

	Engine.Sync2(new(GormPage)) //创建表
	page := GormPage{
		Titile:  title,
		Content: content,
	}
	affect, err := Engine.Insert(&page)
	if err != nil {
		fmt.Println("插入数据失败：", err)
		return
	}

	fmt.Printf("save data:%d条\n", affect)
	fmt.Printf("save data:%v\n", page)
}
