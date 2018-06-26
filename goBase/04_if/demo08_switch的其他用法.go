package main

import "fmt"

func main() {
	/*
		1.标准写法：
			switch 变量{
			case 数值1：
			case 数值2：
			。。。
			}

		2.switch后可以省略变量，相当于作用在true
			switch { //true
			case true：
			case false
			}

		3.case后可以同时跟多个数值
			switch 变量{
			case 数值1,数值2,数值3.。。。：

			case 数值4,数值5,数值6.。。。：

			}

		4.switch后也可以添加一条初始化语句
			switch 语句;变量{

			}
	*/

	monDayCount()
	return

	switch {
	case true:
		fmt.Println("执行此处了。。true。。。")
	case false:
		fmt.Println("执行了。。。false。。。")

	}
	//练习1：使用switch实现成绩：优秀，良好，中等，及格，不及格

	letter := "M"
	switch letter {
	case "A", "E", "U", "I", "O":
		fmt.Println(letter, "是一个元音")
	case "D", "F", "G":
		fmt.Println("是E或F或G。。")
	default:
		fmt.Println("其他。。")
	}
	//练习2：使用switch完成。。
	month := 5
	day := 0
	if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		day = 31
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		day = 30
	} else if month == 2 {
		day = 28
	} else {
		fmt.Println(" 无效的月份。。")
	}
	fmt.Println(month, "，的天数是：", day)

	switch language := "golang"; language { //作用域
	case "golang":
		fmt.Println("Go语言。。")
	case "Java":
		fmt.Println("Java语言。。")
	case "Python":
		fmt.Println("Python语言。。。")
	case "JavaScript":
		fmt.Println("Js。。。")
	}
	//fmt.Println(language) //undefined: language

}

func monDayCount() {

	var mon int
	fmt.Scanln(&mon)

	switch mon {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 days")
	case 2:
		fmt.Println("28 days")
	case 4, 6, 9, 11:
		fmt.Println("30 days")
	}

}
