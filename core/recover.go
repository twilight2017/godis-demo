package main

import "fmt"

func main() {
	funcA()
	funcB()
	funcC()
	fmt.Println("main over")
}

func funcA() {
	fmt.Println("This is funcA")
}

func funcB() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("recover success!")
		}
	}()
	fmt.Println("This is funcB")
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			panic("B this is a panic")
		}
	}
}

func funcC() {
	defer func() {
		// if msg:= recover():msg!=nil{
		// 	fmt.Println("func c recover success!")
		// }
		fmt.Println("执行延迟函数")
		msg := recover()
		fmt.Println("获取recover的返回值", msg)
	}()
	fmt.Println("this is func C")
	panic("func C has a panic")
}
