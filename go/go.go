package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}
type User struct {
	name string
}
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}
func validateUser(u *User) error {
	if u.name == "" {
		return &ValidationError{Field: "name", Message: "cannot be empty"} //can be returned as it implmeents error interface
	}
	return nil
}




type APIError struct {
	StatusCode int
	Message   string
	Err 	error
}
func (e *APIError) Error() string {
	return fmt.Sprintf("Api error %d,%s",e.StatusCode,e.Message,)
}
var ErrNotFound = errors.New("not found")
func handleRequest() error {
	return &APIError{StatusCode: 500, Message: "Internal Server Error", 
	Err: ErrNotFound}
}
func (e *APIError) Unwrap() error {
    return e.Err
}





func normalswap(a, b int) (int, int) {
	x := b
	y := a
	return x, y
}

// named return functions
func namedReturnSwap(a, b int) (x int, y int) {
	x = b
	y = a
	return // returns x and y implicitly
}

// vardiac func (these vardiac parameters need to be at last passed arguemtns)
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

//usage : res = sum(1,2,3,4) // res = 10  or  sum(nums...)

// func logf(level string, args ...interface{})

//usage : logf("INFO", "User %s logged in", username)

func apply(f func(...int) int, nums ...int) int {
	return f(nums...)
}


// Closure: function that captures variables from outer scope
func counter() func() int {
	count := 0  // captured by closure
	return func() int {
		count++  // modifies captured variable
		return count
	}
}



func main() {
	fmt.Println("Hello, World!")

	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday)

	//type conversion
	var i int = 42
	var ff float64 = float64(i)
	fmt.Println("Integer:", i, "Float:", ff)

	//type assertion
	var iface interface{} = "Hello, Go!"
	str, ok := iface.(string)
	if ok {
		fmt.Println("Extracted string:", str)
	} else {
		fmt.Println("Failed to extract string")
	}

	switch v := iface.(type) {
	case string:
		fmt.Println("It's a string:", v)
	case int:
		fmt.Println("It's an integer:", v)
	default:
		fmt.Println("Unknown type")
	}

	//type alias
	type Myint = int
	var x Myint = 10
	fmt.Println("Myint value:", x)

	//named type
	type Celsius float64
	type Fahrenheit float64
	var c Celsius = 100
	var f Fahrenheit = 212
	//still need type conversion to perform operations
	f = Fahrenheit(c)*9/5 + 32
	fmt.Printf("%g°C is %g°F\n", c, f)

	//fmt package
	fmt.Print("no newline")       // prints without newline
	fmt.Println("with newline")   // prints with newline
	fmt.Printf("value: %d\n", 42) // formatted print

	fmt.Printf("type:%T", f) //type
	fmt.Printf("printing any value %v", x)
	type point struct {
		x, y int
	}
	p := point{3, 4}
	fmt.Printf("point: %+v\n", p) // prints field names and values

	// Common format verbs
	// %v  - default format (any value)
	// %+v - struct with field names
	// %#v - Go syntax representation
	// %T  - type of value
	// %d  - decimal integer
	// %s  - string
	// %f  - float
	// %t  - boolean
	// %p  - pointer address
	// %q  - quoted string

	s := fmt.Sprintf("formatted number: %d", 42) // returns formatted string
	fmt.Println(s)

	err := fmt.Errorf("an error occurred: %s", "something went wrong")
	fmt.Println(err)

	//if else s
	// Short statement - declare variable scoped to if/else
	if num := 10; num > 5 {
		fmt.Println("greater")
	} else {
		fmt.Println("lesser")
	}
	// num is NOT accessible here

	// Common pattern: error handling
	// if err := doSomething(); err != nil {
	// 	log.Fatal(err)
	// }

	//all 4 for loops
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	//#2
	condition := false
	for condition {

	}

	//#3
	done := true
	for {
		if done {
			break
		}
		//process
	}

	//#4 range loop
	for i, v := range "Hello" {
		fmt.Printf("Index: %d, Value: %c\n", i, v)
	}

	//fucntions

	//first class funcs
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("Sum:", add(3, 4))

	//pass func as parameters
	//usage : apply(sum,1,2,3,4) // returns 10
	apply(sum, 1, 2, 3, 4)


	//anonymouse func 
	result := func(a, b int) int {
		return a * b
	}(1,2)
	fmt.Println("Result of anonymous function:", result)

 
	greet := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}
	greet("World")

 
	
	vc := counter()
	fmt.Println(vc())  // 1
	fmt.Println(vc())  // 2
	fmt.Println(vc())  // 3
 

	// 
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println(n)  // prints 0,1,2,3,4
		}(i)
	}



	//defer 
