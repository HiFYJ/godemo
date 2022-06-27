package main

import (
	"fmt"
	"os"
)

func main() {

	//Mkdir("./666")

	//MkdirAll("233/666")
	//removeAll("233")
}

/**
	创建单层目录
**/
func Mkdir(name string) error {
	err := os.Mkdir(name, 0666)
	if err != nil {
		fmt.Println("创建目录出错", err)
		return err
	}
	fmt.Println("创建目录成功")
	return nil
}

/**
	创建多层目录
**/
func MkdirAll(path string) error {
	err := os.MkdirAll(path, 06666)
	if err != nil {
		fmt.Println("创建失败:", err)
		return err
	}
	fmt.Println("创建成功")
	return nil
}

/**
	s删除目录
**/
func removeDir(path string) error {
	os.Remove(path)
	return nil
}

/**
	删除目录以及下面的所有
**/
func removeAll(path string) error {
	os.RemoveAll(path)
	return nil
}
