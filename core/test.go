package main

import (
	"fmt"
	"strings"
)

func main() {
	result := StringToLower("AghBNJJFkkjhgjFFHG", processCase)
	fmt.Println(result)
	result1 := StringToLower1("FGHJFGJghjhFFHGddeghDFH", processCase)
	fmt.Println(result1)
}

// 奇数偶数依次显示为大小写
func processCase(str string) string {
	result := ""
	for i, value := range str {
		if i%2 == 0 {
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}
	return result
}

func StringToLower(str string, f func(string) string) string {
	fmt.Printf("%T \n", f)
	return f(str)
}

type caseFunc func(string) string

func StringToLower1(str string, f caseFunc) string {
	fmt.Printf("%T \n", f)
	return f(str)

}
