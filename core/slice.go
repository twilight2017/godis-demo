/*切片的容量是从创建切片的索引开始的底层数组中元素的数量*/
package main

import "fmt"

func main() {
	sliceCap()
}

func sliceCap() {
	arr0 := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	fmt.Println("cap(arr0)=", cap(arr0), arr0)
	//截取数组形成切片
	s01 := arr0[2:8]
	fmt.Printf("%T \n", s01)
	fmt.Println("cap(s01)", cap(s01), s01)
	s02 := arr0[4:7]
	fmt.Println("cap(s02)", cap(s02), s02)
	//截取切片，形成切片
	s03 := s01[3:9]
	fmt.Println("截取s01[3:9]后形成s03", s03)
	s04 := s02[4:7]
	fmt.Println("截取s02[4:7]后形成s04", s04)
	//切片是引用类型
	s04[0] = "x"
	fmt.Print(arr0, s01, s02, s03, s04)
}
