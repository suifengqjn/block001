package main

import "fmt"

//1.定义一个书的结构体
type Book struct {
	bookName string
	price    float64
	author   string
}

//2.定义一个人的结构体：姓名，年龄，书
type Person struct {
	name string
	age  int
	book Book//模拟聚合关系：has a
}

//p1.book = b1

type Person2 struct{
	name string
	age int
	book *Book
}

//p1.book = &b1

func main() {
	b1 := Book{}
	b1.bookName = "金瓶里没有梅"
	b1.price = 45.8
	b1.author = "田中天"
	p1 := Person{}
	p1.name = "王二狗"
	p1.age = 30
	p1.book = b1 //值传递

	fmt.Printf("姓名：%s，年龄：%d，书名：%s，价格：%.2f,作者：%s\n",p1.name,p1.age,p1.book.bookName,p1.book.price,p1.book.author)
	p1.book.bookName = "金瓶里没有梅2"
	fmt.Println(p1)
	fmt.Println(b1)


	p2 :=Person{name:"老王",age:28,book:Book{bookName:"呼啸山庄",price:89.7,author:"艾米莉·勃朗特"}}
	fmt.Println(p2.name,p2.age)
	fmt.Println("\t",p2.book.bookName,p2.book.price,p2.book.author)

	p3:=Person{
		name:"Jerry",
		age:26,
		book:Book{
			bookName:"Go语言是怎样练成的",
			price:55.9,
			author:"王建",
		},
	}
	fmt.Println(p3.name,p3.age)
	fmt.Println("\t",p3.book.bookName,p3.book.price,p3.book.author)

	p4:=Person{"李小花",20,Book{"十万个为啥",44.8,"张晓坤"}}
	fmt.Println(p4.name,p4.age,p4.book.bookName,p4.book.price,p4.book.author)
	/*
	练习1：创建一个student类，book类
			一个学生对象，可以有多本书。
	 */


}
