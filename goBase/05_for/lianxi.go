package main

import "fmt"

func main() {

	/*

			*
		   ***
		  *****
		 *******
		*********

	*/

	// 三角形
	for i := 0; i < 5; i++ { //5 行

		for j := 5; j > 0; j-- {

			if j > i+1 {
				print(" ")
			}

		}

		for k := 0; k < (i*2 + 1); k++ {
			fmt.Print("*")
		}

		fmt.Println()

	}

	fmt.Println("-----------")

	//倒三角形
	for i := 5; i > 0; i-- { //5 行

		for j := 5; j > 0; j-- {

			if j > i {
				print(" ")
			}
		}

		for k := 0; k < (i*2 - 1); k++ {
			fmt.Print("*")
		}

		fmt.Println()

	}
	fmt.Println("-----------")
	//斐波那契额数列求和
	fmt.Println("斐波那契额数列求和", feibonaqie(12))
	fmt.Println("-----------")

	//输入年，输入月，打印日历

	var year int
	var month int
	fmt.Println("请输入年份")
	fmt.Scanln(&year)
	fmt.Println("请输入月份")
	fmt.Scanln(&month)

	//这个月有几天
	var monthCount = monthDays(year, month)

	fmt.Printf("这个月有%d天\n", monthCount)

	if month == 1 || month == 2 {
		year -= 1
		month += 12
	}
	// 这个月的1号是星期几
	var monfirstDay = weekDay(year, month, 1)
	fmt.Printf("这个月的1号是星期%d\n", monfirstDay)

	for i := 1; i <= 7; i++ {
		fmt.Printf("星期%d\t", i)
	}

	//打印前面的几个空格
	fmt.Println()
	for i := 1; i < monfirstDay; i++ {
		fmt.Printf("\t")
	}
	//打印月份
	for i := 1; i <= monthCount; i++ {

		fmt.Printf("%d\t", i)

		if (i+monfirstDay-1)%7 == 0 {
			fmt.Println()
		}
	}

}

// feibonaqie
func feibonaqie(index int) int {

	if index == 0 || index == 1 {
		return 1
	} else if index == 2 {
		return 2
	} else {
		return feibonaqie(index-1) + feibonaqie(index-2)
	}
}

//基姆拉尔森 计算星期几
func weekDay(year int, month int, day int) int {
	var d = (day + 1 + 2*month + 3*(month+1)/5 + year + year/4 - year/100 + year/400) % 7
	if d == 0 {
		d = 1
	}
	return day
}

// 计算当月的天数
func monthDays(year int, month int) int {

	if month == 2 {
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			return 29
		} else {
			return 28
		}
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	} else {
		return 31
	}
}
