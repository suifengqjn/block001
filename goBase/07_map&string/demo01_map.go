package main

import "fmt"

func main() {
	/*
	map：映射，是一种专门用于存储键值对的集合。属于引用类型

	存储特点：
		A：存储的是无序的键值对
		B：键不能重复，并且和value值一一对应的。
				map中的key不能重复，如果重复，那么新的value会覆盖原来的，程序不会报错。

	语法结构：
		A：创建map
			var map1 map[key类型]value类型
				nil map，无法直接使用
			var map2  = map[key类型]value类型 {key：value，key：value，key：value。。。。}
			var map3 = make(map[key类型]value类型)
		B：添加/修改
			如果key不存在，就是添加数据，如果存在就是修改
			map[key] = value,将键值对，存入到map中，无序
		C：获取
			value, ok := map[key]
				根据key 获取map中对应的value值。
					如果键值对存在，value就是对应的数据，ok为true
					如果键值对不存在，value就是值类型的默认值，ok为false

			遍历：for range map
					key，value
		D：修改
		E：删除,使用go语言的内置函数delete()，
			delete(map,key),如果key存在，那么可以直接删除，如果key不存在，删除失败，但是程序不报错。
	 */

	//1.创建map
	var map1 map[int]string
	var map2 = map[string]int{"go": 98, "python": 68, "java": 89, "html": 93}
	var map3 = make(map[int]string)
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	//2.nil map,//panic: assignment to entry in nil map
	fmt.Println(map1 == nil) //true
	fmt.Println(map2 == nil) // false
	fmt.Println(map3 == nil) // false

	if map1 == nil {
		map1 = make(map[int]string)
		fmt.Println(map1 == nil)
	}
	//3.存储键值对到map中
	//map1[key] = value
	map1[1] = "hello"
	map1[2] = "world"
	map1[3] = "memeda"
	map1[4] = "三胖斯密达"
	map1[5] = ""

	// 4.获取数据，根据key获取value值
	/*
	根据key获取对应的value，如果key 存在，取对应的数值，如果key不存在，获取的是该值类型的默认值。
	 */
	fmt.Println(map1)
	fmt.Println(map1[1]) //根据key为1,获取对应的value数据
	fmt.Println(map1[30]) //""
	v1,ok := map1[30]
	if ok {
		fmt.Println("对应的数值是：",v1)
	}else{
		fmt.Println("操作的key不存在，获取到了默认值：",v1)
	}

	//遍历
	for k,v:=range map1{
		fmt.Println(k, v)
	}

	//5.修改数据
	fmt.Println(map1)
	map1[3] = "如花"
	fmt.Println(map1)

	//6.内置函数delete
	delete(map1,1)
	fmt.Println(map1)
	delete(map1,11)
	fmt.Println(map1)

	//7.长度,
	fmt.Println(len(map1))
}
