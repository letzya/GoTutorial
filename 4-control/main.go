package main

import (
	"fmt"
	"time"
)

func main() {
	//forLoop()
	//forever()
	//whileLop()
	doDefer()
}

func doDefer() {
	defer fmt.Println("Deferred")

	fmt.Println("something")
}

func forLoop() {

	for i := 0; i < 10; i++ {
		fmt.Print("For loop", i, "\n")
	}
}

func forever() {

	for {
		time.Sleep(time.Second * 1)
		fmt.Println(time.Now())
	}
}


func whileLop() {

	counter := 0
	finished := false

	for !finished {
		fmt.Println("Counter ", counter)

		if counter == 5 {
			finished = true
		}
		counter++
	}
}