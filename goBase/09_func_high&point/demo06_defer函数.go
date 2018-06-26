package main

import (
	"fmt"
)

func main() {
	/*
	defer函数：一个函数的执行被延迟了。
		用于：defer 对象.close()  临时文件.delete()

	多个函数被defer，先延迟的最后执行，后延迟的先执行。。

	defer传递参数的时候：defer 调用函数时，就已经 传递数据了，只是暂时不执行。

	defer函数：
	1.当外围函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束执行。
	2.当执行外围函数中的return语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正返回。
	3.当外围函数中的代码引发运行恐慌时panic，只有其中所有的延迟函数都执行完毕后，该运行时恐慌才会真正被扩展至调用函数。
	 */
	 //1.defer函数的执行
	 //defer fun1("hello")
	 //fmt.Println("12345")
	 //defer fun1("world")//被延迟
	 //fmt.Println("王二狗")

	 //2.defer的参数传递
	 //a := 2 //2
	 //defer fun2(a) //调用了就传递参数，只不过暂时不执行
	 //a++// 3
	 //fmt.Println("main:a",a)//3


	 //练习题：
	 //s := "helloworld"
	 //for i:=0;i<len(s);i++{
	 //	defer fmt.Printf("%c",s[i])
	 //}

	 //3.defer函数的注意点：
	 res1:=f3()
	 fmt.Println(res1)



}

func f1() (result int) {
	fmt.Println("11111")
	fmt.Println("2222")
	fmt.Println("33333")
	defer func() {
		result++
	}()
	fmt.Println("44444")
	return 0
}
/*
f2()-->
step1:
	r = 0
step2:
	t = 5
step3:
	返回值：r = 5
step4:defer
	t = 10
step5:return r
 */
func f2() (r int) {

	t := 5
	defer func() {
		t = t + 5
		fmt.Println("f2中：t,",t)
	}()
	return t
}
/*


step1:
	r = 0
step2:
	返回值：r =1
step3：defer


step4:return
	1
 */

func f3() (r int) {
	defer func(r int) { //r = 0
		r = r + 5
		fmt.Println("f3中,r:",r)
	}(r)
	return 1
}


func fun1(s string){
	fmt.Println(s)
}
/*
1.利用defer 函数的特点，倒叙打印字符串
 */


 func fun2(num int){
 	fmt.Println("fun2中，num：",num)//2
 }