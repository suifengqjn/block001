package main

import (
	"math"
	"fmt"
)

//1.定义一个接口
type Shape interface {
	peri() float64
	area() float64
}

//2.定义实现类:三角形
type Triangle struct {
	//a float64
	//b float64
	//c float64
	a, b, c float64
}

func (t Triangle)peri() float64  {
	return t.a+t.b+t.c
}
func (t Triangle) area() float64{
	p := t.peri() / 2 //半周长
	s := math.Sqrt(p*(p-t.a)*(p-t.b)*(p-t.c))
	return s
}

type Circle struct {
	radius float64//半径
}

func (c Circle)peri()float64  {
	return c.radius*2*math.Pi
}
func (c Circle)area()float64  {
	return math.Pow(c.radius,2)*math.Pi
}


//测试函数
func testShape(s Shape){
	fmt.Println("周长：",s.peri(),"面积：",s.area())
}

//1.转型
func getType(s Shape){
	/*
	方法一：
		instance，ok:=接口对象.(实际类型)
			如果该接口对象是对应的实际类型，那么instance就是转型之后对象，ok的值为true
	 */

	 if ins,ok:=s.(Triangle);ok{
	 	fmt.Println("是三角形,三边是：",ins.a,ins.b,ins.c)
	 }else if ins,ok:=s.(Circle);ok{
	 	fmt.Println("是圆形，半径是：",ins.radius)
	 }else{
	 	fmt.Println("我也不知道了。。")
	 }
}

//2.转型
func getType2(s Shape){
	/*
	方法二：接口对象.(type),配合switch和case语句使用
	 */
	switch ins:=s.(type) {
	case Triangle:
		fmt.Println("三角形。。。",ins.a,ins.b,ins.c)
	case Circle:
		fmt.Println("圆形。。。。",ins.radius)
	//case int:
	//	fmt.Println("整型数据。。")
	}
}




func main() {
	/*
	多态：一个事物的多种形态
		go语言：通过接口模拟多态性
		一个实现类的对象：
			看作是一个实现类类型：能够访问实现类中的方法和属性
			还可以看作是对应的接口类型：只能够访问接口中定义的方法

	接口的用法：
	用法一：一个函数如果接收接口类型作为参数，那么实际上可以传入该接口的任意实现类对象作为参数。
	用法二：定义一个类型为接口类型，那么实际上可以赋值任意实现类的对象。
			如果定义了一个接口类型的容器，实际上该容器中可以存储任意的实现类对象。
	 */

	 var t1 Triangle
	 t1 =Triangle{3,4,5}
	 fmt.Println(t1.peri())
	 fmt.Println(t1.area())
	 fmt.Println(t1.a,t1.b,t1.c)

	 var s1 Shape
	 s1 = t1
	 fmt.Println(s1.peri())
	 fmt.Println(s1.area())

	 var c1 Circle
	 c1 = Circle{4}
	 fmt.Println(c1.peri())
	 fmt.Println(c1.area())
	 fmt.Println(c1.radius)

	 var s2 Shape = Circle{5}
	 //s2 = c1
	 fmt.Println(s2.area(),s2.area())


	testShape(t1)

	//定义一个接口类型的数组
	//var arr[4]Shape
	//arr[0] = t1
	arr := [4]Shape{t1,s1,c1,s2}
	fmt.Println(arr)


	fmt.Println("----------------------------")
	//接口类型的对象-->对应实现类类型
	getType(c1)

	getType2(t1)
}
/*
定义一个接口：Animal
eat(),sleep()
定义两个实现类：
Cat，属性name
Dog，属性color


创建一个数组，存储5个Animal，设计一个函数，接受这个数组作为参数。
打印输出数组中每个动物的eat(),sleep(),如果是猫显示name，如果是狗现实color
 */
