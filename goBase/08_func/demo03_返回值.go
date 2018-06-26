package main

import (
	"fmt"
)

func main() {
	/*
	函数的返回值：将一个函数的执行结果，返回给调用处。执行结果就叫该函数的返回值。
		return语句

	一个函数的定义有返回值，那么函数中必须使用return语句返回结果。
			返回的结果必须和函数的定义一致：类型，个数，顺序

	return关键字：
		A：将 函数的结果返回给调用处
		B：同时结束了函数的执行

	 */
	//求1-10的和，并放大2倍输出
	res := getSum() // 相当于res = sum
	fmt.Printf("%T,%d\n", res, res)
	fmt.Println(res * 2)

	fmt.Println(getSum2()) // 函数名() = return后的数值

	peri,area :=rectangle(3,4)
	fmt.Println("周长：",peri,"面积：",area)

	peri2,area2 :=rectangle2(5,8)
	fmt.Println(peri2,area2)

	//函数的返回值有多个，如果只要某个，其他的可以使用_进行舍弃
	peri3,_ := rectangle(1,4)
	fmt.Println(peri3)

	_,area3 := rectangle(1,4)
	fmt.Println(area3)

}

//1.函数名：getSum,没有参数，有一个int类型数据返回
func getSum() int {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	//fmt.Println("1-10的和：",sum)
	return sum //将return后的数据，返回给调用处
}

//也可以写成这样：定义函数时，直接写清楚返回值是谁
func getSum2() (sum int) {
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return //省略sum不写，也表示返回sum
}

//2.支持一个函数同时返回多个数据
//求一个矩形的周长和面积
func rectangle(len, wid float64) (float64,float64){
	peri := (len + wid) * 2 //周长
	area := len * wid
	return peri, area
}
//也可以写成
func rectangle2(len, wid float64)(peri float64,area float64){
	peri = (len +wid) * 2
	area = len * wid

	return
}
