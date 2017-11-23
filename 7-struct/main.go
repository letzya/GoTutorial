package main

import "fmt"

func main() {
	tyk := Company{
		"Tyk Technology",
		[]Employee{
			{"yaara", 555, "developer", 1000.0},
			{"Josh", 29, "developer", 444},
		},
	}

	fmt.Println("%#v", tyk )
}

type Employee struct {
	Name string
	Age int
	JobTitle string
	salary float64 //this is private member
}


type Company struct {
	Name string
	Employees []Employee
}