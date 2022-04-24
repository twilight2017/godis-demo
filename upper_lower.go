package main

import (
	"fmt"
	"strings"
)

func main() {
	TestTitle()
	TestToTitle()
	TestToLower()
	TestToUpper()
}

//将字符串s每个单词首字母大写并返回
func TestTitle() {
	fmt.Println(strings.Title("hka sf adfa asdw wer df"))
}

//将字符串s转换为大写并返回
func TestToTitle() {
	fmt.Println(strings.ToTitle("daAdasdzD"))
}

//将字符串s转换为小写并返回
func TestToLower() {
	fmt.Println(strings.ToLower("AdBJHJJKJM"))
}

//将字符串转换为大写返回
func TestToUpper() {
	fmt.Println(strings.ToUpper("Godah"))
}
