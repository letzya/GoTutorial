package main

import (
	"fmt"
)

var a = "Hello"
var b bool
var c int

type myType struct {
	MyString string
	MyInt    int
}

type myInterface struct {
	MyString string
	MyInt    int
}

func main() {

	var d string
	var e int = 3

	f := 99

	g := myType{"my string", 33}

	fmt.Printf("%T: %v\n", a, a)
	fmt.Printf("%T: %v\n", b, b)
	fmt.Printf("%T: %v\n", c, c)
	fmt.Printf("%T: %v\n", d, d)
	fmt.Printf("%T: %v\n", e, e)
	fmt.Printf("%T: %v\n", f, f)

	fmt.Printf("%T: %v\n", g, g)


	switchCase(a)
	switchCase(b)
	switchCase(c)
	switchCase(g)
	switchCase( )
}

func switchCase(myInterface interface{}) {
	switch v := myInterface.(type) {
	case int:
		fmt.Printf("v is an int - %T: %v\n", v, v)
	case string:
		fmt.Printf("v is a string - %T: %v\n", v, v)
	case bool:
		fmt.Printf("v is a bool - %T: %v\n", v, v)
	case myType:
		fmt.Printf("v is a myType - %T: %v\n", v, v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}