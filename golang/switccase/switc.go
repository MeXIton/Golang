package main

import (
	"fmt"
)

var input string
var colorCode string

func main() {
	fmt.Scan(&input)
	switch input {
	case "blue":
		fmt.Println("#0000FF")

	case "Green":
		fmt.Println("#008000")

	case "Yellow":
		fmt.Println("#FFF00")

	default:
		fmt.Println("No color")
	}
}
