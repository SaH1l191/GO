package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func add(a int, b int) int {
	return a + b
}

func addProduct(a int, b int) (int, int) {
	return a + b, a * b
}

// named return values
func addSubtract(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

// variadic function
// takes in any no of int arguments
func addAll(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

func main() {
	
	// variables
	// var varName type = value
	var greeting string = "Hello, Go!"
	fmt.Println(greeting)
	fmt.Println(math.Ceil(3.28))
	fmt.Println(strings.ToUpper("adsasd"))

	//inferred type	
	var age = 30
	fmt.Println("Age:", age)
	

	// short declaration
	city1, city2 := "New York", "Los Angeles"
	fmt.Println("City1:", city1)
	fmt.Println("City2:", city2)

	//boolean
	isAdmin, isLoggedIn := true, false
	canDeletePost := (isAdmin && isLoggedIn) || isLoggedIn
	fmt.Println("Can delete post:", canDeletePost)

	age1 := 342
	isAdult := age1 >= 18
	fmt.Println("Is adult:", isAdult)

	//constants
	const ENV_VAR string = "production"

	//ifelse
	score := 85
	if score > 30 {
		fmt.Println("You passed!")
	} else if score <= 30 {
		fmt.Println("Just passed!")
	}

	//if with short statement
	if temp := 25; temp > 20 {
		fmt.Println("It's warm outside.")
	} else {
		fmt.Println("It's cold outside.")
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	//switch

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the work week.")
	case "Friday":
		fmt.Println("End of the work week.")
	default:
		fmt.Println("It's just another day.")
	}

	// arrays
	var m = [3]int{1, 2, 3}
	fmt.Println("Array m:", m)

	//literals
	n := []string{"Go", "Python", "Java"}
	n = append(n, "C++")
	fmt.Println("Slice n:", n)

	//len : , capacity : how many elements can be stored in the underlying array

	// make([]T , length, capacity)
	scores := make([]int, 3, 5)
	fmt.Println("Scores slice:", scores, len(scores), cap(scores))
	scores = append(scores, 90, 85, 88, 92)
	// on exceeding capacity, the capacity is doubled
	fmt.Println("Updated Scores slice:", scores, len(scores), cap(scores))

	// for loop
	views := []int{100, 200, 300, 400, 500}
	for index, value := range views {
		fmt.Printf("Day: %d, Value: %d\n", index, value)
	}

	//maps
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	//OR ages:=make(map[string]int)
	ages["Charlie"] = 35
	fmt.Println("Ages map:", ages)
	delete(ages, "Bob")
	fmt.Println("Updated Ages map:", ages)

	points := map[string]int{
		"A": 90,
		"B": 85,
		"C": 88,
	}
	valA, okA := points["A"]
	if okA {
		fmt.Println("Point A exists : ", valA)
	} else {
		fmt.Println("Point A does not exist")
	}
	for key, value := range points {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	//functions
	fmt.Println(add(2, 3))
	sum, product := addProduct(2, 3)
	fmt.Println("Sum:", sum, "Product:", product)
	//OR if we want to ignore one return value
	// sum,_:=addProduct(2,3)

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(addAll(nums...))

	//anonymous function ( immediate invocation of fnc)
	res := func(n int) int {
		return n * n
	}
	fmt.Println("Square of 5:", res(5))

	level, err := parseLevel("5")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Level:", level)
	}
	// shorter way :
	// if level, err := parseLevel("3"); err != nil {
	// 	fmt.Println("level:", level)
	// }else{
	// 	fmt.Println("Error:", err)
	// }

	

}

// error handling
// go doesnt raise exceptions for errors :
// instead it returns error as a value which can be checked
func parseLevel(s string) (int, error) {
	// return type : (value ,error)
	n, err := strconv.Atoi(s)
	if n == 5 {
		fmt.Println(err)
		// return 0, fmt.Errorf("level 5 is not allowed")
		return 0, err
	}
	return n, nil
}
   