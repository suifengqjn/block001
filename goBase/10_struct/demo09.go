package main

import (
	"day10_struct/employee"
	"fmt"
)

func main() {
	//创建一个emp对象
	e1:=employee.Emp{Name:"张三",Age:30,Salary:88.9}
	fmt.Println(e1)
}
