package main

import "fmt"

type cat struct {
	name string
	age  int
}

func changeName1(c cat) { // c = c1,值传递，传递的是数据的副本
	c.name = "花花"
}
func changeName2(c *cat) { // c = &c1,引用传递，传递的地址
	c.name = "默默"
}


func main() {
	/*
	作为参数：
		结构体对象作为参数？
		结构体对象的指针作为参数？
	作为返回值：
		结构体作为返回值
		结构体指针作为返回值
	 */
	c1 := cat{"tom", 1}
	fmt.Println(c1) //tom,1
	changeName1(c1)
	fmt.Println(c1) //tom，1

	/*
	语法糖：简化的写法：
		数组指针：
			*[4]int

		p1 --->[4]int
		*p1
			(*p1)[0]-->p1[0]

		结构体指针：
			*person
		p2--->struct
		*p2
			*p2.name--->p2.name
	 */
	changeName2(&c1)
	fmt.Println(c1)

	c2 := getCat1()//露娜
	c3 := getCat1()//露娜
	fmt.Println(c2,c3)
	c2.name = "jerry"
	fmt.Println(c2,c3)

	c4 := getCat2()//加菲猫
	c5 := getCat2()//加菲猫
	fmt.Println(c4,c5)
	c4.name="花猫"
	fmt.Println(c4,c5)
}
func getCat1() cat {
	c := cat{"露娜", 1}
	return c
}
func getCat2() *cat {
	//c :=new(cat)
	c := cat{"加菲猫", 2}
	return &c
}
