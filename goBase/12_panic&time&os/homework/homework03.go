package main

import "fmt"

//1.定义一个发动机
type IEngine interface {
	start()
	speedUp()
	stop()
}

type YAMAHA struct {
	name string
}
func (y YAMAHA)start(){
	fmt.Println(y.name,"启动，速度60迈。。")
}
func (y YAMAHA)speedUp(){
	fmt.Println(y.name,"加速，速度80迈。。")
}
func (y YAMAHA)stop(){
	fmt.Println(y.name,"停止，速度0迈。。")
}

type HONDA struct {
	name string
}
func (h HONDA)start(){
	fmt.Println(h.name,"启动，速度40迈。。")
}
func (h HONDA)speedUp(){
	fmt.Println(h.name,"加速，速度120迈。。")
}
func (h HONDA)stop(){
	fmt.Println(h.name,"停止，速度0迈。。")
}

type Car struct {
	iEngine IEngine //接口类型作为的字段属性
}


func (c Car) testIEngine(){
	c.iEngine.start()
	c.iEngine.speedUp()
	c.iEngine.stop()
}
func main() {
	/*
	1.综合题：

定义一个IEngine接口，3个方法：start(),speedup(),stop()，表示启动，加速，停止
定义两个结构体实现该接口：YAMAHA和HONDA
实现方式分别是：
	YAMAHA
		启动：60，加速80，停止0

	HONDA
		启动：40，加速120，停止0

定义一个Car结构体，将IEngine作为字段，再提供一个car的方法：testIEngine()，用于测试该小汽车的发动机。也就是在testIEngine()中调用start(),speedup(),stop()方法。
	 */

	 c1 := Car{}
	 //c1.iEngine = YAMAHA{"雅马哈"}
	 c1.iEngine = HONDA{"本田"}
	 c1.testIEngine()
}
