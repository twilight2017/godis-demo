package main

import (
	"fmt"
)

type processFunc func(int) bool //声明了一个函数类型

//判断元素是否是偶数
func IsEven(integer int) bool {
	if integer%2 == 0 {
		return true
	} else {
		return false
	}
}

//判断元素是否是奇数
func IsOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	} else {
		return true
	}
}

//根据函数来处理切片，根据元素奇数偶数分组，返回新的切片
func filter(slice []int, f processFunc) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice=", slice)
	odd := filter(slice, IsOdd) //把函数当作值进行传递
	fmt.Println("奇数元素是： ", odd)
	even := filter(slice, IsEven) //把函数当成值来传递
	fmt.Println("偶数元素是：", even)
}
