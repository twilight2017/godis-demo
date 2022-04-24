package main

import (
	"fmt"
	"strconv"
)

func main() {
	TestItoa()
	TestFormatInt()
	TestFormatUint()
	TestFormatFloat()
	TestFormatBool()
}

//Int转换为String
func TestItoa() {
	s := strconv.Itoa(199)
	fmt.Printf("%T, %v, 长度：%d \n", s, s, len(s))
	fmt.Println("--------")
}

//返回给定基数的i的字符串表示
func TestFormatInt() {
	s := strconv.FormatInt(-19968, 16)
	s = strconv.FormatInt(-40869, 16)
	fmt.Printf("%T, %v, 长度：%d \n", s, s, len(s))
	fmt.Println("--------")
}

//返回给定基数i的字符串表示
func TestFormatUint() {
	s := strconv.FormatUint(19968, 16)
	s = strconv.FormatUint(40869, 16)
	fmt.Printf("%T, %v, 长度： %d\n", s, s, len(s))
	fmt.Println("---------")
}

//将浮点数f转换为字符串
func TestFormatFloat() {
	s := strconv.FormatFloat(3.1415926, 'g', -1, 64)
	fmt.Printf("%T, %v, 长度: %d \n", s, s, len(s))
	fmt.Println("----------")
}

func TestFormatBool() {
	s := strconv.FormatBool(true)
	fmt.Printf("%T, %v, 长度： %d \n", s, s, len(s))
	fmt.Println("-----------")
}
