package main

import (
	"fmt"
)

func main() {
	sum, avg, count := GetScore(90, 82.5, 73, 64.8)
	fmt.Printf("学员共有%d门成绩，总成绩为：%.2f, 平均成绩为：%.2f", count, sum, avg)
	fmt.Println()
	scores := []float64{92, 72.7, 89.6, 67.9, 98}
	sum, avg, count = GetScore(scores...)
	fmt.Printf("学员共有%d门成绩，总成绩为： %.2f, 平均成绩为：%.2f", count, sum, avg)
}
func GetScore(scores ...float64) (sum, avg float64, count int) {
	for _, value := range scores {
		sum += value
		count++
	}
	avg = sum / float64(count)
	return
}
