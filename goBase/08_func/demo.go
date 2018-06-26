package main

import "fmt"

func main() {
	s1:=[]int{1,2,3,4}
	s2 :=[] string{"a","b","c"}
	test1(s1)
	test1(s2)
}
func test1(a interface{})  {
	fmt.Println(a)
}
