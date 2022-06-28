package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	body := Fetch("https://gorm.io/zh_CN/docs/")
	Parse(body)
}

func Fetch(url string) string {
	//创建请求
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")
	req.Header.Add("Cookie", "BIDUPSID=FC89240F9FB90D6212051B9534E1B5DF; PSTM=1627539380; __yjs_duid=1_57e0408ad622192cfdac7115504b00281627550884177; BAIDUID=2001FEA6156A0353AC92DC507F2C0A7F:SL=0:NR=10:FG=1; BDSFRCVID=KGuOJexroG0C3dcD6ufkUJX_ivGM4w5TDYLEOwXPsp3LGJLVgxPTEG0PtjJ5HU4bLrA9ogKKL2OTHm_F_2uxOjjg8UtVJeC6EG0Ptf8g0M5; H_BDCLCKID_SF=tRk8oK-atDvDqTrP-trf5DCShUFsWb-OB2Q-XPoO3KJCqP8GyMTl3MuAyR790tQiWbRM2MbgylRp8P3y0bb2DUA1y4vpKh5GygTxoUJ2XhrPDfcoqtnWhfkebPRiWTj9QgbLMhQ7tt5W8ncFbT7l5hKpbt-q0x-jLTnhVn0MBCK0hD0wDT8hD6PVKgTa54cbb4o2WbCQ3ncz8pcN2b5oQT8wKp-tBU6PQgJPafnv3J5vOIJTXpOUWfAkXpJvQnJjt2JxaqRC5MQIoh5jDh3MBUCzhJ6ie4ROamby0hvcBIocShnzBUjrDRLbXU6BK5vPbNcZ0l8K3l02V-bIe-t2XjQhDHt8J50ttJ3aQ5rtKRTffjrnhPF35PvDXP6-hnjy3bRNQ-8aWprlSKOayl7cQ4LUyN3MWh3RymJ42-39LPO2hpRjyxv4bpKYW-oxJpOJX2olWxccHR7WDqnvbURvX5Dg3-7LyM5dtjTO2bc_5KnlfMQ_bf--QfbQ0hOhqP-jBRIEoC0XtK-hhCvPKITD-tFO5eT22-usatjl2hcHMPoosIJ-Kn-hyMCA3hQPBqv93DTiaKJjBMbUoqRHXnJi0btQDPvxBf7p5acGhp5TtUJM_prDhq6vqt4bht7yKMniBnr9-pnEXhQrh459XP68bTkA5bjZKxtq3mkjbPbDfn028DKuD6-2j63yDNts5JtXKD600PK8Kb7Vbpvj5fnkbJkXhPtj35JbK2PqLRQoQlci_nTCQUFM5TL7Qbrr0xRfyNReQIO13hcdSR390hOpQT8r5M6ZtxvPJIrrQqDXab3vOp44XpO1hq4zBN5thURB2DkO-4bCWJ5TMl5jDh3Mb6ksD-FtqtJHKbDe_CL-JMK; BDSFRCVID_BFESS=KGuOJexroG0C3dcD6ufkUJX_ivGM4w5TDYLEOwXPsp3LGJLVgxPTEG0PtjJ5HU4bLrA9ogKKL2OTHm_F_2uxOjjg8UtVJeC6EG0Ptf8g0M5; H_BDCLCKID_SF_BFESS=tRk8oK-atDvDqTrP-trf5DCShUFsWb-OB2Q-XPoO3KJCqP8GyMTl3MuAyR790tQiWbRM2MbgylRp8P3y0bb2DUA1y4vpKh5GygTxoUJ2XhrPDfcoqtnWhfkebPRiWTj9QgbLMhQ7tt5W8ncFbT7l5hKpbt-q0x-jLTnhVn0MBCK0hD0wDT8hD6PVKgTa54cbb4o2WbCQ3ncz8pcN2b5oQT8wKp-tBU6PQgJPafnv3J5vOIJTXpOUWfAkXpJvQnJjt2JxaqRC5MQIoh5jDh3MBUCzhJ6ie4ROamby0hvcBIocShnzBUjrDRLbXU6BK5vPbNcZ0l8K3l02V-bIe-t2XjQhDHt8J50ttJ3aQ5rtKRTffjrnhPF35PvDXP6-hnjy3bRNQ-8aWprlSKOayl7cQ4LUyN3MWh3RymJ42-39LPO2hpRjyxv4bpKYW-oxJpOJX2olWxccHR7WDqnvbURvX5Dg3-7LyM5dtjTO2bc_5KnlfMQ_bf--QfbQ0hOhqP-jBRIEoC0XtK-hhCvPKITD-tFO5eT22-usatjl2hcHMPoosIJ-Kn-hyMCA3hQPBqv93DTiaKJjBMbUoqRHXnJi0btQDPvxBf7p5acGhp5TtUJM_prDhq6vqt4bht7yKMniBnr9-pnEXhQrh459XP68bTkA5bjZKxtq3mkjbPbDfn028DKuD6-2j63yDNts5JtXKD600PK8Kb7Vbpvj5fnkbJkXhPtj35JbK2PqLRQoQlci_nTCQUFM5TL7Qbrr0xRfyNReQIO13hcdSR390hOpQT8r5M6ZtxvPJIrrQqDXab3vOp44XpO1hq4zBN5thURB2DkO-4bCWJ5TMl5jDh3Mb6ksD-FtqtJHKbDe_CL-JMK; delPer=0; PSINO=5; ZFY=VwLPDF1yE5vIxYSAK54lhoPTl2diiDyoPYiVbp5fKCc:C; BAIDUID_BFESS=B2A6685CC83D22ACA4208857A184DE49:FG=1; MCITY=-:; BDUSS=mtEenNYeW45UXV2WUFWSDZPSlJlUVBVWmYtYzFIWVZsQjk3YXJUVlBVTnNxLUJpSVFBQUFBJCQAAAAAAAAAAAEAAABwH~Ap0KHQobfraGVsbG8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGweuWJsHrlia; BDUSS_BFESS=mtEenNYeW45UXV2WUFWSDZPSlJlUVBVWmYtYzFIWVZsQjk3YXJUVlBVTnNxLUJpSVFBQUFBJCQAAAAAAAAAAAEAAABwH~Ap0KHQobfraGVsbG8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGweuWJsHrlia; BA_HECTOR=84810h80a18k2k20a41hbi8tt15; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; H_PS_PSSID=36549_36459_36673_36455_31253_36452_36692_36167_36693_36696_36653_36658_26350_36468; ab_sr=1.0.1_NmNmZjRjNDZhYzY5Yjc5YWIyODFjNzE5NmYxNzhhNzYwZjkyYTk5YzNlNzJhMDQwZjMwNWM4MzMyOWFhM2I3YTBhOWE2N2I1ZjhhMTcxMDM1ZmRjZGY0NjhiM2MyNTFlYjZlZjUxZjIyZjAxNmViNzkxOTRkZjEyZGM5YzI3YTA2ZmUzNzQ1MGYwYzExNDk0ZDYxOWY1MWFiN2ZhZWE5Yw==")
	rep, err := client.Do(req)
	if err != nil {
		fmt.Println("http request error")
		return ""
	}
	if rep.StatusCode != 200 {
		fmt.Println("http status code:", rep.StatusCode)
		return ""
	}
	defer rep.Body.Close()

	body, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		fmt.Println("read body error :", err)
		return ""
	}
	//fmt.Println(string(body))
	return string(body)

}
func Parse(html string) {
	//替换掉空格
	html = strings.Replace(html, "\n", "", -1)
	//边栏内容块正则
	re_sidebar := regexp.MustCompile(`<aside id="sidebar" role="navigation">(.*?)</aside>`)
	//找到边栏内容块
	sidebar := re_sidebar.FindString(html)
	//链接正则
	re_link := regexp.MustCompile(`href="(.*?)"`)
	//找到所有链接
	links := re_link.FindAllString(sidebar, -1)
	base_url := "https://gorm.io/zh_CN/docs/"
	for _, v := range links {
		fmt.Printf("v: %v \n", v)
		s := v[6 : len(v)-1]
		url := base_url + s
		//url = strings.ReplaceAll(url,"//","/")
		fmt.Printf("url: %v \n", url)
		body := Fetch(url)
		//fmt.Println(body)
		//启动另外一个线程处理go
		go Parse2(body)
	}

}
func Parse2(body string) {
	//替换页面空格
	body = strings.ReplaceAll(body, "\n", "")
	//页面内容
	re_content := regexp.MustCompile(`<div class="article">(.*?)</div>`)
	//找到页面内容
	content := re_content.FindString(body)

	//标题
	re_title := regexp.MustCompile(`<h1 class="article-title" itemprop="name">(.*?)</h1>`)
	//找到页面内容
	title := re_title.FindString(content)
	fmt.Println("title is :", title)

	title = title[42 : len(title)-5]
	fmt.Printf("title replace after : %v \n", title)
}
