package main

import (
	"fmt"
	"net/http"
)

func main() {
	testClientGet()
}

func testClientGet() {
	//创建客户端
	client := http.Client{}
	//通过client发起请求
	response, err := client.Get("https://www.baidu.com/")
	CheckErr(err)
	fmt.Printf("响应状态码： %v \n", response.StatusCode)
	if response.StatusCode == 200 {
		fmt.Println("request network success")
		defer response.Body.Close()
	}
}

//CheckErr
func CheckErr(err error) {
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("program has a panic", ins.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}
