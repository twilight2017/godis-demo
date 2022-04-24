package main

import "fmt"

type Address struct {
	province string
	city     string
}
type Person struct {
	name    string
	age     int
	address *Address
}

func main() {
	p := Person{}
	p.name = "Steven"
	p.age = 35
	addr := Address{}
	addr.province = "北京市"
	addr.city = "海淀区"
	p.address = &addr
	fmt.Println(p.address.city)
	p.address.city = "昌平区"
	fmt.Println(p.address.city)
	addr.city = "大兴区"
	fmt.Println(p.address.city)
}
