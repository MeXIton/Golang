package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID       int
	Employee string
	Tel      string
	Email    string
	Password string
}

func main() {

	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":101,"Employee":"Tonnam","Tel":"0646253202","Email":"Naruepai@gmail.com","Password":"TonnamNamton"}`), &e )
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)
	fmt.Println(e.Password)
}
