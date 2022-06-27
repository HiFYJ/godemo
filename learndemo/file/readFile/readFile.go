package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//file := ReadFile("./view.text")
	//fmt.Println(string(file))

	//BufIORead("./view.text")

	IOUtileRead("./view.text")
}

/**
	通过Read方法读取文件
	注意最后的切片拼接的问题
**/
func ReadFile(path string) []byte {
	//一、打开文件
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println("openfile error")
		return nil
	}
	//fmt.Println(file) //&{0xc000122780}
	//	二、读取文件内容
	var fileSlice []byte
	tempSlice := make([]byte, 128)
	for {
		n, er := file.Read(tempSlice)
		if er == io.EOF {
			fmt.Println("读取文件结束")
			break
		}
		if er != nil {
			fmt.Println("读取文件出错")
			return nil
		}
		fmt.Println("读取文件大小：", n)
		//最后不一定够128个字节，所以有多少字节读取多少字节
		fileSlice = append(fileSlice, tempSlice[:n]...)
	}

	//fmt.Println("切片：",string(fileSlice))
	return fileSlice
}

/**
  	使用bufio流读取文件
	注意处理最后结束时的情况，内容可能会丢失（bufio读取文件最后出现EOF错误可能还会返回str）
**/
func BufIORead(path string) {
	//一、打开文件
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println("openfile error")
		return
	}
	reader := bufio.NewReader(file)
	var fileStr string
	for {
		str, er := reader.ReadString('\n')
		if er == io.EOF {
			fmt.Println("ReadString finish")
			//TODO 注意这里最后如果有换行会有问题 内容不完整
			fileStr += str
			break
		}
		if er != nil {
			fmt.Println("ReadString error")
			return
		}
		fileStr += str
	}

	fmt.Println(fileStr)
}

/**
	不同于前两者这个方法是一次性就读出数据
	适合较小的文件
**/
func IOUtileRead(path string) {
	fileByte, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("ioutil error")
		return
	}
	fmt.Println(string(fileByte))
}
