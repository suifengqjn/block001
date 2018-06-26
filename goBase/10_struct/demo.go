package main

import "fmt"

var students []Student
type Student struct{
	no int
	name string
	count int
}

func initInfo(){
	fmt.Println("请输入参加候选人的人数：")
	sum := 0
	fmt.Scanln(&sum)
	students = make([] Student,0,sum)
	for i:=1;i<=sum;i++{
		fmt.Println("请输入第",i,"位候选人姓名：")
		name := ""
		fmt.Scanln(&name)

		s := Student{i,name,0}

		students = append(students,s)
	}
}

func printInfo(){
	//排序
	for i:=1;i<len(students);i++{
		for j:=0;j<len(students)-i;j++{
			if students[j].count < students[j+1].count{
				students[j],students[j+1] = students[j+1],students[j]
			}
		}
	}
	for i:=0;i<len(students);i++{
		s := students[i]
		fmt.Printf("[%d] %s同学，票数是：%d\n",s.no,s.name,s.count)
	}
}

func vote(){
	for{
		fmt.Println("请输入候选人编号，0表示退出")
		n := -1
		fmt.Scanln(&n)
		if n == 0 {
			fmt.Println("结束投票环节")
			break
		}else if n > 0 && n <= len(students){
			students[n-1].count = students[n-1].count+1
		}else{
			fmt.Println("此票无效，请重新输入：")
		}
	}
}

func main() {
	//初始化信息
	initInfo()
	//显示信息
	printInfo()
	//投票
	vote()
	//显示信息
	printInfo()
}
