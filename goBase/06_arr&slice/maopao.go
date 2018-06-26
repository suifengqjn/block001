package main

import "fmt"

func main()  {


	var arr = [...]int{2,43, 10, 7, 8, 5, 20, 11}

	// 冒泡排序
	for i := 0;i<len(arr) ;i++  {
		for j:= 0;j <len(arr) - 1 ; j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}

	fmt.Println(arr)


	//选择排序
	arr = [...]int{2,43, 10, 7, 8, 5, 20, 11}
}
