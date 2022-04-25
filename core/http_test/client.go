package main

import (
	"fmt"
	"net/http"
)

func main() {
	testHttpNewRequest()
}

func testHttpNewRequest() {
	//1.创建一个客户端
	client := http.Client{}
	//2.创建一个请求，请求方式既可以是Get，也可以是POST
	request, err := http.NewRequest("GET", "https://www.baidu.com/", nil)
	CheckErr(err)
	//3.客户端发送请求
	cookName := &http.Cookie{Name: "username", Value: "Steven"}
	//添加cookie
	request.AddCookie(cookName)
	response, err := client.Do(request)
	CheckErr(err)
	//设置请求头
	request.Header.Set("Accept-Language", "zh-cn")
	defer response.Body.Close()
	//查看请求头的数据
	fmt.Printf("Header:%+v\n", request.Header)
	fmt.Printf("响应状态码： %v\n", response.StatusCode)
	//操作数据
	if response.StatusCode == 200 {
		fmt.Println("request network success")
		CheckErr(err)
	} else {
		fmt.Println("request network failed", response.Status)
	}
}

//CheckErr
func CheckErr(err error) {
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("program has an error", ins.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}
