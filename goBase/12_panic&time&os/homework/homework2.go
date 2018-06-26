package main

import "fmt"

type Student struct {
	name string
	age int
	sex string
	englishScore, chineseScore,mathScore float64
}
func (s Student) getSum()float64{
	sum :=s.englishScore+s.chineseScore+s.mathScore
	return sum
}
func (s Student) getAvg()float64{
	avg := s.getSum() / 3
	return avg
}

func (s Student) printInfo(){
	fmt.Printf("姓名：%s,年龄：%d,性别：%s\n",s.name,s.age,s.sex)
	fmt.Printf("\t语文成绩：%.2f,英语成绩：%.2f,数学成绩：%.2f\n",s.chineseScore,s.englishScore,s.mathScore)
	fmt.Printf("\t总分%.2f，平均分：%.2f\n",s.getSum(),s.getAvg())
}
func main() {
	/*
	定义一个学生类，6个属性：姓名，年龄，性别，英语成绩，语文成绩，数学成绩，提供3个方法，求总分(),求平均分(),打印信息()
	 */

	 s1 := Student{"王二狗",18,"男",88.5,69.7,88}
	 //fmt.Println(s1.getSum(),s1.getAvg())
	 s1.printInfo()
}
