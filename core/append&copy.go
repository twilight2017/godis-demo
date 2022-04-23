package main

import "fmt"

func main() {
	fmt.Println("1.-------------")
	numbers := make([]int, 0, 20) //底层数组元素个数为20，当前元素个数为0的切片
	printSlices("numbers", numbers)
	numbers = append(numbers, 0)
	printSlices("numbers", numbers)
	/*向切片添加一个元素*/
	numbers = append(numbers, 1)
	printSlices("numbers", numbers)
	/*同时向切片添加多个元素*/
	numbers = append(numbers, 2, 3, 4, 5, 6, 7)
	printSlices("numbers", numbers)
	fmt.Println("2.-------------")
	//追加一个切片
	s1 := []int{100, 200, 300, 400, 500, 600, 700}
	numbers = append(numbers, s1...)
	printSlices("numbers", numbers)
	fmt.Println("3.-------------")
	//删除第一个元素
	numbers = numbers[1:]
	printSlices("numbers", numbers)
	//删除最后一个元素
	numbers = numbers[:len(numbers)-1]
	printSlices("numbers", numbers)
	//删除中间一个元素
	a := int(len(numbers) / 2)
	fmt.Println("中间数：", a)
	numbers = append(numbers[:a], numbers[a+1:]...)
	printSlices("numbers", numbers)
	fmt.Println("4.--------------")
	/*创建切片number1是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), 2*(cap(numbers)))
	//复制numbers的内容到numbers1
	count := copy(numbers1, numbers) //copy函数返回的是复制后的个数
	fmt.Println("复制个数：", count)
	printSlices("numbers1", numbers1)
	numbers[len(numbers)-1] = 99
	numbers1[0] = 100
	printSlices("numbers1", numbers1)
	printSlices("numbers", numbers)
}

//输出切片格式化信息
func printSlices(name string, x []int) {
	fmt.Print(name, "\t")
	fmt.Printf("addr: %p \t len=%d \t cap=%d \t slice = %v \n", x, len(x), cap(x), x)
}
