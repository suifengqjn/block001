package main

import "fmt"

//空接口
type A interface {

}

type Cat struct {
	name string
	age int
}

type Person struct {
	name string
	sex string
}


//空接口：理解为代表了任意类型
func test1(a A){

}
//将匿名空接口作为参数：表示该函数可以接受任意类型的数据
func test2(a interface{}){

}

func test3(slice2 []interface{}){
	for i:=0;i<len(slice2);i++{
		fmt.Println("第",i+1,"个数据：")
		switch ins:=slice2[i].(type) {
		case Cat:
			fmt.Println("\tcat对象：",ins.name,ins.age)
		case Person:
			fmt.Println("\tperson对象：",ins.name,ins.sex)
		case int:
			fmt.Println("\tint类型：",ins)
		case string:
			fmt.Println("\tstring类型：",ins)
		}
	}
}

func main() {
	/*
	空接口：也是一个接口，但是该接口中没有任何的方法。
		所以可以将任意类型作为该接口的实现。
	 */

	 var a1 A = Cat{name:"花猫",age:1}
	 var a2 A = Person{"王二狗","男性"}
	 var a3 A = "haha"
	 var a4 A = 100

	test1(3.14)
	test2(a1)

	//定义一个map：string作为key，任意类型作为value
	map1 := make(map[string]interface{})
	map1["name"] = "王二狗"
	map1["age"] = 30
	fmt.Println(map1)
	fmt.Println(100,"hahah","helo",Cat{"小天",2})


	//定义一个切片，存储任意类型的数据
	slice1 := make([]interface{},0,10)
	slice1 = append(slice1,a1,a2,a3,a4)
	fmt.Println(slice1)

	test3(slice1)


}
