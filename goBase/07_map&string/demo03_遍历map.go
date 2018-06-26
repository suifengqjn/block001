package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]string)
	map1[1] = "红孩儿"
	map1[2] = "小钻风"
	map1[3] = "白骨精"
	map1[4] = "白素贞"
	map1[5] = "金叫大王"

	//1.遍历map
	for k,v:=range map1{
		fmt.Println(k,v)
	}
	fmt.Println("------------------------")
	//2.有序遍历map-->按照key的大小顺序
	for i:=1;i<=len(map1);i++{
		fmt.Println(i,map1[i])
	}

	/*
	step1：将所有的key取出--->切片/数组
	step2：进行排序
	step3：遍历key，-->map[key]
	 */
	 keys := make([]int,0,len(map1))
	 fmt.Println(keys)
	 for k :=range map1{
	 	keys = append(keys,k)
	 }
	fmt.Println(keys)
	//冒泡排序,或者使用sort包下的排序的方法
	sort.Ints(keys)
	fmt.Println(keys)
	for _,key :=range keys{
		fmt.Println(key,map1[key])
	}
	fmt.Println("------------------------")
	s1:=[]string{"Apple","Windows","Orange","abc","阿楠","acd","af","aaaaa"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)
}
