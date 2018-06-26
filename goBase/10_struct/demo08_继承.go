package main

import "fmt"

//1.定义父类
type Person struct {
	name string
	age int
}

//2定义子类
type Student struct {
	Person //模拟继承结构
	school string //子类的新增属性
}
/*
s1.Person.name ---> s1.name
student类中将Person作为一个匿名字段了，那么个该Person中的的字段，对于student来将，就叫做提升字段。
student对象直接访问perosn中的字段。
 */



func main() {
	/*
	继承性：一个类和另一个类的关系。也是面向对象 的三大特征之一
		一个子类继承一个父类。

		意义：A：避免重复的代码。B：扩展类的功能

		A：子类可以直接访问父类的属性和方法
		B：子类可以新增自己的属性和方法
		C：子类可以重写父类已有的方法

	语法：go的实现
	结构体的嵌套
		模拟聚合：has - a
		type A struct{}
		type B struct{
			a A //聚合关系
		}

		模拟继承：is - a
		type C struct{}
		type D struct{
			C //继承
		}



	cat:name,age,eat(),sleep()，爬树()
	dog:name,age,eat(),sleep()，看家()

	dog 继承 cat

	cat 继承 dog


	子类是一个特殊的父类

	学生 是 人

	小汽车 是 车

	工人 是 人

	猫 是 动物
	狗 是 动物


	 */
	//1.创建父类对象
	 p1 := Person{name:"张三",age:30}
	 fmt.Println(p1.name, p1.age)
	 //2.创建子类对象
	 var s1 Student //
	 s1.name = "里斯"//子类对象直接访问父类的属性,提升字段
	 s1.age = 18//子类访问父类的属性，提升字段
	 s1.school = "北京大学"//子类访问自己新增的属性

	 fmt.Println(s1.name,s1.age,s1.school)

	 s2 :=Student{Person{"王五",17},"清华大学"}
	 fmt.Println(s2.name,s2.age,s2.school)

	 s3:=Student{Person:Person{name:"rose",age:19},school:"新东方"}
	 fmt.Println(s3.Person)
	 fmt.Println(s3.Person.name,s3.Person.age)
}
/*
父类：动物name，age
猫：新增属性color
狗：新增属性kind
 */
