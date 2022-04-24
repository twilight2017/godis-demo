package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randTest()
	randAnswer()
}

//生成随机数
func randTest() {
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(50))
	fmt.Println(rand.Float64())
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randnum := r1.Intn(10)
	fmt.Println(randnum)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Float64())
	num := rand.Intn(7) + 5
	fmt.Println(num)
}
