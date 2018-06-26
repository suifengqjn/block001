package main

import "fmt"

func main() {
	/*
		双层for循环
	Print()
	Println()
	Printf()
	 */
	// i是用于控制行数
	/*
	for i:=1;i<=5;i++{ //i=1,2,3,4,5
		//j是用于控制*的个数
	   for j:=0;j<5;j++{ //j=0,1,2,3,4
		   fmt.Print("*")
	   }
	   fmt.Println()
   }
	*/

	out:for i := 1; i < 5; i++ {

		fmt.Println("i",i)

		for j := 1; j <5 ;j++ {

			if j ==2 {
				break out
			}
			fmt.Println("j",j)
		}

	}

}

/*		行	*	换行
*****	1	5	1
*****	2	5	1
*****	3	5	1
*****	4	5	1
*****	5	5	1


*
**
***
****
*****

    *
   **
  ***
 ****
*****

    *****
   *****
  *****
 *****
*****
 */
