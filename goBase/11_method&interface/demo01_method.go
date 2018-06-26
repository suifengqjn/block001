package main

import "fmt"

//1.定义一个工人类,猫类
type Worker struct {
	//字段
	name string
	age int
	sex string
}

type Cat struct {
	color string
	age int
}

//2.定义类的行为:方法
/*
语法：
func （接受者） 方法名(参数列表)(返回值列表){}

函数可以直接进行调用
方法必须由接受者调用，谁调用谁就是接受者。
 */
func (w Worker) work(){//w = w1  w = w2
	fmt.Println(w.name,"在工作。。。")
}

func (w *Worker) rest(){ // w = w1的地址，w=w2的地址
	fmt.Println(w.name,"在休息。。")
}

func (w *Worker) printInfo(){
	fmt.Printf("工人姓名：%s,年龄：%d,性别：%s\n",w.name,(*w).age,w.sex)
}


func (c Cat) eat(){
	fmt.Println(c.color,"颜色的猫，在吃东西。。")
}

func (c *Cat) printInfo(){
	fmt.Printf("猫咪的颜色：%s，年龄%d\n",c.color,c.age)
}

func test(){
	fmt.Println("我是一个函数啊。。。")
}

func main() {
	/*
	方法：method，同函数类似，区别需要有接受者。(也就是调用者)

	方法同函数的对比：
		A：意义
			方法：某个类别的行为功能。需要有接收者调用。
			函数：一段独立功能的代码。可以直接调用

		B：语法
			函数：命名不能冲突
			方法：方法名可以一致，只需要接收者不同即可。
	 */

	 //3.创建对象
	 w1 := Worker{name:"王二狗",age:30,sex:"男"}
	 w1.work()//对象调用方法

	 w2 := Worker{"李小花",28,"女"}
	 w2.work()

	 w3 := &w2
	 fmt.Println(w3)
	 w3.work()

	 w3.rest()
	 w1.rest()

	 w1.printInfo()
	 w2.printInfo()
	 w3.printInfo()


	 c1 := Cat{}
	 c1.color="白色"
	 c1.age=1
	 c1.eat()
	 c1.printInfo()//方法

	 test()//函数

}
/*
练习1：创建一个小汽车类：车的品牌，车的颜色，方法：跑()，加速()，停止()
练习2：继承中方法的使用?
	A：子类可以直接访问 父类的属性和方法
	B：子类可以新增自己的属性和方法
	C：子类可以重写父类已有的方法
 */
