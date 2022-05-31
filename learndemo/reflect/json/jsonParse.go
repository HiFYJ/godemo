package main

import (
	"encoding/json"
	"fmt"
)

/**
json 使用
使用标签可以进行字段重命名
**/
func main() {
	movie := Movie{Actors: []string{"zbz", "xy"}, Title: "xjzw", Year: 2000, Price: 15}

	//一、结构体转成json字符串
	str, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error:", err)
	}
	fmt.Printf("***%s \n", str) //***{"title":"xjzw","year":2000,"rmb":15,"actors":["zbz","xy"]}

	//二、json转成结构体
	movie2 := Movie{}
	err2 := json.Unmarshal(str, &movie2)
	if err2 != nil {
		fmt.Println("unmarshal error:", err2)
	}
	fmt.Println("movie2 is", movie2)
}

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

/*type Movie struct {
	Title  string
	Year   int
	Price  int
	Actors []string
}*/
