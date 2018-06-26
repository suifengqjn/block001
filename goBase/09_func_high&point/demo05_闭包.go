package main

import "fmt"

func main() {
	/*
		函数式编程：
			支持将函数作为另一个函数的参数，叫回调函数
			支持将函数作为另一个函数的返回值，


		一个函数有内层函数，该内层函数中，还会操作外层函数的局部变量(外层函数中定义，外层函数的参数)，并且该外层函数的返回值就是这个内层函数。
		该内层函数和局部变量，统称叫做闭包结构。closure
	*/

	res1 := increment() //res1=fun
	fmt.Println(res1)   //0x489e00
	fmt.Printf("%T\n", res1)

	res2 := res1()    //调用res1(),那么就是调用fun(),实际是调用内层的匿名函数
	fmt.Println(res2) //1
	res3 := res1()
	fmt.Println(res3) //2
	res4 := res1()
	fmt.Println(res4) //3

	res5 := increment()
	res6 := res5()
	fmt.Println(res6) //1
	fmt.Println(res5())
	fmt.Println(res5())
	fmt.Println(res5())
	fmt.Println(res5())
}

func increment() func() int { //increment：外层函数
	i := 0

	fun := func() int { //内层函数
		i++
		return i
	}
	return fun //将内层函数返回：本质是返回该内层函数的地址
	//return fun()//将内层函数的执行结果，返回。
}

/*
练习题：利用闭包，写一个倒计时：
	10,9,8,7.....
*/
