package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//WriteFile("./test.text")
	//BufIOWrite("./test.text")
	//IOUtileWrite("./test.text")
	SaveFile("./test627.text")
}

/**
	通过Write方法写文件
**/
func WriteFile(path string) {
	//一、打开文件
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("error")
		return
	}
	file.Write([]byte("msg:\n"))
	file.WriteString("helloworld")
}

/**
  	使用bufio写文件
	注意最后需要flush从缓存写到文件
**/
func BufIOWrite(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("error")
		return
	}
	writer := bufio.NewWriter(file)
	writer.WriteString("hello world \n")
	writer.Write([]byte("你好中国！"))
	writer.Flush()
}

/**
	默认是覆盖
**/
func IOUtileWrite(path string) {
	err := ioutil.WriteFile(path, []byte("你好golang233"), 0666)
	if err != nil {
		fmt.Println("ioutil error")
		return
	}
}

/**
	使用os.WriteFile创建并写入数据
**/
func SaveFile(path string) {
	err := os.WriteFile(path, []byte("你好os"), 0666)
	if err != nil {
		fmt.Println("writefile error", err)
		return
	}
}
