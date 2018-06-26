package main

import "fmt"

func main(){
	/*
	 位运算符：了解性
		将数值，转为二进制后，按位操作
		按位&，
			位的值如果都为1才为1,有一个为0就为0
		按位|，
			位的值如果都为0才为0,有一个为1就为1
		异或^，
			不同为1,相同为0

	位移运算符：
	<<：按位左移,将a转为二进制，向左移动b位
		a << b
	>>：按位右移，将a转为二进制，向右移动b位
		a >> b
	 */

	 a := 60
	 b := 13
	 /*
	 60:32+16+8+4
	 111100
	 13:8+4+1
	 1101

	 60:0011 1100
	 13:0000 1101
	 & :0000 1100
	 | :0011 1101 32+16+8+4+1
	 ^ :0011 0001 32+16+1
	  */
	  fmt.Printf("%d,%b\n",a,a)
	  fmt.Printf("%d,%b\n",b,b)

	  res1 := a & b
	  fmt.Println(res1) //12
	  res2 :=a | b
	  fmt.Println(res2)
	  res3 :=a ^ b
	  fmt.Println(res3)

	  c := 8 //
	  /*
	  c :100
	  0000 0100
	   */
	   res4 := c << 2
	   fmt.Println(res4)
	   res5:=c >> 2
	   fmt.Println(res5)

}
