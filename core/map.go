package main

import "fmt"

func main() {
	//1.声明时同时初始化
	var country = map[string]string{
		"China":  "Beijing",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
		"France": "Paris",
		"Italy":  "Rome",
	}
	if value, ok := country["USA"]; ok {
		fmt.Println("首都", value)
	} else {
		fmt.Println("首都未检测到")
	}
	fmt.Println(country)
	//短变量声明初始化方法
	rating := map[string]float64{"c": 5, "Go": 4.5, "Python": 4.5, "C++": 3}
	fmt.Println(rating)
	//2.创建map后再复制
	countryMap := make(map[string]string)
	countryMap["China"] = "Beijing"
	countryMap["Japan"] = "Tokyo"
	countryMap["India"] = "New Delhi"
	countryMap["France"] = "Paris"
	countryMap["Italy"] = "Rome"
	//3.遍历map(无序)
	for k, v := range countryMap {
		fmt.Println("国家", k, "首都", v)
	}
}
