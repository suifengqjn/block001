package main

import "fmt"

func main()  {
	/*
	算术运算符：+，-，*，/,%,++,--
	/,取商
	%,取余，取模

	++，自赠1
	--，自减1
	 */
	 a := 10
	 b := 3
	 mul := a * b
	 fmt.Println(mul)
	 div := a / b //取商
	 mod := a % b//取余，取模
	 fmt.Println(div,mod)

	 m := 3
	 fmt.Println(m) //3
	 m++ //自增加1
	 fmt.Println(m)//4
	 m-- //自减1
	 fmt.Println(m)

}
/*
8/5=1
8%5=3

3/4=0
3%4=3

4/3=1
4%3=1

2/1=2
2%1=0

1/7=0
1%7=1

6/3=2
6%3=0

 */
