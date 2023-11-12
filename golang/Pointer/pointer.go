package main

import "fmt"


func zero(ivalue int ) {
	ivalue = 0
}

func zeroPointer(ipointer *int) {
	*ipointer = 0 
}

func main() {
	i := 1
	fmt.Println("i = ", i)


	zero(i)
	fmt.Println("i from function zero = ",i)

	zeroPointer(&i)
	fmt.Println(" i from function zero = ",&i)
}
