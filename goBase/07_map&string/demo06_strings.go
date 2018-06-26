package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	strings包下的关于字符串的函数
	A：大小写转换
	ToLower()
	ToUpper()

	B：切割和拼接
	Split(s,sep)-->[]string
	SplitN(s,sep,n)-->[]string

	Join([]string, sep)-->string

	C：查找
	以xx前缀开头，以xx后缀结尾
	HasPrefix(s,prefix)-->bool
	HasSuffix(s,suffix)-->bool

	是否包含
	Contains(s, substr)-->bool，是否包含指定的内容
	ContainsAny(s,chars)-->bool,是否包含chars中任意一个字符即可

	索引
	Index(s,substr)-->int,返回substr在s中第一次出现的位置，如果不存在返回-1
	IndexAny(s,chars)-->int,返回chars中的任意一个字符在s中第一次出现的位置，如果不存在，返回-1
	LastIndex(s,substr)-->int,返回substr最后一次在s中出现的位置
	LastIndexAny()

	统计次数：
	Count(s, substr)-->int,子串substr，在s中出现的次数

	替换：
	Replace(s,old,new,n)-->string,

	重复：
	Repeat(s,count)-->string,

	去除首尾指定的内容：
	Trim()
	TrimSpace()

	截取子串：
	s[start:end]-->substr
		包含start，不包含end

	 */
	 s1 := "helloworld"
	 fmt.Println(strings.Contains(s1,"abc"))
	 fmt.Println(strings.ContainsAny(s1,"hello"))

	 fmt.Println(strings.Count(s1,"lloo"))


	 //以xxx前缀开头，以xx后缀结尾
	 s2 := "20180522课堂笔记.txt"
	 if strings.HasPrefix(s2,"201805"){
	 	fmt.Println("18年5月的文件")
	 }

	 if strings.HasSuffix(s2,".txt"){
	 	fmt.Println("文本文档。。")
	 }

	 //索引
	 //helloworld
	 fmt.Println(strings.Index(s1,"l")) //查找substr在s中的位置，如果不存在返回-1
	 fmt.Println(strings.IndexAny(s1,"abcdh")) //查找chars中任意一个字符，在s中的位置，
	 fmt.Println(strings.LastIndex(s1,"l")) //查找substr在s中最后一次出现的位置

	 //字符串的拼接
	 ss1:=[]string{"abc","world","hello","rose"}
	 s3:=strings.Join(ss1,"*")
	 fmt.Println(s3)

	 //切割
	 s4 := "123,475858,dd,,30404,ddd,,"
	 ss2:=strings.Split(s4,",")
	 for i:=0;i<len(ss2);i++{
	 	fmt.Println(ss2[i])
	 }
	//按照数量切割，切割后数量不能超过n，如果想切割所有，设置n为-1
	 ss3 := strings.SplitN(s4,",",20)
	 for i:=0;i<len(ss3);i++{
	 	fmt.Println(ss3[i])
	 }

	 //重复,自己拼接自己count次
	 fmt.Println(strings.Repeat("hello",5))

	 //替换:helloworld
	 fmt.Println(s1)
	 fmt.Println(strings.Replace(s1,"l","*",-1))

	 //大小写转换
	 s5 := "heLLO wORld**123"
	 fmt.Println(strings.ToLower(s5))
	 fmt.Println(strings.ToUpper(s5))

	 //去除首尾指定的内容
	 s6 :=" *++*hello**world**** "
	 fmt.Println(strings.Trim(s6,"*+ "))
	 fmt.Println(strings.TrimLeft(s6,"+* "))
	 fmt.Println(strings.TrimRight(s6,"+* "))
	 fmt.Println(strings.TrimSpace(s6)) //直接去除首尾空格


	 //截取子串
	 /*
	 substring(start,end)-->substr
	  */
	  fmt.Println(s1) //helloworld
	  s7 := s1[5:]
	  fmt.Println(s7)
	  s8:=s1[2:5]
	fmt.Println(s8)
}
