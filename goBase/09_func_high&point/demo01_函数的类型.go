package main

import "fmt"

func main() {
	/*
	数据类型的分类：
		基本类型：
			int，float，bool，string
		复合类型
			array，slice，map，function

	数据类型的分类：
		值类型：
			int，float，bool，string，array，struct
		引用类型
			slice，map，function

	%T
	 */
	 a := 10
	 fmt.Printf("%T\n",a)//int
	 b :=[4]int{1,2}
	 fmt.Printf("%T\n",b)//[4]int
	 /*
	 [4]int,
	 [8]float64,
	 [5]string
	 [3]map[string]map[int][string]
	  */
	 c :=[]string{}
	 fmt.Printf("%T\n",c)//[]string

	 d:=make(map[int]string)
	 fmt.Printf("%T\n",d) //map[int]string

	 fmt.Printf("%T\n",fun1) //func(int, int)
	 fmt.Printf("%T\n",fun2) //func(float64, float64, string, int)
	 fmt.Printf("%T\n",fun3) //func(string, string) int
	 fmt.Printf("%T\n",fun4) //func(float64,int)(string,int)
}


func fun1(a, b int){//func(int, int)

}
//func(float64,float64,string,int)
func fun2(a, b float64,c string, d int){

}

func fun3(a, b string)(int){
	return 0
}

func fun4(a float64,b int)(string, int){
	return "",0
}

//func (int,int,float64,string, string)(float64,int, string,string)
func fun5(a,b int,c float64,d,e string)(float64,int, string,string){

}


