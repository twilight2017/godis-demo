package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(getMultiple(5))
}

//通过递归实现阶乘计算
func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

//通过循环实现阶乘
func getMultiple(num int) (result int) {
	result = 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}
