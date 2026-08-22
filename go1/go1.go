package main

import "fmt"

type Address struct {
	City    string
	Country string
}
type User struct {
	Name string
	Address
}

type Contact struct {
	Email string
	Phone string
}

type Employee struct {
	User
	Contact
	EmployeeID int
}

func main() {
	// var arr []int
	arr := []int{1, 2, 3, 4, 5}
	// arr := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	s := make([]int, 4)
	fmt.Print(s)

	// s1:= []int{1,2}
	// s2:= []int{3,4}
	// s2 = append(s1,s2...)

	//slices
	//s2:= s2[1:3] // last is exclusive

	
	


}
