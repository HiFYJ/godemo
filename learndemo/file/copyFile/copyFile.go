package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//err := CopyFile("./view.text","./test2.text")

	err := BufIOCopy("./view.text", "./test2.text")

	//IOUtileCopy("./view.text","./test.text")
	if err != nil {
		fmt.Println("复制文件出错：", err)
	}
}

/**
	先通过read读文件
	通过Write方法写文件
**/
func CopyFile(sourceFile, target string) error {
	//一、打开文件
	source, err := os.Open(sourceFile)
	distination, werr := os.OpenFile(target, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil || werr != nil {
		fmt.Println("openfile error")
		return err
	}
	tempSlice := make([]byte, 128)
	for {
		//读文件
		n, er := source.Read(tempSlice)
		if er == io.EOF {
			fmt.Println("读取文件结束")
			break
		}
		if er != nil {
			fmt.Println("读取文件出错")
			return er
		}
		fmt.Println("读取文件大小：", n)

		//写文件
		n, e := distination.Write(tempSlice[:n])
		if e != nil {
			fmt.Println("write error")
			return e
		}
		fmt.Println("写入", n, "个字节\n")
	}
	return nil
}

/**
  	使用bufio写文件
	注意最后需要flush从缓存写到文件
**/
func BufIOCopy(sourceFile, target string) error {
	//一、打开文件
	source, serr := os.Open(sourceFile)
	tar, terr := os.OpenFile(target, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer source.Close()
	defer tar.Close()
	if serr != nil {
		fmt.Println("read open file error")
		return serr
	}
	if terr != nil {
		fmt.Println("write openfile error")
		return terr
	}
	reader := bufio.NewReader(source)
	writer := bufio.NewWriter(tar)
	for {
		str, er := reader.ReadString('\n')
		if er == io.EOF {
			fmt.Println("ReadString finish")
			//TODO 注意这里最后如果有换行会有问题 内容不完整
			writer.WriteString(str)
			break
		}
		if er != nil {
			fmt.Println("ReadString error")
			return er
		}
		writer.WriteString(str)
	}

	writer.Flush()
	return nil
}

/**
	默认是覆盖
**/
func IOUtileCopy(sourceFile, target string) error {
	fileByte, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println("read file error")
		return err
	}

	er := ioutil.WriteFile(target, fileByte, 0666)
	if er != nil {
		fmt.Println("write file  error")
		return er
	}
	return nil
}
