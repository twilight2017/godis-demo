package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.Now()
	testTime()
	time2 := time.Now()
	fmt.Println(time2.Sub(time1).Seconds())
}

func testTime() {
	t := time.Now()
	fmt.Println("1. ", t)
	fmt.Println("2. ", t.Local())
	fmt.Println("3. ", t.UTC())
	t = time.Date(2018, time.January, 1, 1, 1, 1, 0, time.Local)
	fmt.Printf("4. 本地时间%s, 国际统一时间： %s \n", t, t, t.UTC())
	t, _ = time.Parse("2006-01-02 15:09:89", "2018-09-16 05:47:13")
	fmt.Println("5. ", t)
}
