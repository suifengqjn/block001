package main

import "fmt"

func main() {


	for i := 1; i <= 5; i++ {

		for j := 1;j <=5 ;j++  {
			if i >= j {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}

	fmt.Println("---------")


	for i := 1; i <= 5; i++ {

		for j := 5;j >=1 ;j--  {
			if i >= j {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}


	fmt.Println("---------")

	for i := 1; i <= 5; i++ {

		for j := 1;j <= 5 ;j++  {
			if i <= j {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}

	fmt.Println("---------")

	for i := 1; i <= 5; i++ {

		for j := 1;j <= 5 ;j++  {
			if i >= j {
				fmt.Print(" ")
			}

		}
		for j := 1;j <= 5 ;j++  {

			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("---------")

	for i := 1; i <= 5; i++ {

		for j := 1;j <= 5 ;j++  {
			if i <= j {
				fmt.Print(" ")
			}

		}
		for j := 1;j <= 5 ;j++  {

			fmt.Print("*")
		}
		fmt.Println()
	}
	
}
