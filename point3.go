/*
  有一个指针数组，另有一个元素个数与之相同的数组，将该数组中每个元素的地址赋值给该指针数组。
  也就是说该指针数组与某一个数组完全对应。可以通过*指针变量获取到该地址所对应的数值
*/
package main

import "fmt"

const COUNT = 4

func main() {
	a := [COUNT]string{"as", "cs", "ADS", "poi"}
	i := 0
	//定义指针数组
	var ptr [COUNT]*string
	fmt.Printf("%T, %v \n", ptr, ptr)
	for i = 0; i < COUNT; i++ {
		//将数组中每个元素的地址赋值给指针数组
		ptr[i] = &a[i]
	}
	//获取指针数组中第一个值并打印
	fmt.Println(ptr[0])
	//根据指针数组元素的每个地址获取该地址所指向元素的值
	for i = 0; i < COUNT; i++ {
		fmt.Printf("a[%d] =%s \n", i, *ptr[i])
	}
}
