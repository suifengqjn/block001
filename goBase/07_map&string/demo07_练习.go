package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	练习2：给定一个路径名：
pathName:="http://192.168.10.1:8080/Day33_Servlet/aa.jpeg"
获取文件名称：aa.jpeg

	给定一个以下字符串：统计大写字母的个数，小写字母的个数，非字母的个数。
str:="aekjffjkJDJ294384848DKFJFJkdjfhfh2943845593nfnJRIEIFJ"
	 */
	pathName := "http://192.168.10.1:8080/Day33_Servlet/aa.jpeg"

	//index:=strings.LastIndex(pathName,"/")
	//fmt.Println(index)
	fileName := pathName[strings.LastIndex(pathName, "/")+1:]
	fmt.Println(fileName)

	str := "aekjffjkJDJ294384848DKFJFJkdjfhfh2943845593nfnJRIEIFJ"
	count1 := 0 //大写字母个数 [65,90]
	count2 := 0 //小谢字母个数 [97,122]
	count3 := 0 //非字母个数

	for i := 0; i < len(str); i++ {
		if str[i] >= 65 && str[i] <= 90 {
			count1++
		} else if str[i] >= 97 && str[i] <= 122 {
			count2++
		} else {
			count3++
		}
	}
	fmt.Printf("大写字母：%d，小写字母：%d，非字母：%d\n", count1, count2, count3)

	count1 = 0
	count2 = 0
	count3 = 0

	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			count1++
		} else if str[i] >= 'a' && str[i] <= 'z' {
			count2++
		} else {
			count3++
		}
	}

	fmt.Printf("大写字母：%d，小写字母：%d，非字母：%d\n", count1, count2, count3)
	s1 := "qwertyuiopasdfghjklzxcvbnm" //小写
	s2 := "QWERTYUIOPASDFGHJKLZXCVBNM" //大写

	count1 = 0
	count2 = 0
	count3 = 0
	for i := 0; i < len(str); i++ { //str[i]
		if strings.ContainsRune(s2, rune(str[i])) {
			count1++
		} else if strings.ContainsRune(s1, rune(str[i])) {
			count2++
		} else {
			count3++
		}
	}
	fmt.Printf("大写字母：%d，小写字母：%d，非字母：%d\n", count1, count2, count3)
	count1 = 0
	count2 = 0
	count3 = 0

	for i := 0; i < len(str); i++ {
		if strings.IndexByte(s2, str[i]) >= 0 {
			count1++
		} else if strings.IndexByte(s1, str[i]) >= 0 {
			count2++
		} else {
			count3++
		}
	}
	fmt.Printf("大写字母：%d，小写字母：%d，非字母：%d\n", count1, count2, count3)

	s3 := strings.ToLower(str) //转为小写
	s4 := strings.ToUpper(str) //转为大写
	count1 = 0
	count2 = 0
	count3 = 0
	for i := 0; i < len(str); i++ {
		if str[i] != s3[i] {
			count1++
		}
		if str[i] != s4[i] {
			count2++
		}
	}
	count3 = len(str) - count1 - count2
	fmt.Printf("大写字母：%d，小写字母：%d，非字母：%d\n", count1, count2, count3)

	a := '0'//字符9-->编码
	b := 9 //数值9
	fmt.Println(a, b)

	count1=0
	count2=0
	count3=0
	sum:=0//原串的编码总和
	sum1:=0//求小写字母的编码总和
	sum2:=0//求大写
	for i:=0;i<len(str);i++{
		sum += int(str[i])
		sum1+=int(s3[i])
		sum2+=int(s4[i])
	}
	count1 = (sum1-sum) / 32
	count2 = (sum-sum2) / 32
	count3 = len(str)-count1-count2
	fmt.Printf("大写字母：%d，小写字母：%d，非字母：%d\n", count1, count2, count3)

}
