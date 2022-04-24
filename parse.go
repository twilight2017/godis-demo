package main

import (
	"fmt"
	"strconv"
)

func main() {
	TestAtoi()
	TestParseInt()
	TestParseUint()
	TestParseFloat()
	TestParseBool()
}

//将字符串转换为Int型
func TestAtoi() {
	a, _ := strconv.Atoi("100")
	fmt.Printf("%T, %v \n", a, a+2)
	fmt.Println("-------------")
}

//解释给定基数（2到36）的字符串s并返回相应的值i
func TestParseInt() {
	num, _ := strconv.ParseInt("-4e00", 16, 64)
	fmt.Printf("%T, %v \n", num, num)
	num, _ = strconv.ParseInt("01100001", 2, 64)
	fmt.Printf("%T, %v \n", num, num)
	num, _ = strconv.ParseInt("-01100001", 10, 64)
	fmt.Printf("%T, %v \n", num, num)
	num, _ = strconv.ParseInt("4e00", 10, 64)
	fmt.Printf("%T, %v \n", num, num)
	fmt.Println("--------------")
}

//ParseUint类似于ParseInt，但是用于无符号数字
func TestParseUint() {
	num, _ := strconv.ParseUint("4e00", 16, 64)
	fmt.Printf("%T, %v \n", num, num)
	num, _ = strconv.ParseUint("01100001", 2, 64)
	fmt.Printf("%T, %v \n", num, num)
	num, _ = strconv.ParseUint("-1100001", 10, 64)
	fmt.Printf("%T, %v \n", num, num)
	num, _ = strconv.ParseUint("4e00", 10, 64)
	fmt.Printf("%T, %v \n", num, num)
	fmt.Println("--------------")
}

//parseFloat用于将字符串s转换为float型
func TestParseFloat() {
	pi := "3.1415926"
	num, _ := strconv.ParseFloat(pi, 64)
	fmt.Printf("%T, %v \n", num, num*2)
	fmt.Println("--------------")
}

//将字符串转换为bool类型
func TestParseBool() {
	flag, _ := strconv.ParseBool("steven")
	fmt.Printf("%T, %v \n", flag, flag)
	fmt.Println("--------------")
	flag1, _ := strconv.ParseBool("1")
	fmt.Printf("%v \n", flag1, flag1)
}
