package main

import (
	"fmt"
	"reflect"
)

/**
使用tag
**/
func main() {
	var re resume
	findTag(&re)
}

type resume struct {
	Name string `info:"name" doc:"名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info is ", tagInfo, " doc is ", tagDoc)
	}
}
