package main

import "fmt"

func main() {
	/*
	return词义:"返回"
		用法：
			A：函数中使用return 将结果返回给调用处。
			B：同时结束了函数的执行。

	注意点：
		1.如果一个函数定义了返回值，那么必须使用return将结果返回。return后的数值，必须和声明一致。
			个数，类型，顺序
		2.如果一个函数定义了返回值，函数中如果有分支，循环等，那么保证无论执行哪个分支，都有return可以被执行。
		3.如果一个函数没有定义返回值，那么函数中可以使用return，专门用于结束函数。
		4.可以使用_，舍弃多余的返回值。
	 */

	 res1:= fun1(-30)
	 fmt.Println(res1)

	 fmt.Println(fun2())

	 fun3()

	 fun5()

	 _,res2:=fun4()
	 fmt.Println(res2)
}

func fun1(age int) int{
	if age < 0 {
		return 0 //同时表示结束函数
	}else if age == 0{
		return 0
	}else{
		return age
	}
}

func fun2() int{
	n := 10
	if n > 0{
		fmt.Println("是正数。。")
		return n//结束了函数
	}
	fmt.Println("n的数值是：",n)
	return 0
}


func fun3(){
	m := 1
	if m > 0{
		return
	}
	fmt.Println(m)
	fmt.Println("fun3...")
	//return //?结束了函数的执行
	fmt.Println("over...")

}

func fun4()(int,string){
	return 0,"abc"
}

func fun5(){
	for i:=1;i<=10;i++{
		if i == 5{
			//continue //结束了某一次循环，循环下次继续
			//break //结束了for循环，彻底的结束
			return //结束了函数
		}
		fmt.Println(i)
	}
	fmt.Println("fun5...over...")
}


