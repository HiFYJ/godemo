package main

import (
	"os"
)

func main() {

	RenameFile("./test2.text", "./test002.text")
}

/**
	修改文件名称
**/
func RenameFile(source, target string) error {
	err := os.Rename(source, target)
	if err != nil {
		return err
	}
	return nil
}
