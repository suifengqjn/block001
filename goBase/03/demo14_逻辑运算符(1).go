package main

import "fmt"

func main()  {
	/*
	逻辑运算符：操作数必须是bool，运算结果也是bool
	逻辑与：&&
			运算规则：所有的操作数都为真，结果才为真，有一个为假，就为假。
				"一假则假，全真才真"
	逻辑或：||
			运算规则：所有的操作数都为假，结果就为假，有一个为真，就为真
				"一真则真，全假为假"
	逻辑非：!
			运算规则：!T-->F,!F-->T
	 */

	 f1 := true
	 f2 := false
	 f3 := true

	 res1 := f1 && f2 && f3
	 fmt.Println(res1)

	 res2 := f1 || f2 || f3
	 fmt.Println(res2)

	 fmt.Println(f1,!f1)
	 fmt.Println(f2,!f2)

	 a := 3
	 b := 2
	 c := 5
	 res3 := a > b && c % a == b && a < (c / b)

	 res4 := b * 2 < c || a / b != 0 || c / a > b
	 res5 := !(c / a == b)
	 fmt.Println(res3,res4,res5)

}
