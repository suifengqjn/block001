package main

import "fmt"

func main() {
	/*
	参数 的传递：值传递，引用传递
	 */
	 a := 10
	 fmt.Println("函数调用前：a，",a)
	 fun1(a)
	 fmt.Println("函数调用后：a,",a)

	 fun2(&a)
	 fmt.Println("函数fun2调用后：a,",a)

	 arr1:=[4]int{1,2,3,4}
	 fmt.Println("函数调用前：",arr1)//[1,2,3,4]
	 fun3(arr1)
	 fmt.Println("函数调用后：",arr1)//[1,2,3,4]

	 fun4(&arr1)
	 fmt.Println("函数fun4调用后：",arr1)


	 //s1:=[]int{1,2,3,4,5}
}



func fun5(s []int){

}

func fun4(p2 *[4]int){
	fmt.Println("函数中数组：",p2)
	p2[0] = 200
	fmt.Println("函数中数据：",p2)
}

func fun3(arr2 [4]int){//arr2 = arr1,数组是值数据
	fmt.Println("函数中数组：",arr2)//[1,2,3,4]
	arr2[0]= 100
	fmt.Println("函数中数组：",arr2)//[100,2,3,4]
}

func fun1(num int){//值传递  num = a=10
	fmt.Println("函数中，num：",num)
	num= 100
	fmt.Println("函数中修改num：",num)
}

func fun2(p1 *int){//传递的是a的地址，相当于引用传递
	fmt.Println("函数fun2,p1，",*p1)
	*p1 = 200
	fmt.Println("函数fun2,p1,",*p1)

}
