package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
	fmt包：
		打印输出类：
		fmt.Println(...interface{})//打印后换行
		fmt.Print()//打印
		fmt.Printf(string, ...interface{})//格式化打印

		读取键盘：
		Scanln()
		Scanf()

		拼接字符串：
		Sprint()
		Sprintf()
		Sprintln()

		获取error对象
		Errorf()


	 */
	 fmt.Println("aa","bb","cc")
	 fmt.Printf("name是：%s,年龄是：%d\n","王二狗",30)

	 //fmt.Scanln()

	 s1 := fmt.Sprint("helloworld",100,3.14,"memeda") //获取字符串
	 fmt.Println(s1)

	 s5 :=fmt.Sprintf("name：%s，年龄：%d","王二狗",30)
	 fmt.Println(s5)



	 s2:="memeda"
	 i1,err:=strconv.ParseInt(s2,10,64)
	 if err !=nil{
	 	fmt.Println(err.Error())
	 }
	 fmt.Println(i1)

	 var i2 int64 = 100

	 s4 := strconv.FormatInt(i2,10)
	 s3 := "hello" +s4
	 fmt.Println(s3)

}
