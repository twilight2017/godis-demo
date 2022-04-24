package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
	sex  string
}

type Student struct {
	Person
	schoolName string
}

func main() {
	//1.实例化并初始化Person
	p1 := Person{"Steven", 35, "男"}
	fmt.Println(p1)
	fmt.Println("---------")
	//2.实例化并初始化Student
	s1 := Student{p1, "北科大"}
	printInfo(s1)
	//
	s2 := Student{Person{"David", 89, "男"}, "北京外国语大学"}
	printInfo(s2)
	//
	s3 := Student{Person: Person{
		name: "Penny",
		age:  98,
		sex:  "男"},
		schoolName: "北大",
	}
	printInfo(s3)
	//
	s4 := Student{}
	s4.name = "Sunny"
	s4.age = 12
	s4.sex = "男"
	s4.schoolName = "北航"
	printInfo(s4)
}

func printInfo(s Student) {
	fmt.Println(s)
	fmt.Printf("%+v \n", s)
	fmt.Printf("姓名：%s, 年龄:%d, 性别:%s, 学校：%s \n", s.name, s.age, s.sex, s.schoolName)
}
