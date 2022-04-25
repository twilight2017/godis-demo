package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	testHttpPost()
}

func testHttpPost() {
	//construct param
	data := url.Values{
		"theCityName": {"重庆"},
	}

	//参数转化为body
	fmt.Println("data.Encode value:", data.Encode())
	reader := strings.NewReader(data.Encode())
	//发起post请求
	response, err := http.Post("http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName",
		"application/x-www-form-urlencoded", reader)
	CheckErr(err)
	fmt.Printf("响应状态码： %v \n", response.StatusCode)
	if response.StatusCode == 200 {
		defer response.Body.Close()
		fmt.Println("request network success")
		CheckErr(err)
	} else {
		fmt.Println("request network failed", response.Status)
	}

}

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
