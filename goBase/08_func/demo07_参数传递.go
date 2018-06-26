package main

import "fmt"

func main() {
	/*
	参数的传递：
		1.值传递：传递的是数据副本，对于原始数据没有影响的
			值类型的数据，默认都是值传递
			基本类型，array

		2.引用传递：传递的是数据的地址，导致多个变量指向同一块内存，
			引用类型的数据：slice，map
	 */
	a := 1
	fmt.Println("函数调用前a：",a)//1
	fun1(a)
	fmt.Println("函数调用后a：",a)//1

	arr1:=[4]int{1,2,3,4}
	fmt.Println("函数调用前arr：",arr1)//[1,2,3,4]
	fun2(arr1)
	fmt.Println("函数调用后arr：",arr1)//[1,2,3,4]

	fmt.Println("-----------------------------")
	s1:=[]int{1,2,3,4,5}
	fmt.Println("函数调用前s1：",s1)
	fun3(s1)
	fmt.Println("函数调用后s1：",s1)

}
func fun3(s2 []int){
	//fmt.Println(s2)
	s2[0] = 100
	fmt.Println("函数中：s2：",s2)
}

func fun1(a int){
	fmt.Println(a)//1
	a = 100
	fmt.Println("函数中a：",a)//100
}

func fun2(arr [4]int){
	fmt.Println(arr)//[1,2,3,4]
	arr[0] = 100
	fmt.Println("函数中arr：",arr)//[100,2,3,4]
}
/*
练习：设计一个函数，用于接受一个map作为参数。在函数中尝试修改map的值，观察原始数据，数据是否发生变量。
 */


