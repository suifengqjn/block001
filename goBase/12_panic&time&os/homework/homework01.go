package main

import (
	"fmt"
	"math"
)

type Point struct {
	x,y,z float64
}

func (p Point) printInfo(){
	fmt.Printf("x轴：%.2f,y轴：%.2f,z轴：%.2f\n",p.x,p.y,p.z)
}

func (p Point)getDistance1(x1,x2,y1,y2,z1,z2 float64) float64{
	dis := math.Sqrt(math.Pow(x1-x2,2)+math.Pow(y1-y2,2)+math.Pow(z1-z2,2))
	return dis
}
func (p Point) getDistance2(p1,p2 Point)float64{
	dis :=math.Sqrt(math.Pow(p1.x-p2.x,2)+math.Pow(p1.y-p2.y,2)+math.Pow(p1.z-p2.z,2))
	return dis
}

func (p Point)getDistance3(p2 Point)float64{
	dis :=math.Sqrt(math.Pow(p.x-p2.x,2)+math.Pow(p.y-p2.y,2)+math.Pow(p.z-p2.z,2))
	return dis
}

func main() {
	/*
	定义一个三维空间的点(Point),设置三个属性分别表示x轴的值，y轴的值，z轴的值。提供一个方法，用于求两点之间的距离
	 */
	 p1:=Point{2,4,3}
	 p2:=Point{0,0,0}

	 p1.printInfo()
	 p2.printInfo()

	 res1:= p1.getDistance1(p1.x,p2.x,p1.y,p2.y,p1.z,p2.z)
	 fmt.Println(res1)

	 res2 := p1.getDistance2(p1,p2)
	 fmt.Println(res2)

	 res3:= p1.getDistance3(p2)
	 fmt.Println(res3)

	 res4:=p2.getDistance3(p1)
	 fmt.Println(res4)
}
