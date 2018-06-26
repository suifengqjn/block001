package main

import "fmt"

func main() {
	/*
	函数指针：一个指针，指向了一个函数的指针。
		因为go语言中，function，默认看作一个指针，没有*。
			slice，map，function是引用类型的数据，存储的就是数据的地址。看作一个默认的指针，没有*
	指针函数：一个函数，该函数的返回值是一个指针。

	 */
	var a func()
	a = fun1 //将fun1赋值给a
	a()

	arr1 := fun2()
	arr2 := fun2()
	fmt.Printf("arr1的类型：%T,地址：%p,数值：%v\n",arr1,&arr1,arr1)
	fmt.Printf("arr2的类型：%T,地址：%p,数值：%v\n",arr2,&arr2,arr2)
	p3 := fun3()
	p4 := fun3()
	fmt.Printf("p3：%T,地址：%p,数值：%v\n",p3,p3,p3)
	fmt.Printf("p4：%T,地址：%p,数值：%v\n",p4,p4,p4)

}

func fun2() [4]int { //普通函数
	arr := [4]int{1, 2, 3, 4}
	return arr
}

func fun3() *[4]int { //指针函数
	arr := [4]int{1, 2, 3, 4}
	return &arr
}

func fun1() {
	fmt.Println("fun1.。。")
}
