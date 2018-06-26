package main

import "fmt"

type dog struct {
	color string
	age int
	kind string//品种
}

func main() {
	/*
	struct的数据类类型：
		值类型：默认是深拷贝
		如果传递指针地址，可以实现浅拷贝
	 */
	d1 := dog{"黑色",1,"二哈"}//dog
	fmt.Println(d1)
	d2 := d1//深拷贝  dog
	fmt.Println(d2)

	d2.kind = "泰迪"
	fmt.Println(d1,d2)

	d3 :=dog{"白色",3,"萨摩耶"} //dog
	fmt.Println(d3)

	//实现结构体的浅拷贝
	d4 :=new(dog)//*dog
	d4.color = "黄色"
	d4.age= 2
	d4.kind="中华田园犬"

	d5 :=d4//*dog
	fmt.Println(d4,d5)
	d5.kind = "牧羊犬"
	fmt.Println(d4,d5)


	d6 := &d1//*dog
	d6.kind = "金毛"
	fmt.Println(d1,d6)
	fmt.Println(d1,d2,d3)
	fmt.Println(*d4,*d5,*d6)


}
