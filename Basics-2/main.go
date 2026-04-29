package main

import (
	"errors"
	"fmt"
)

// defer statement :
// primarily used for cleanup tasks and resource management
func doWork(success bool) error {
	fmt.Println("acquiring resources...")
	if !success {
		return errors.New("Something went wrong ! Returning Early! ")
	}
	fmt.Println("Doing some work...")
	defer fmt.Println("cleanup : Releasing resources...")
	return nil
}

// pointers : when we want to mod
func addScore(score *int) {
	*score += 5
	// no need to return directly modifies the original variable
}

func main() {
	fmt.Println("Case:1")
	if err := doWork(true); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("\nCase:2")
	if err := doWork(false); err != nil {
		fmt.Println("Error:", err)
	}

	// pointers
	score := 10
	addScore(&score)
	fmt.Println(score)

	//structs
	type Person struct {
		name string
		age  int
	}
	p1 := Person{name: "Alice", age: 30}
	//mutable by default
	fmt.Println("Person 1:", p1)


	//partial struct 
	p2 := Person{name: "Bob"}
	fmt.Println("Person 2:", p2) // age will be 0 by default 
 
	// Value Receiver
	// func (p Person) greet() {
	// 	fmt.Println("Hello", p.Name)
	// }

	// Works on a copy.

	// Pointer Receiver
	// func (p *Person) birthday() {
	// 	p.Age++
	// }
	
}
