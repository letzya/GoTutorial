package main

import (
	"fmt"
	"encoding/json"
	"image/color"
	"reflect"
)

func main() {
	tyk := Company{
		"Tyk Technology",
		[]Employee{
			{"yaara", 555, "developer", 1000.0},
			{"Josh", 29, "developer", 444},
		},
	}

	fmt.Println("tyk object %#v", tyk )


	jsBytes, _ := json.Marshal(tyk)
	fmt.Println("tyk bytes: ", string(jsBytes))


	jsBytes1, _ := json.MarshalIndent(tyk, "", "  ")
	fmt.Println("tyk bytes: ",string(jsBytes1))

	jsBytes2 := []byte{}
	fmt.Println("tyk bytes2: ",string(jsBytes2))

	tykCopy:= reflect.New(reflect.TypeOf(tyk))
	fmt.Println(" tykCopy after new: %#v", tykCopy)

	tykCopy2 := tykCopy.Interface()
	fmt.Println(" tykCopy2 after Interface: %#v", tykCopy2)

	jsErr := json.Unmarshal(jsBytes, &tykCopy2)
	if jsErr != nil {
		fmt.Println("Couldn't unmarshal configuration: ", jsErr)
	} else {
		fmt.Println("unmarshaling tykCopy2: %#v", tykCopy2)

		fmt.Println("orig object tyk      :  %#v", tyk)
	}

}

type Employee struct {
	Name string `json:"name"`
	Age int `json:"age"`
	JobTitle string `json:"job_title"`
	salary float64 `json: salary`//this is private member
}


type Company struct {
	Name string
	Employees []Employee
}

type human interface {
	bodySize() float64
	eysColor() []color.RGBA
}