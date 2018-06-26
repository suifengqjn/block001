package main

import "fmt"

func main() {
	/*
	递归函数：一个函数自己调用自己，就叫做递归函数。
		递归函数要有出口，逐渐向出口靠近

	 */

	res1 := getSum(5) // 15
	fmt.Println(res1)

	//练习：递归算法求5的阶乘
	//练习：兔子生兔子：fibonacci数列：第1,2两项数值都为1,从第3项开始，是前两项之和。使用递归算法，求第12项的数值。


	fmt.Println(getfactorial(5))
	fmt.Println(getFibonacci(12))
}

func getFibonacci(n int)int{
	if n ==1 || n ==2{
		return 1
	}
	return getFibonacci(n-1) + getFibonacci(n-2)
}


func getfactorial(n int)int{
	if n== 1{
		return 1
	}
	return getfactorial(n-1) *n
}

func getSum(n int) int {
	// 5   4   3   2   1
	fmt.Println("************")
	if n == 1 {
		return 1
	}
	return getSum(n-1) + n
	// getSum(4) + 5 = 10+5 = 15
	//  getSum(3) + 4 = 6+4=10

	//   getSum(2)+3 = 3+3 = 6

	//getSum(1)+2 = 1+2 = 3

	// 1
}
