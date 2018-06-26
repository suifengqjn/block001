package main

import "fmt"

func main() {
	/*
	二维数组：存储的是一维的一维
		arr :=[3][4]int
		该二维数组的长度，就是3。
		存储的元素是一维数组，一维数组的元素是数值
	 */

	 arr :=[3][4]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}//[0-2][0-3]
	 fmt.Println(arr)
	 fmt.Println("二维数组的长度：",len(arr))//一维的个数
	 fmt.Println("一维数组的长度：",len(arr[0]))//一维中数值的个数
	 fmt.Println(arr[0][0])//1
	 fmt.Println(arr[2][2])//11
	 fmt.Println(arr[2][1])//10
	 fmt.Println(arr[1][3])//8
	 //遍历二维数组
	 for i:=0;i<len(arr);i++{
	 	for j:=0;j<len(arr[i]);j++{
	 		fmt.Print(arr[i][j],"\t")
		}
		fmt.Println()
	 }
	 //for range遍历二维数组
	 for _,arr1:=range arr{
	 	//fmt.Println(arr1)
	 	for _,val :=range arr1{
	 		fmt.Print(val,"\t")
		}
		fmt.Println()
	 }




}
