package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	TestFields()
	TestFieldsFunc()
	TestSplitAfterN()
	TestSplit()
	TestSplitN()
	TestSplitAfter()
}

//将字符串以空白字符分割，并返回一个切片
func TestFields() {
	fmt.Println(strings.Fields("  abd 123 AVD xy hk"))
}

//以字符串满足f(r)==true的字符分割，返回一个切片
func TestFieldsFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.FieldsFunc("abc$asd%ABC&xyz%XYZ", f))
}

//将字符串以sep作为分隔符进行分割，分割后字符去掉sep
func TestSplit() {
	fmt.Printf("%q\n", strings.Split("a, b, c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}

//将字符串s以sep作为分隔符进行分割，分割后字符最后附上sep， n决定返回的切片数
func TestSplitN() {
	fmt.Printf("%q\n", strings.SplitN("a, b, c", ",", 2))
	fmt.Printf("%q\n", strings.SplitN("a, b, c", ",", 1))
}

//将字符s以sep作为分割符进行分割，分割后字符最后附上sep
func TestSplitAfter() {
	fmt.Printf("%q\n", strings.SplitAfter("a, b, c", ","))
}

//将字符串s以sep作为分隔符进行分割，分割后字符最后附上sep，n决定返回的切片数
func TestSplitAfterN() {
	fmt.Printf("%q\n", strings.SplitAfterN("a, b, c", ",", 2))
}
