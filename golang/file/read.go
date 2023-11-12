package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("TH-CTF-2022_1.pdf")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line %d : %s\n", count, line)

		count++
	}
}
