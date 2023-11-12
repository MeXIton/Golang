package main

import "fmt"

var score int 

func main() {
	fmt.Println("Grade")
	fmt.Scanf("%d",&score)
	fmt.Println("score = ",score)
	if score >= 80 {
		fmt.Println("Grade A")
	} else if score >=70  {
		fmt.Println("Grade B")
	} else if score >= 60 {
		fmt.Println("Grade C")
	} else if score >= 50 {
		fmt.Println("Grade D")
	} else {
		fmt.Println("Grade F")
	}
}
