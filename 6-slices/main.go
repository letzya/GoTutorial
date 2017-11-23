package main

import (
	"fmt"
)

func main() {

	slice := []string{}
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, "first1")
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, "second1")
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, "third1")
	fmt.Println(slice, len(slice), cap(slice))

	fmt.Printf("Array: %#v\n", slice)

	nilSlice()



	slices := []int{10,20,30,40,50,60}
	for key, value := range slices {
		fmt.Println(key, value)
	}


	fmt.Println("********")


	var nilArray []int
	fmt.Println("nilArray ", nilArray)
	myPrint(nilArray)

	var copyNilArray2 = nilArray
	fmt.Println("copyNilArray2 ", copyNilArray2)
	myPrint(copyNilArray2)

	copyNilArray2 = append(copyNilArray2, 5)
	fmt.Println("copyNilArray2 ", copyNilArray2)

	copyNilArray := nilArray

	fmt.Println("copyNilArray ", copyNilArray)
	myPrint(copyNilArray)

	emptySlice := []int{}

	fmt.Println("emptySlice ", emptySlice)
	myPrint(emptySlice)
	emptySlice = append(emptySlice, 5)
	myPrint(emptySlice)

	var emptyArray []int = []int{}
	fmt.Println("emptyArray ", emptyArray)
	myPrint(emptyArray)

	var a interface{}
	fmt.Println("interface ", a)

	a = nilArray
	fmt.Println("interface after assign ", a)
	if a == nil {
		fmt.Println("nil!" , a)
	} else {
		fmt.Println("not nil!", a)
	}

	a = append(nilArray, 5,4)
	fmt.Println("nilArray ", nilArray)
	fmt.Println("interface ", a)
	//myPrint( a)

}

func myPrint(z []int) {
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!" , z)
	} else {
		fmt.Println("not nil!", z)
	}
}

func nilSlice(){

	var s []int

	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil slice")
	}
}