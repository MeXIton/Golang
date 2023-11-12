package main

import "fmt"

const count = 10

func main() {

	for i := 10; i < count ; i++ {
		fmt.Println("i = ", i)
	}
	var input string
	for {
		fmt.Scanf("%s",&(input))
		fmt.Println("input = ",input)

		if input == "exit"{
			break
		}
	}

	
}