// 	The key idea:
// defer executes later, but its arguments are evaluated immediately.
// And for named returns:
// defer runs after the return value is assigned but before the function actually exits.


	// func example() {
	//     x := 10
	//     defer fmt.Println(x)  // prints 10, not 11
	//     x++
	// }
 
	// func double(x int) (result int) {
	// 	defer func() {
	// 		result *= 2  // modifies return value after function body
	// 	}()
	// 	return x  // returns x, then defer doubles it
	// }

 
	// // Multiple defers run in LIFO (stack) order
	// func example() {
	//     defer fmt.Println("first")
	//     defer fmt.Println("second")
	//     defer fmt.Println("third")
	//     // Output: third, second, first
	// }

	
	// func readFile(filename string) error {
	// 	f, err := os.Open(filename)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer f.Close()  // guaranteed to run when function exits

	// 	// use f...
	// 	return nil
	// }

	//panic 
	// panic in Go is like a runtime emergency stop.
	// It immediately stops normal execution of the current function and 
	// starts unwinding the call stack (running deferred functions on the way up). 
	// If nobody handles it with recover(), the program crashes.
	// Your example:

	// func mustParse(s string) int {
	// 	n, err := strconv.Atoi(s)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	return n
	// }


	//error as value 
	// Go uses explicit error returns, not exceptions
	// func divide(a, b float64) (float64, error) {
	// 	if b == 0 {
	// 		return 0, fmt.Errorf("division by zero")
	// 	}
	// 	return a / b, nil
	// }

	// // Always check errors!
	// result, err := divide(10, 0)
	// if err != nil {
	// 	log.Printf("error: %v", err)
	// 	return
	// }
	// fmt.Println(result)


	//creating new errors 
	// var ErrNotFound  = errors.New("not found")
	// var ErrUnauthorized = errors.New("unauthorized")
	 
	// if errors.Is(err, ErrNotFound) {
	// 	// true - works even though error is wrapped
	// 	fmt.Println("user not found")
	// }
	
	//also in returning nil, err :
	// 	You lose context:
	// sql: no rows in result set
	// use nil , fmt.Errorf("user %d not found", userId) to wrap surroding context
	//getUser: sql: no rows in result set (much better )



	//sentinal errors 
	 
	// var (
	// 	ErrNotFound     = errors.New("not found")
	// 	ErrAlreadyExists = errors.New("already exists")
	// 	ErrPermission   = errors.New("permission denied")
	// 	ErrTimeout      = errors.New("operation timed out")
	// )

	// // Usage in code
	// func GetByID(id string) (*Item, error) {
	// 	item, err := db.Find(id)
	// 	if err != nil {
	// 		if errors.Is(err, sql.ErrNoRows) {
	// 			return nil, fmt.Errorf("GetByID: %w", ErrNotFound)
	// 		}
	// 		return nil, fmt.Errorf("GetByID: %w", err)
	// 	}
	// 	return item, nil
	// }

	//errors.Is when you're comparing against a known error value.
	// error.as :  when you're looking for a particular error type and want to extract it.
	 
	// err = validateUser(&User{name: ""})
	// var ve *ValidationError
	// if errors.As(err, &ve) {
	// 	fmt.Printf("Validation failed: %s\n", ve)
	// } else if err != nil {
	// 	fmt.Printf("Other error: %v\n", err)
	// } else {
	// 	fmt.Println("User is valid")
	// }
	

	err = handleRequest()
	//It checks whether an error matches a specific sentinel error, even if it’s wrapped.
	//Even though ErrNotFound is wrapped inside another error, errors.Is() walks the chain:
	if errors.Is(err, ErrNotFound) { //needs unwrap func to work
		fmt.Println("User not found !!!!!!!!!!!!!!!")
	}
	var apiError *APIError
	if errors.As(err, &apiError) { // works without unwrap  , compares type and extracts if matches
		//`As` extracts typed error from chain
		fmt.Printf("API error: %d - %s\n", apiError.StatusCode, apiError.Message)
	}
	//What errors.As() does
	// It:
	// walks the error chain
	// finds a matching type (*MyError)
	// assigns it to myErr
	// So you can safely access fields.

	// A sentinel error in Go is a predefined, fixed error value that is used to represent a specific, well-known failure condition.
	// Its called “sentinel” because it acts like a marker you compare against.
 
  
}
