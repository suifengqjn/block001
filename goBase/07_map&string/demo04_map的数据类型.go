package main

import "fmt"

func main() {
	/*
	map的数据类型：%T
	一：数据类型
		基本数据类型：
			int，float，string，bool
		复合数据类型：
			array，slice，map，function，struct，interface。。。
			数组：[size]数据类型
			切片：[]数据类型
					[]map[string]string

			映射：map[key数据类型]value数据类型
				map[string]map[string]string
	二：存储特点
		值类型：
			int，float，string，bool，array，struct
		引用类型：
			slice，map
	 */

	 map1:=make(map[int]string)
	 map2:=make(map[string]int)
	 fmt.Printf("%T\n",map1) //map[int]string
	 fmt.Printf("%T\n",map2) //map[string]int

	 map3:=make(map[string]map[string]string)
	 m1:=make(map[string]string)
	 m1["name"]="王二狗"
	 m1["age"]= "30"
	 m1["salary"]="3000"

	 map3["hr"] = m1
	 fmt.Println(map3)

	 m2 :=make(map[string]string)
	 m2["name"]="Rose"
	 m2["age"]="28"
	 m2["salary"]="8000"

	 map3["总经理"] = m2
	 fmt.Println(map3)

	 fmt.Println("---------------------------------------------")
	 //引用类型
	 map4 :=make(map[string] string)
	 map4["王二狗"] = "矮矬穷"
	 map4["rose"]= "白富美"
	 map4["田中天"] = "住在隔壁"
	 fmt.Println(map4)

	 map5 :=map4
	 fmt.Println(map5)

	 map5["王二狗"] = "高富帅"
	fmt.Println(map4,map5)
}
