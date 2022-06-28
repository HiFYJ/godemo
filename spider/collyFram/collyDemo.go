package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	//设置属性
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36"
	//方法一：设置请求头
	/*c.OnRequest(func(request *colly.Request) {
		request.Headers.Add("cookie","")
	})*/
	//方法二：自动设置cookie，根据url拿到cookie
	/*cookie := c.Cookies("url")
	c.SetCookies("",cookie)*/

	//回调方法：当请求返回html的时候执行该方法，***重点****
	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		s := e.Attr("href")
		e.Request.Visit(s)
		fmt.Printf("标签名称：%v 标签文本：%v\n", e.Name, e.Text)
		url := e.Request.AbsoluteURL(s)
		fmt.Printf("请求绝对路径:%v\n", url)
	})

	//回调方法：当发生请求时调用该方法
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("url is: %v \n", r.URL)
	})

	c.OnResponse(func(re *colly.Response) {
		//fmt.Printf("response recived:%v \n",string(re.Body))
		fmt.Printf("response recived.. \n")
	})
	c.OnError(func(res *colly.Response, err error) {
		fmt.Printf("has error:%v \n", err)
	})
	//回调：每次请求完成之后
	c.OnScraped(func(response *colly.Response) {
		fmt.Println("spider finished...")
	})

	//发送一个请求，注意要放在回调方法后面
	c.Visit("https://gorm.io/zh_CN/docs/")
}
