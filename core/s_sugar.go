package main

import "fmt"

type Emp struct {
	name string
	age  int8
	sex  byte
}

func main() {
	emp1 := new(Emp)
	fmt.Printf("emp1: %T, %v, %p \n", emp1, emp1, emp1)
	//语法糖
	emp1.name = "Das"
	emp1.age = 45
	emp1.sex = 1
}
