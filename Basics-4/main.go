package main

import (
	"errors"
	"fmt"
	"strings"
)

//Structs group related fields together.
//Interfaces define behavior (method sets).
// interface satisfaction is checked at compile time.
type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func counter() func() int {
	count := 0
	return func() int {
		count++;
		return count ;
	}
}

func checkType(v interface{}) {
	switch val := v.(type) { 
	case int:
		fmt.Println("Integer:", val)

	case string:
		fmt.Println("String:", val)

	case bool:
		fmt.Println("Boolean:", val)

	default:
		fmt.Println("Unknown type")
	}
}


type Reader interface {
    Read(p []byte) (n int, err error)
}


func readConfig() error {
	return errors.New("file missing")
}

func startApp() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("startup failed: %w", err)
	}
	return nil
}


// init() Function
// init() runs automatically before main().
func init() {
    fmt.Println("Init function runs first")
}
 
func main() {
	fmt.Println("Main function runs second")
	c := counter(); 
	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) //3 


	arr := []int{1, 2, 3, 4, 5}
	fmt.Print(arr[1:2]) //excludes last index


	src := []int{1,2,3}
	dst := make([]int,3) 
	copy(dst, src) 
	fmt.Println(dst) // [1 2 3]]

	circle := Circle{radius: 5}
	fmt.Println("Area of circle:", circle.Area())


	// type assertion : to check if an interface value holds a specific type 
	checkType(10)
	checkType("Go")
	checkType(true)
	

	//io.Reader 
	r := strings.NewReader("Hello")
	buff := make([]byte,5)
	n ,_  := r.Read(buff)
	fmt.Println(string(buff[:n]))
	//similar stringer  ,reader ,writer , err ,

	// errors in GO 
	err := errors.New("file not found")
	fmt.Println("Error:", err)

	//fmt.Errorf() : Allows formatted errors.
	id := 23;
	if id == 23 {
		fmt.Errorf("User %d not found\n", id)
	} 

	//Error Wrapping (%w) : Go allows wrapping one error inside another. This preserves the original error chain.
	if err := startApp(); err != nil {
		fmt.Println("Error:", err)
	}

	//errors.Is() : Checks if an error is in the chain of wrapped errors.
	// var ErrNotFound = errors.New("not found")
	// func findUser(id int) error {
	// 	if id != 1 {
	// 		return fmt.Errorf("db lookup failed: %w", ErrNotFound)
	// 	}
	// 	return nil
	// }

	// func main() {
	// 	err := findUser(5)

	// 	if errors.Is(err, ErrNotFound) {
	// 		fmt.Println("User does not exist")
	// 	}
	// }

	// go.mod defines the module name and dependencies.
	// go.sum stores checksums of dependencies.
	//go get is used to add or update dependencies.
	//go mod tidy cleans up unused dependencies and adds missing ones.
	//A package can have multiple init functions.




	//context in GO : 
	// 	A context carries: 
	// Cancellation signal 
	// Deadline / timeout 
	// Request-scoped values 
	// Propagation across goroutines 
	// Think of it as a control signal for a request.


	//context.Background() is the root context. : It never cancels and has no deadline.
	// 	Starting the application
	// In main()
	// In tests
	// When no parent context exists


	//context.WithCancel() :Creates a child context that can be cancelled manually.
	// ctx, cancel := context.WithCancel(context.Background()) 
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	cancel()
	// }()

	//context.WithTimeout() : Creates a child context that automatically cancels after a specified duration.
	// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	//Always pass context as first argument





	



}




 