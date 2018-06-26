package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//var slice = []interface{}{1}
	//slice = append(slice, 2)
	//slice = append(slice, "sd")
	//fmt.Println(slice)
	//fmt.Printf("%T\n", slice)
	//
	//funcBack()
	//
	//
	//defer fmt.Println("1")
	//defer fmt.Println("2")
	//fmt.Println(1,f1())
	//fmt.Println(2,f2())
	//fmt.Println(3,f3())
	//
	//var a = 10
	//fmt.Println(&a)

	var x *string = nil
	var arr = []int{1}
	fmt.Println(x, arr, arr[0])

	pointer()

	func(unused string) {
		fmt.Println("Unused arg. No compile error")
	}("what?")


	data := "1234中国"
	fmt.Println(utf8.RuneCountInString(data))
	fmt.Println(len(data))
}

//回调函数

func funcBack() {

	fmt.Printf("%T", opEr)
	fmt.Printf("%T\n", sum)

	var re = sum(2, 3, func(a int, b int) int {
		return a + b
	})

	fmt.Println("==",re)
}

func sum(a, b int, fu func(int, int) int) int {

	var s = fu(a, b)
	return s

}

func opEr(m, n int, fu func(int, int) int) int { //m=100,n=8,fun
	fmt.Println(m, n, fu)
	res := fu(m, n) //res=40
	return res
}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
		fmt.Println(r)
	}(r)
	return 1
}

func pointer()  {

	//创建变量，创建指针，操作指针更改变量，指针的指针。。
	var a = 10
	var p = &a
	fmt.Println(a, p)

	*p = 20
	fmt.Println(a)

	var p2 = &p
	**p2 = 100
	fmt.Println(a)




}