package main

import "fmt"

type Book struct {
	bookName string
	price float64
}
type Student struct {
	name string
	age int
	books []*Book
}
/*
array,定长
slice，变长
map,key，value
 */

func main() {
	//1.创建书
	b1 := Book{"孙子兵法",120.8}
	b2 := Book{"易筋经",68.9}
	b3 := Book{"西游记",88.9}
	//2.创建书架，切片
	bb1:=make([]*Book,0,10)
	//3.将书，添加到书架切片上
	bb1 = append(bb1,&b1,&b2,&b3)
	//4.创建学生对象
	s1:=Student{"王二狗",19,bb1}
	//5.获取对象的属性
	fmt.Printf("姓名：%s，年龄：%d\n",s1.name,s1.age)
	//fmt.Println(s1.books)
	for i:=0;i<len(s1.books);i++{
		p := s1.books[i] //书对象的地址
		fmt.Printf("\t第%d本书：书名：%s,价格：%.2f\n",i+1,(*p).bookName,p.price)
	}

	s2 :=Student{"李小花",18,make([]*Book,0,10)}
	//s2.books--->切片
	s2.books = append(s2.books,&Book{"红楼梦",77.7},&Book{"三国演义",88.5})

	fmt.Println(s2.name,s2.age)
	//fmt.Println(s2.books)
	for i:=0;i<len(s2.books);i++{
		p:=s2.books[i] //书的指针
		fmt.Printf("\t书名：%s,价格：%.2f\n",p.bookName,p.price)
	}
	fmt.Println("---------------------------")
	//创建一个容器，存储学生对象
	ss1:=make([]*Student,0,10)
	ss1 = append(ss1,&s1,&s2)
	for i:=0;i<len(ss1);i++{
		fmt.Printf("第%d个学生的信息：\n",i+1)
		fmt.Printf("\t学生姓名：%s,学生年龄：%d\n",ss1[i].name,ss1[i].age)
		bb := ss1[i].books // 装书的容器：切片
		if len(bb) == 0{
			fmt.Println("\t\t该学生不看书。。")
		}else{
			//遍历
			for j :=0;j<len(bb);j++{
				p :=bb[j] //书架中书的地址
				fmt.Printf("\t\t第%d本书：书名：《 %s 》，价格：%.2f\n",j+1,ss1[i].books[j].bookName,p.price)
			}
		}
	}


}


