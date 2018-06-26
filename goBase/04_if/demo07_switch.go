package main

import "fmt"

func main(){
	/*
	switch语句，词意"开关"
	switch 变量{
	case 数值1:分支语句1
	case 数值2：分支语句2.。
	。。。
	default：
		最后一个分支语句
	}

	注意点：
	1.switch后变量类型必须和case后的数值类型匹配
	2.case后的数值必须唯一，否则duplicate case 1 in switch
	3.case是无序
	4.default是可选的操作。
	 */
	 a := 1
	switch a {
	default:
		fmt.Println("数据有误。。")
	case 3:
		fmt.Println("这是第三个季度。。")
	case 4:
		fmt.Println("这是第四个季度。。")
	case 1:
		fmt.Println("这是第一季度")
	case 2:
		fmt.Println("这是第二个季度。。")


	}
	fmt.Println("over...")
}
/*
练习1：使用switch语句，实现熊大，熊二，光头强。。。。
练习2：使用switch语句，实现计算器：
	两个数，操作：+，-,*,/
 */
