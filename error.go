package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	//1
	res := math.Sqrt(-100)
	fmt.Println(res)
	res, err := Sqrt(-100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	//2
	res, err = Divide(100, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	//3 打开不存在的文件
	f, err := os.Open("/abc.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T", err)
	} else {
		fmt.Println(f.Name())
	}
}

//Sqrt
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//Go语言errors包对外提供了可供用户自定义的方法，errors包下的New()函数返回error对象
		//New()的函数参数为字符串，返回值为error类型
		return 0, errors.New("不能给负数计算平方根")
	} else {
		return math.Sqrt(f), nil
	}
}

//Divide
func Divide(dividee float64, divider float64) (float64, error) {
	if divider == 0 {
		return 0, errors.New("除数不能为0")
	} else {
		return dividee / divider, nil
	}
}
