package main

import (
	"time"
	"math/rand"
	"fmt"
)

func main()  {

	//1. 设置随机数种子
	var t = time.Now()
	var timesnap = t.Unix() // 精确到秒
	rand.Seed(timesnap)
	//2. 获取随机数
	var random1 = rand.Intn(10)
	var random2 = rand.Intn(10)

	fmt.Println(random1, random2)

	//如果是在for 循环内，需要精确到毫秒级
	for  i := 0;i < 5 ;i++  {
		var t = time.Now()
		var timesnap = t.UnixNano() //精确到毫秒级
		rand.Seed(timesnap)
		var random = rand.Intn(22)
		fmt.Println(random)
	}



}