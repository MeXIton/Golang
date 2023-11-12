package main

import "fmt"

var number, number2 int = 1000, 1001
var Msg string = "Hello"

func main() {
	// Shot  format  //
	numberTonnam := 25.4
	fmt.Println(number, number2)
	fmt.Println(numberTonnam)
	fmt.Println(Msg)

	fmt.Println(number + number2)
	fmt.Println(numberTonnam + float64(number))

	fmt.Println(Msg + "TonnamNM")

	fmt.Println("My money =", numberTonnam)
}
