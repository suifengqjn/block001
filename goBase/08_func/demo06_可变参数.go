package main

import "fmt"

func main() {
	/*
	可变参数：一个函数的参数，类型确定，但是个数不定(0-多个)，可以使用可变参数
	语法：
		参数名 ... 参数的类型

	函数中，可变参数相当于切片
	调用该函数：可以传入0-多个该中类型的数值。或者是直接传入该中类型的切片...


	注意点：
		A：如果一个函数的参数是可变参数，同时还有其他的参数。可变参数要放在形参列表的最后
		B：一个函数的参数列表中，最多只能有一个可变参数。
	 */
	 //需求：n个整数的和
	 getAdd(1,2,3)

	 getAdd(1,2,3,4,5,6,7,8,9,10)

	 getAdd()

	 s1:=[]int{5,6,7,8}
	 getAdd(s1...)

	 //append()
	 //fmt.Println()
}
func getAdd(nums... int){
	fmt.Printf("%T\n",nums) //	[] int
	sum := 0
	for i:=0;i<len(nums);i++{
		sum += nums[i]
	}
	fmt.Println(sum)

}


//fun1("a","b",1,2,3,4,5,6)
func fun1(s1,s2 string,nums...int){

}
