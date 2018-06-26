package main

import "fmt"

func main()  {
	/*
	if语句的嵌套：
	if 条件1{
		//A段
	}else{
		if条件2{
			//B段
		}else{
			//C段
		}
	}
	简写成：
	if 条件1{
		A段
	
	}else if 条件2{
		B段
	}else if 条件3{
		C段
	}else{
		D段。。
	}
	 */
	sex := "猪"
	 if sex == "男"{
	 	fmt.Println(sex,"，去男厕所。。")
	}else{
		if sex =="女"{
			fmt.Println(sex,"，去女厕所。。")
		}else{
			fmt.Println(sex,"我也不知道了。。")
		}
	 }
	 fmt.Println("----------------------------")
	if sex == "男"{
		fmt.Println(sex,"，去男厕所。。")
	}else if sex =="女"{
		fmt.Println(sex,"，去女厕所。。")
	}else{
		fmt.Println(sex,"我也不知道了。。")
	}

}
