package main

import "fmt"

//1.定义一个父类
type Person struct {
	name string
	age int
}
type Student struct {
	Person
	school string
}
//2.方法
func (p Person) eat(){
	fmt.Println("父类的方法，吃窝窝头。。")
}

func (s Student) study(){
	fmt.Println("学生学习啦。。")
}

func (s Student)eat(){
	fmt.Println("子类重写的方法：吃炸鸡喝啤酒。。。")
}

func main() {
	/*
	继承中的方法：
		1.子类可以直接访问父类的属性恶化方法
		2.子类可以新增自己的属性和方法
		3.子类可以重写父类的方法(orverride，就是将父类已有的方法，重新实现)
	 */
	 p1:=Person{"王二狗",30}
	 fmt.Println(p1.name,p1.age)//父类的对象访问父类的属性
	 s1 := Student{Person{"rose",18},"清华大学"}
	 fmt.Println(s1.Person.name) //可以省略Person
	 fmt.Println(s1.name,s1.age)//子类对象直接访问父类的属性
	 fmt.Println(s1.school) //子类访问自己的新增属性
	 p1.eat()//父类的对象，访问父类的方法
	 s1.eat()//子类访问父类的方法，但是存在重写，那么会访问子类重写的方法。。

	 s1.study() //子类对象访问子类的新增方法
}
