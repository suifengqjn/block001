package main

import (
	"sort"
	"fmt"
)

//使用sort包对任意类型元素的集合进行排序
type Person struct {
	name string
	age  int
}

func (p *Person)String()string{
	return fmt.Sprintf("{%s,%d}",p.name,p.age)
}

type PersonSlice []*Person

//定义排序的规则：
func (slice PersonSlice) Less(i, j int) bool {
	if slice[i].age < slice[j].age {
		return true
	} else if slice[i].age > slice[j].age {
		return false
	} else {
		return slice[i].name < slice[j].name
	}
}
//交换指定下标i，j的数据
func (slice PersonSlice) Swap(i,j int){
	slice[i],slice[j] = slice[j],slice[i]
}
//要排序的容器中的对象的个数：len（容器名）
func (slice PersonSlice) Len()int  {
	return len(slice)
}

func main()  {
	//自定义排序：给Person的对象
	//规则：年龄从小到大，年龄相同按照姓名
	p1 :=Person{"aaa",30}
	p2 :=Person{"bbb",29}
	p3 :=Person{"ccc",31}
	p4 :=Person{"ddd",29}
	p5 :=Person{"eee",35}

	s1:=make([]*Person,0,10)
	s1 = append(s1,&p1,&p2,&p3,&p4,&p5)
	//s1 := PersonSlice([]*Person{&p1,&p2,&p3,&p4,&p5})
	s2:=PersonSlice(s1)
	sort.Sort(s2)

	fmt.Println(s2)



}
