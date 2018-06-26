package main

import "fmt"

func main() {
	/*
	1.创建一个map用于存储一个人的信息：
	"name"：王二狗
	"age"：30
	sex：男性
	address："北京市xx路xx号"

2.打印遍历该map
3.修改sex为女性
4.如果想map中添加重复的key，会如何？
	 */
	//1.创建map存储一个人的信息
	map1 := make(map[string]string)
	map1["name"] = "王二狗"
	map1["age"] = "30"
	map1["sex"] = "男性"
	map1 ["address"] = "北京市XX路XX号"
	fmt.Println(map1)


	//2.第二个人
	map2 :=make(map[string] string)
	map2["name"]= "李小花"
	map2["age"] = "18"
	map2["sex"] ="女性"
	map2["address"] = "上海市。。。"
	fmt.Println(map2)

	//3，第三个人的信息
	map3 :=map[string]string{"name":"rose","age":"19","sex":"女性","address":"杭州市"}
	fmt.Println(map3)


	//将map存储到slice中
	//s1 := []map[string]string{}
	s1 := make([] map[string]string,0,3)
	s1 = append(s1,map1)
	s1 = append(s1,map2)
	s1 = append(s1,map3)

	//遍历切片
	for i,val :=range s1{
		//val -->map1,map2,map3
		fmt.Println("第",i+1,"个人的信息：")
		fmt.Printf("\t姓名：%s,年龄：%s,性别：%s,地址：%s\n",val["name"],val["age"],val["sex"],val["address"])
	}

}
