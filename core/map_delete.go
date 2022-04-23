package main

import "fmt"

func main() {
	//1.声明并初始化一个map
	map1 := map[string]string{
		"element":    "div",
		"width":      "100px",
		"height":     "100px",
		"border":     "solid",
		"background": "none",
	}
	//根据key删除map中某个元素
	fmt.Println("删除前：", map1)
	//不真正用到value，所以传给_
	if _, ok := map1["background"]; ok {
		delete(map1, "background")
	}
	fmt.Println("删除后：", map1)
	//3.清空map
	map1 = make(map[string]string)
	fmt.Println("清空后", map1)
}
