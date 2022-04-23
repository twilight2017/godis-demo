package main

import "fmt"

func main() {
	//声明实际变量
	var a int = 110
	//声明指针变量
	var ip *int
	//给指针变量赋值，把变量a的地址赋给指针变量ip
	ip = &a
	//打印变量a的类型和值
	fmt.Printf("a的类型是%T,值是%v\n", a, a)
	//打印&a的类型和值
	fmt.Printf("&a的类型和值是%T,值是%v\n", &a, &a)
	//打印ip的类型和值
	fmt.Printf("ip的类型是%T,值是%v\n", ip, ip)
	//打印变量*ip的类型和值,*ip可以获取指针指向的变量值
	fmt.Printf("*ip的类型是%T， 值是%v\n", *ip, *ip)
}
