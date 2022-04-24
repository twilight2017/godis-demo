package main

import "fmt"

/*接口用于定义对象的行为，只定义对象应该做什么
实现这种行为的方式（实现细节）由对象来决定
如果某个对象实现了该接口的所有方法，则此对象就实现了该接口*/
type Phone interface {
	call()
}

type AndroidPhone struct {
}

type IPhone struct {
}

func (a AndroidPhone) call() {
	fmt.Printf("I'm an android phone")
}

func (i IPhone) call() {
	fmt.Printf("I'm an iphone")
}

func main() {
	//定义接口类型的变量
	var phone Phone
	phone = new(AndroidPhone)
	fmt.Printf("%T, %v, %p \n", phone, phone, phone)
	phone.call()
	phone = AndroidPhone{}
}
