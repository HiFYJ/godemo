package main

import "fmt"

/**
关闭channel
特点：
	1、channel不像文件一样需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显示的结束rang循环之类的，才去关闭channel
	2、关闭channel后，无法向channel再发送数据
	3、关闭channel后，可以继续从channel接收数据
	4、关闭nil channel，无论收发都会被阻塞
**/
func main() {
	c := make(chan int)

	go func() {
		defer close(c)
		fmt.Println(" go程end")
		for i := 0; i < 5; i++ {
			c <- i
		}

		//close(c)//关闭channel
	}()

	//方案一：for循环遍历
	/*for {
		//ok如果为true表示channel没有关闭，如果为false表示channel已经关闭
		if value,ok:=<-c;ok{
			fmt.Println("value is ",value)
		}else {
			break
		}
	}*/

	//range遍历
	for data := range c {
		fmt.Println("value is ", data)
	}

	fmt.Println("main finished....")
	/** output:
	value is  0
	value is  1
	value is  2
	value is  3
	value is  4
	main finished....
	**/
}
