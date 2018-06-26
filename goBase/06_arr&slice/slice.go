package main

import (
	"fmt"
	"reflect"
)

func main() {


	var slice = []string{"a", "b", "c", "d"}

	fmt.Println(&slice[0])

	slice = removeString(slice, 2)
	fmt.Printf("%T\n", slice)
	fmt.Println(slice)
	fmt.Println(&slice[0],cap(slice))

	v := reflect.ValueOf("a")

	fmt.Println("qq",v)

	fmt.Println("------------")
	var slice2 = [][]string{}


	fmt.Println(slice2)

}


func removeString(s []string, i int) []string {
	return append(s[:i],s[i+1:]...)

}

