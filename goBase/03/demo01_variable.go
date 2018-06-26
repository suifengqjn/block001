package main

import "fmt"

func main(){
	/*
	变量：
	1.概念：本质就是一小块内存。用于存储某个数值。该数值在程序的运行过程中可以被改变。
	2.使用：
		step1：变量的定义，也叫声明(开辟内存)
			方式一：var 变量名 数据类型
					变量名 = 赋值
			方式二：缩写成一行
				var 变量名 数据类型 = 赋值
			方式三：类型推断，省略数据类型
				var 变量名 = 赋值
			方式四：省略var，简短声明
				变量名:= 赋值
		step2：变量的访问，进行赋值或取值
				直接根据变量名即可
					变量名 = 赋值
	3.注意点：
		A：go语言是静态语言，要求变量的类型和赋值的类型必须一致。
		B：变量名不能冲突。(同一个作用于域内不能冲突)
		C：简短定义方式，左边的变量名至少有一个是新的
		D：简短定义方式，不能定义全局变量。
		E：变量的零值。也叫默认值。
	 */
	 //1.定义变量
	 var a int
	 a = 10
	 fmt.Println("a:",a) //a,10

	 //2.定义，赋值写一行
	 var b int = 20
	 fmt.Println(b)

	 //3.类型推断：变量赋值后，可以根据数值，推断该变量的类型
	 var c = 30
	 fmt.Println(c)

	 fmt.Printf("%T,%d，%b，%o \n",c,c,c,c) //int'

	 //4.省略var，简短声明
	 d := 40
	 fmt.Printf("%T,%d\n",d,d) //int ,40

	 //更改变量的数值
	 a = 100
	 fmt.Println("a:",a)

	 //5.定义多个变量
	 var m, n ,k int
	 m = 1
	 n = 2
	 k = 3
	 fmt.Printf("m:%d,n:%d,k:%d\n",m,n,k)

	 var p,q int = 10,20
	 fmt.Printf("p:%d,q:%d\n",p,q)

	 var r, t ,s = 100, 3.14,"memeda"
	 fmt.Printf("r:%d,%T     t:%f,%T     s:%s,%T\n",r,r,t,t,s,s)

	 //定义一组变量
	 var(
	 	y = 10
	 	z int
	 )
	 fmt.Println(y,z)

	 name,age := "王二狗",30
	 fmt.Println(name,age)
	 name,sex:="rose","女" //no new variables on left side of :=
	 fmt.Println(name,age,sex)

	 //name= 1000//cannot use 1000 (type int) as type string in assignment
}
