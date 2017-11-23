package main

import (
	"math"
	"fmt"
	"os"
	"errors"
)

func main() {
	answer, err := sqrt(-10)
	fmt.Println("sqrt 100 = ", answer, err)

	answer, err = sqrt(-10)
	if err != nil {
		fmt.Println("Oops ", err)
		os.Exit(1)
	}

}

func sqrt(x float64) (f float64, err error) {

	if (x < 0) {
		return 0, errors.New("sqrt cant be negative")
	}
	return math.Sqrt(x), nil
}