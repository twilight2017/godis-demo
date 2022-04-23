package main

import "fmt"

func main() {
	a := [4]float64{65.3, 78.9, 89.7, 90.3}
	b := [...]int{2, 6, 9}
	fmt.Printf("数组a的长度为 %d, 数组b的长度为 %d \n", len(a), len(b))
}
