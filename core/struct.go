package main

import "fmt"

type Teacher struct {
	name string
	age  int8
	sex  byte
}

func main() {
	var t1 Teacher
	fmt.Println(t1)
	fmt.Printf("t1:%T, %v, %q \n", t1, t1, t1)
	t1.name = "Steven"
	t1.age = 35
	t1.sex = 1
	fmt.Println(t1)
	fmt.Println("----------")
	t2 := Teacher()
	t2.name = "DAvid"
	t2.age = 30
	t2.sex = 1
}
