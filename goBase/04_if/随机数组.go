package main

import "fmt"
import (
	"math/rand"
	"time"
)

func main() {
	var slice = make([]int, 0, 23)
	for i := 1; i <= 23; i++ {
		slice = append(slice, i)
	}

	fmt.Println(slice)

	var result = make([]int, 0, 10)
	for i := 0; i < 10; i++ {

		var ranIndex = getRandIndex(slice)
		result = append(result, slice[ranIndex])
		removeSlice(slice, ranIndex)
	}

	fmt.Println(slice)
	fmt.Println(result)

	//randomCount()

}

func getRandIndex(s []int) int {

	var count = len(s)
	var time = time.Now()
	var timesnap1 = time.UnixNano()
	rand.Seed(timesnap1)
	var rad = rand.Intn(count - 1)
	fmt.Println(rad)
	return rad

}

func randomCount() {
	var time = time.Now()
	fmt.Println(time)

	//时间戳 秒
	var timesnap1 = time.Unix()
	//时间戳 毫秒
	var timesnap2 = time.UnixNano()
	println(timesnap1, timesnap2)

	//随机数
	//1. 设置随机数种子，
	rand.Seed(timesnap1)
	var rad = rand.Intn(10)
	fmt.Println(rad)

}

func removeSlice(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}
