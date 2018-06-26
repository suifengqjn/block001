package main

import "fmt"

//1.定义一个结构体，表示自定义错误的类型
type errorRect struct {
	msg string//错误的信息
	wid float64//宽度
	len float64//高度
}

func (e *errorRect) Error()string{
	return fmt.Sprintf("宽度：%.2f,长度：%.2f,错误信息：%s",e.wid,e.len,e.msg)
}


//2.定义一个函数，用于求矩形的面积
func getArea(wid,len float64)(float64,error){
	errorMsg:=""
	if wid < 0{
		errorMsg = "宽度为负数"
	}
	if len < 0{
		if errorMsg == ""{
			errorMsg = "长度为负数"
		}else{
			errorMsg += ",长度也为负数"
		}
	}
	if errorMsg != ""{
		return 0,&errorRect{errorMsg,wid,len}
	}
	area := wid * len
	return area,nil

}


func main() {
	/*
	自定义error
		error是一个内置的接口interface，定义了一个方法
			Error() string

	 */

	 res1,err:=getArea(-4,-6)
	 if err != nil{
	 	fmt.Printf(err.Error())
	 }else{
	 	fmt.Println("面积是:",res1)
	 }
}
