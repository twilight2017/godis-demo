package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "大概顿巴斯的"
	fmt.Println("字节长度", len(s))
	len := 0
	for i, ch := range s {
		fmt.Printf("%d:%c", i, ch)
		len++
	}
	//遍历所有字节
	fmt.Println("\n字符串长度", len)
	for i, ch := range []byte(s) {
		fmt.Printf("%d:%x", i, ch)
	}
	//遍历所有字符
	count := 0
	//rune可以将字符串转化为unicode码点
	for i, ch := range []rune(s) {
		fmt.Printf("%d:%c", i, ch)
		count++
	}
	fmt.Println("字符串长度", count)
	fmt.Println("字符串长度", utf8.RuneCountInString(s))
}
