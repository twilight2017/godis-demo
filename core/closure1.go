package main

import "fmt"

func main() {
	myfunc := Counter()
	fmt.Println("myfunc", myfunc)
	/*调用myfuc函数，变量i自增1后并返回*/
	fmt.Println(myfunc())
	fmt.Println(myfunc())
	fmt.Println(myfunc())
	/*创建新的函数nextNumber1，并查看结果*/
	myfunc1 := Counter()
	fmt.Println("myfunc1", myfunc1)
	fmt.Println(myfunc1())
	fmt.Println(myfunc1())
}

//计数器 闭包函数
func Counter() func() int {
	i := 0
	res := func() int {
		i += 1
		return i
	}
	fmt.Println("Counter中的内部函数：", res)
	return res
}
