package main

import "fmt"

type student struct{
	//name string
	//age int
	string //匿名字段
	int//匿名字段
	//string//因为字段没有名字，默认使用类型作为名字，那么匿名字段的类型不能冲突。
}

func main() {
	/*
	匿名结构体和匿名字段：
		匿名结构体：没有名字的结构体，在创建匿名结构体时，同时创建对象
			变量名:=struct{
				定义字段
			}{
				字段进行赋值
			}

		匿名字段：一个结构体的字段没有名字
			理解为，如果一个字段没有名字，那么默认使用类型作为字段名

	匿名函数：没有名字的函数，随着定义的时候直接进行调用
	 */
	//匿名函数
	func() {
		fmt.Println("helloworld...")
	}()

	//c1 := car{color:"火红"}
	//匿名结构体,同时创建该匿名结构体的对象
	p1 := struct {
		name string
		age  int
	}{
		name: "张三",
		age:  30,
	}
	fmt.Println(p1.name, p1.age)

	//fmt.Println(struct {
	//	name string
	//}{"李晓华"}.name)

	s1 :=student{"张三",18}
	fmt.Println(s1)
	fmt.Println(s1.string,s1.int)
}
/*
练习1：设计几个函数，分别接收结构体，结构体指针作为参数，返回值是结构体，结构体指针。
练习2：设计一个匿名结构体，并创建对象，设计一个结构体中有匿名字段，创建对象使用。
 */
