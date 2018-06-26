package main

import "fmt"

func main()  {
	/*
	go语言是静态语言，定义，赋值必须类型一致。
		运算：+，-，=....

	数据类型转换：type convert
		兼容类型之间可以转换:int,float
		语法格式：Type(Value)

	常数，在有需要的时候，可以自动扎u内环
	变量，需要手动转换
	 */
	 var a int8 = 10
	 var b int16 = 20
	 sum := int16(a) + b
	 fmt.Printf("%T,%d\n",sum,sum)

	 f1 := 3.74
	 sum2 := f1 + 100
	 fmt.Printf("%T,%f\n",sum2,sum2)
	 c := 100 //int
	 sum3 := int(f1) + c
	 fmt.Printf("%T,%d\n",sum3,sum3)
}
