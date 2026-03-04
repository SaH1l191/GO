package main

import (
	"fmt"
)

// method  value receiver

type Person struct {
	name string
	age  int
}

// greet() is a method with a value receiver
// The method reads data but doesn’t modify it
//p is a copy of the Person value that calls the method.read-only behavior
func (p Person) greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.name, p.age)
}


// method pointer receiver
type User struct {
	name string
	age  int
}

func (u *User) Birthday() {
	u.age += 1
}

func main() {
	p1 := Person{name: "Alice", age: 30}
	fmt.Println(p1.greet())

	u1 := User{name: "john_doe", age: 25}
	u1.Birthday()
	fmt.Println(u1.age)
}
