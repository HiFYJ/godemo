package main

import "time"

type GormPage struct {
	Id         int `xorm:"pk autoincr"`
	Titile     string
	Content    string    `xorm:"text"`
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

/**
	修改默认表名
**/
/*func (gormPage *GormPage)TableName()string{
	return "test_tb"
}*/
