package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	TestTrim()
	TestTrimFunc()
	TestTrimLeft()
	TestTrimLeftFunc()
	TestTrimRight()
	TestTrimRightFunc()
	TestTrimSpace()
	TestTrimPrefix()
	TestTrimSuffix()
}

//将字符串s首尾包含在cutset中的任一字符去掉并返回
func TestTrim() {
	fmt.Println(strings.Trim(" steven wang ", " "))
}

//将字符串s首位满足函数f(r)==true的字符去掉返回
func TestTrimFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.TrimFunc("!jk%^*)(^", f))
}

//将字符串s左边包含在cutset中的任一字符去掉返回
func TestTrimLeft() {
	fmt.Println(strings.TrimLeft("  steven wang  ", " "))
}

//将字符串s左边满足函数f(r)==true的字符去掉返回
func TestTrimLeftFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.TrimLeftFunc("$%*(fhjkjhghjklkh", f))
}

//将字符串s右边包含在cutset中的任一字符去掉返回
func TestTrimRight() {
	fmt.Println(strings.TrimRight("  steven wang  ", " "))
}

//将字符串s右边满足函数f(r)==true的字符去掉返回
func TestTrimRightFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.TrimRightFunc("$%*(fhjkjhghjklkh", f))
}

//将字符串s首位空白去掉并返回
func TestTrimSpace() {
	fmt.Println(strings.TrimSpace(" dauhdis  "))
}

//将字符串s中前缀字符串prefix去掉并返回
func TestTrimPrefix() {
	var s = "dahdkkjidgauidjo"
	s = strings.TrimPrefix(s, "dahd")
	fmt.Println(s)
}

//将字符串s中后缀字符串suffix去掉并返回
func TestTrimSuffix() {
	var s = "dagdhuohjgiuh"
	s = strings.TrimSuffix(s, "giuh")
	fmt.Println(s)
}
