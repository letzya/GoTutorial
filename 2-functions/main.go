package main

import "fmt"

func main() {
	fmt.ln(multifly(10, 3))

	name, age := GetNameAndAge()
	fmt.ln("Name: ", name, " Age:", age)
}

func multifly(x int, y int) int {
	return x * y
}

//Public
func GetNameAndAge() (string, int) {
	return "Ahmet", 35
}