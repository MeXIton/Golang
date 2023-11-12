package main

import "fmt"

type employee struct {
	ID    string
	Name  string
	Phone string
}

func main() {
	employeeList := []employee{}
	employee1 := employee {
		ID: "101",
		Name: "Tonnam",
		Phone: "00000",
	}

	employee2 := employee {
		ID: "102",
		Name: "Ton",
		Phone: "000001",
	}
	

	employeeList = append(employeeList, employee1, employee2)
	fmt.Print("employee = ",employeeList)
}	

