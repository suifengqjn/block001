package main

import "fmt"

func main()  {
	/*
	常量：
	1.常量：同变量类似，程序执行过程中，数值不可以改变。
	2.常数：固定的数值
		100,"abc",3.14
	3.定义
		显示类型定义：
			const 常量名 数据类型 = 赋值
		隐式类型定义：
			const 常量名 = 赋值
	4.常量名的习惯用法：
		所有的字母都大写：
			公共的
			如果私有，前加小写字母c
	 */
	 //1.定义常量
	 const PATH string = "http://wwww.baidu.com" //显示定义
	 const PI = 3.14                             //隐士定义
	 fmt.Println(PATH)
	 fmt.Println(PI)

	 //2.尝试修改常量的数值
	 //PATH = "http://www.baidudu.com" //cannot assign to PATH

	 //3.定义一组常量
	 const c1,c2,c3 = 1,3.14,"haha"

	 const(
	 	MALE = 1
	 	FEMALE = 2
	 	UNKNOW = 0
	 )
	 fmt.Println(MALE,FEMALE,UNKNOW)

	 //4. 一组常量中，如果某个常量没有初始值，那么默认和上一行一致
	 const(
	 	a int = 100
	 	b
	 	c string = "斯密达"
	 	d
	 	e
	 )
	 fmt.Printf("a:%T,%d\n",a,a)
	 fmt.Printf("b:%T,%d\n",b,b)
	 fmt.Printf("c:%T,%s\n",c,c)
	 fmt.Printf("d:%T,%s\n",d,d)
	 fmt.Printf("e:%T,%s\n",e,e)

	 //5.枚举类型：使用常量组作为枚举类型。一组相关数值的数据
	 const(
	 	spring = 1
	 	summer = 2
	 	autumn = 3
	 	winter = 4
	 )


}
