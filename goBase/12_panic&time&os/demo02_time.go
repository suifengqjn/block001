package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	/*
	time包：
	1年=365天
		day
	1天=24小时hour
	1小时=60分钟minute
	1分钟=60秒
		second
	1秒=1000毫秒
		millisecond
	1毫秒=1000微秒
		microsecond-->μs
	1μs = 1000纳秒
		nanosecond-->ns
	1ns = 1000皮秒
		picosecond-->ps


	A:获取时间time对象
		time.Now()当前时间
		time.Date()指定时间
	B：格式化：
		time-->string
			t1.Format(layout string)-->string
		string-->time
			time.Parse(layout string, value string)-->time,error

	C：时间戳：timeStamp
		t1.Unix()
		t1.UnixNano()

	D：获取指定内容
		t1.Date()-->year,month,day
		t1.Clock()-->hour,minute,second
		t1.Year()-->year
		t1.Month()-->Month
		t1.Day()-->day
		t1.Week()
		t1.Hour()
		t1.Minute()
		t1.Second()
		t1.NanoSecond()

	E：时间间隔：Duration
		type Duration int64

		t1.Add(d)-->time
		t1.Sub(t2)-->d

	F：睡眠
		time.Sleep(duration)//NanoSecond



	 */
	 //1.获取当前的时间
	 t1:=time.Now()
	 fmt.Printf("%T\n",t1)// time.Time
	 fmt.Println(t1)//2018-05-29 10:58:23.789029492 +0800 CST m=+0.001181079
	 //2.获取指定的时间
	 t2 :=time.Date(2008,7,15,16,30,28,0,time.Local)
	 fmt.Println(t2)

	 //3.time--string之间的转换
	 //fmt.Println(t1.String())
	 /*
	 t1.Format("格式木板")-->string
	 	模板的日期必须是固定：6-1-2-3-4-5

	 time.parse("模板",str)-->time

	  */
	 s1 := t1.Format("2006年1月2日 15:04:05")
	 fmt.Println(s1)

	 s2 :=t1.Format("2006-01-02")
	 fmt.Println(s2)

	 s3 := "1999年10月10日" //string
	 t3,err:=time.Parse("2006年1月2日",s3)
	 if err != nil{
	 	fmt.Println("err:",err.Error())
	 }
	 fmt.Println(t3)

	 //4.时间戳：指定的日期(当前日期，指定的某个日期)，距离1970年1月1日0点0时0分0秒的时间差值：秒，纳秒
	 t4 := time.Date(1970,1,1,1,0,0,0,time.UTC)
	 timeStamp1:=t4.Unix() //秒的差值
	 fmt.Println(timeStamp1)//3600
	 timeStamp2 :=t1.Unix()
	 fmt.Println(timeStamp2)

	 timeStamp3 :=t4.UnixNano()//纳秒差值
	 fmt.Println(timeStamp3)//3600 000 000 000
	 timeStamp4:=t1.UnixNano()
	 fmt.Println(timeStamp4)

	 //5.根据当前的时间获取指定的内容：年，月，日，
	 year,month,day:=t1.Date()
	 fmt.Println(year,month,day)
	 hour,min,sec:=t1.Clock()
	 fmt.Println(hour,min,sec)

	 year2:=t1.Year()
	 fmt.Println("年：",year2)
	 fmt.Println(t1.YearDay())
	 month2:=t1.Month()
	 fmt.Println("月：",month2)
	 fmt.Println("日：",t1.Day())
	 fmt.Println("小时：",t1.Hour())
	 fmt.Println("分钟：",t1.Minute())
	 fmt.Println("秒：",t1.Second())
	 fmt.Println(t1.Weekday())//星期几
	 fmt.Println(t1.ISOWeek()) //年份，和该年份中的第几个星期
	 //t1.Nanosecond()

	 //5.时间间隔

	 t5 := t1.Add(time.Minute)//duritaion时间间隔之前/后
	 fmt.Println(t1)
	 fmt.Println(t5)
	 fmt.Println(t1.Add(24*time.Hour)) //1天之后

	 t6:=t1.AddDate(1,0,0)
	 fmt.Println(t6)

	 fmt.Println(t1.Add(-5*time.Minute))

	 d1:=t5.Sub(t1) //两个time的时间差值： Duration
	 fmt.Printf("%T\n",d1)
	 fmt.Println(d1)

	 //6.睡眠
	 time.Sleep(3*time.Second)//让当前的程序进入睡眠状态
	 fmt.Println("main...over.....")

	 //睡眠[1-10]的秒随机数
	 rand.Seed(time.Now().UnixNano())
	 randNum := rand.Intn(10)+1 //int
	 fmt.Println(randNum)
	 time.Sleep(time.Duration(randNum)*time.Second)//Duration
	 fmt.Println("睡醒了。。")



}
