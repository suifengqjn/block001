package main

import "fmt"

func main() {
	/*
	深浅拷贝：
		深拷贝：拷贝数据的副本。对原始数据没有影响
			值类型的数据，默认都是深拷贝
				int，float，string，bool，array，struct
		浅拷贝：拷贝的是数据的地址。
			引用类型的数据，默认都是浅拷贝
				slice，map，function
	 */
	map1 := make(map[int]string)
	map1[1] = "aaa"
	map1[2] = "bbb"
	fmt.Println(map1)
	fmt.Printf("%p\n", map1)

	//1.深拷贝
	a := 10
	b := a //深拷
	fmt.Println(a, b)
	b = 20
	fmt.Println(a, b)

	//2.数组的拷贝
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := arr1
	fmt.Println(arr1, arr2)
	arr2[0] = 100
	fmt.Println(arr1, arr2)
	//数组的浅拷贝
	arr3 := &arr1
	fmt.Println(arr1, arr3)
	(*arr3)[0] = 100 //理论的写法
	arr3[0] = 200 //go优化写法：语法糖
	fmt.Println(arr1, arr3)


	//3.引用类型：默认就是浅拷贝

}
