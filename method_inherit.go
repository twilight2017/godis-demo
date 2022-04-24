/*如果匿名字段实现了一个方法，那么包含这个匿名字段的struct也可以调用这个方法*/
package main

import "fmt"

type Human struct {
	name, phone string
	age         int
}

type Stu struct {
	Human  //匿名字段
	shcool string
}

type Employee struct {
	Human   //匿名字段
	company string
}

func main() {
	s1 := Stu{Human{"Daniel", "1246i9789", 13}, "十一中学"}
	e1 := Employee{Human{"Steven", "1263889", 45}, "1000dada"}
	s1.SayHi()
	e1.SayHi()
}

func (h *Human) SayHi() {
	fmt.Printf("大家好！我是%s, 我%d岁， 我的联系方式是： %s\n", h.name, h.phone, h.age)
}
