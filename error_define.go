package main

import (
	"fmt"
	"time"
)

//1.定义结构体，表示自定义错误类型
type MyError struct {
	when time.Time
	what string
}

//Error()
func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.when, e.what)
}

//getArea()
func getArea(width, length float64) (float64, error) {
	errorInfo := ""
	if width < 0 && length < 0 {
		errorInfo = fmt.Sprintf("长度:%v, 宽度:%v, 均为负数", length, width)
	} else if length < 0 {
		errorInfo = fmt.Sprintf("长度为负数")
	} else if width < 0 {
		errorInfo = fmt.Sprintf("宽度为负数")
	}
	if errorInfo != "" {
		return 0, MyError{time.Now(), errorInfo}
	} else {
		return width * length, nil
	}
}

func main() {
	res, err := getArea(-4, -5)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("面积为：", res)
	}
}
