package main

import "fmt"

//1.定义一个类，就是定义一个结构体
type Person struct{//field,method,function
	name string
	age int
	sex string
}


func main() {
	//方法一：创建对象
	//2.根据类，实例化对象--->创建结构体类型的对象
	var p1 Person
	fmt.Println(p1)
	//3.对象访问该类属性，进行赋值和取值
	p1.name = "王二狗"
	p1.age = 20
	p1.sex = "男"
	fmt.Println(p1)
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n",p1.name,p1.age,p1.sex)
	//map[key] =value, map[key]-->value

	//方法二：
	p2 := Person{}
	p2.name = "rose"
	p2.age = 19
	p2.sex ="女"
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n",p2.name,p2.age,p2.sex)

	//方法三：
	p3 := Person{name:"李晓华",age:28,sex:"女"}
	fmt.Println(p3)

	p4 := Person{
		name:"隔壁老王",
		age:45,
		sex:"男",
	}
	fmt.Println(p4)
	//方法四：
	p5 := Person{"田中天",13,"男"}//注意赋值顺序
	fmt.Println(p5)

	//方法五：使用内置的函数new(),go语言中专门用于创建某种类型的指针的函数。
	p6 := new(Person)
	fmt.Printf("%T,%T\n",p1,p6) //main.Person,*main.Person
	(*p6).name = "jerry"
	p6.age = 30//语言糖，简化写法  (*p6).name--->p6.name
	p6.sex= "女"
	fmt.Println(p6)
	fmt.Printf("%p\n",p6)
}
/*
1.练习1：定义一个狗
	color，age，sex，type
2.练习2：创建3个狗，分别赋值，并打印狗的属性信息
3.练习3：结构体对象，是值类型？还是引用类型？
 */
