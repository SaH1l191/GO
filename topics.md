# Master Go Topics - Fresher Interview Guide 2026

> Based on your codebase analysis + industry standards for 2026

---

## 1. Go Fundamentals (Must Know - Day 1)

### 1.1 Variables & Data Types

- [x] `var` declaration, short `:=`, type inference

- [x] Constants, `iota` (enumerated constants)

- [x] Zero values for all types

- [x] Type conversion vs type assertion

- [x] Named types, type aliases vs type definitions

#### Variable Declaration Styles

```go
// Style 1: var with explicit type
var name string = "Go"

// Style 2: var with type inference (initialization required)
var name = "Go"

// Style 3: Short declaration (only inside functions)
name := "Go"

// Style 4: Multiple declarations
var (
    name   string = "Go"
    age    int    = 10
    active bool   = true
)

// Style 5: Multiple short declarations
name, age, active := "Go", 10, true
```

#### Constants and iota

```go
// Basic constants
const Pi = 3.14159
const (
    StatusOK    = 200
    StatusNotFound = 404
)

// iota - auto-incrementing constants (starts at 0)
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
)

// iota with expressions
type ByteSize int64
const (
    _           = iota             // ignore
    KB ByteSize = 1 << (10 * iota) // 1 << 10 = 1024
    MB                                // 1 << 20
    GB                                // 1 << 30
)

// iota with custom starting value
const (
    First = iota + 1  // 1
    Second             // 2
    Third              // 3
)
```

#### Zero Values (CRITICAL for interviews)

Every type in Go has a zero value - the value assigned when no explicit value is provided:

```go
var i int         // 0
var f float64     // 0.0
var b bool        // false
var s string      // ""
var p *int        // nil
var sl []int      // nil
var m map[string]int // nil
var ch chan int    // nil
var fn func()     // nil
var iface interface{} // nil

// Struct zero value has all fields at their zero values
type User struct {
    Name string
    Age  int
}
var u User  // User{Name: "", Age: 0}
```

**Interview Q:** What is the zero value of `int`? `string`? `bool`? `pointer`? `slice`? `map`? **Answer:** 0, "", false, nil, nil, nil

#### Type Conversion vs Type Assertion

```go
// Type Conversion (explicit, works between compatible types)
var i int = 42
var f float64 = float64(i)    // int -> float64
var u uint = uint(f)          // float64 -> uint
var s string = string(65)     // int -> string (Unicode code point) = "A"

// Type Assertion (for interfaces)
var iface interface{} = "hello"
str, ok := iface.(string)     // type assertion with comma-ok
if ok {
    fmt.Println(str)           // "hello"
}

// Type Switch (better for multiple types)
switch v := iface.(type) {
case string:
    fmt.Println("string:", v)
case int:
    fmt.Println("int:", v)
default:
    fmt.Println("unknown type")
}
```

#### Named Types vs Type Aliases vs Type Definitions

```go
// Type Definition - creates a NEW type
type Celsius float64
type Fahrenheit float64

// Cannot directly mix types!
var c Celsius = 100
var f Fahrenheit = float64(c)  // ERROR: need explicit conversion
f = Fahrenheit(c * 9/5 + 32)  // OK: explicit conversion

// Type Alias - just another name for existing type
type MyInt = int  // alias, not new type

var x int = 10
var y MyInt = x   // OK: same type

// When to use each:
// Type Definition: when you want type safety (Celsius vs Fahrenheit)
// Type Alias: when you want compatibility (deprecated type migration)

// Example: Custom type with methods
type Distance float64

func (d Distance) InMeters() float64 {
    return float64(d) * 0.3048
}

// Example: Type alias for readability
type (
    UserID = string
    OrderID = string
)
```

#### fmt Package Essentials

```go
import "fmt"

// Print variations
fmt.Print("no newline")           // prints without newline
fmt.Println("with newline")       // prints with newline
fmt.Printf("value: %d\n", 42)    // formatted print

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

// Sprintf returns string instead of printing
s := fmt.Sprintf("age: %d, name: %s", 25, "Alice")

// Errorf creates formatted error
err := fmt.Errorf("invalid age: %d", age)
```

---

### 1.2 Control Flow

- [x] `if/else` with short statement

- [x] `for` loop (only loop in Go - no while/do)

- [x] `switch` statement (no fallthrough by default)

- [x] `range` (for arrays, slices, maps, strings, channels)

#### if/else with Short Statement

```go
// Short statement - declare variable scoped to if/else
if num := 10; num > 5 {
    fmt.Println("greater")
} else {
    fmt.Println("lesser")
}
// num is NOT accessible here

// Common pattern: error handling
if err := doSomething(); err != nil {
    log.Fatal(err)
}
```

#### The 4 Forms of for Loop

```go
// Form 1: Traditional (C-style)
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Form 2: While-style (condition only)
for condition {
    // do something
}

// Form 3: Infinite loop
for {
    // break to exit
    if done {
        break
    }
}

// Form 4: Range-based
for i, v := range slice {
    fmt.Printf("index: %d, value: %v\n", i, v)
}
```

#### switch Statement

```go
// Basic switch (no fallthrough by default!)
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("Almost weekend")
case "Saturday", "Sunday":  // multiple values
    fmt.Println("Weekend!")
default:
    fmt.Println("Midweek")
}

// Expression-less switch (replaces if-else chains)
switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
case score >= 70:
    grade = "C"
default:
    grade = "F"
}

// Type switch
switch v := value.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
}

// fallthrough keyword (rare, breaks Go conventions)
switch x {
case 1:
    fmt.Println("one")
    fallthrough  // forces next case to execute
case 2:
    fmt.Println("one or two")  // always executes if case 1 matches
}
```

#### range Patterns

```go
// Slice/Array
for i, v := range []int{10, 20, 30} {
    fmt.Printf("index %d: %d\n", i, v)
}

// Map
for k, v := range map[string]int{"a": 1, "b": 2} {
    fmt.Printf("%s = %d\n", k, v)
}

// String (range gives rune, not byte!)
for i, ch := range "Hello, 世界" {
    fmt.Printf("byte %d: %c (rune %d)\n", i, ch, ch)
}

// Channel
for v := range ch {
    fmt.Println(v)
}

// Blank identifier (ignore index or value)
for _, v := range slice {  // ignore index
    fmt.Println(v)
}
for i := range slice {  // ignore value (just want index)
    fmt.Println(i)
}

// Only need value, not index (Go 1.22+)
for v := range 10 {  // 0 to 9
    fmt.Println(v)
}
```

**Interview Q:** What are the 4 forms of `for` loop in Go? **Answer:**

```go
for i := 0; i < 5; i++ {}    // traditional
for condition {}               // while-style
for {}                         // infinite loop
for i, v := range slice {}    // range
```

---

### 1.3 Functions

- [x] Multiple return values

- [x] Named return values

- [x] Variadic functions (`...int`)

- [x] First-class functions (assign to variables)

- [x] Anonymous functions / IIFE

- [x] Closures

- [x] `defer` in depth (LIFO order, panic/recover)

#### Multiple Return Values

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Usage
result, err := divide(10, 3)
if err != nil {
    log.Fatal(err)
}
```

#### Named Return Values

```go
func swap(a, b int) (x, y int) {
    x = b
    y = a
    return  // naked return - returns x and y
}

// Named returns are just regular variables
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // returns x=20, y=25 for sum=45
}
```

#### Variadic Functions

```go
// Accept any number of int arguments
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// Usage
sum(1, 2, 3)       // 6
sum(1, 2, 3, 4, 5) // 15

// Slice unpacking with ...
nums := []int{1, 2, 3}
sum(nums...)  // unpack slice

// Variadic with other parameters (must be last!)
func logf(level string, args ...interface{}) {
    fmt.Printf("[%s] %v\n", level, args)
}
logf("INFO", "user", 123, "logged in")
```

#### First-Class Functions

```go
// Functions are values - assign to variables
add := func(a, b int) int {
    return a + b
}
result := add(5, 3)  // 8

// Pass functions as arguments
func apply(f func(int, int) int, a, b int) int {
    return f(a, b)
}
result := apply(add, 5, 3)

// Return functions from functions
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}
double := multiplier(2)
triple := multiplier(3)
fmt.Println(double(5))  // 10
fmt.Println(triple(5))  // 15
```

#### Anonymous Functions / IIFE

```go
// Immediately Invoked Function Expression (IIFE)
result := func(a, b int) int {
    return a + b
}(5, 3)  // 8

// Anonymous function assigned to variable
greet := func(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
greet("World")

// Goroutine with anonymous function
go func() {
    fmt.Println("running in goroutine")
}()
```

#### Closures

```go
// Closure: function that captures variables from outer scope
func counter() func() int {
    count := 0  // captured by closure
    return func() int {
        count++  // modifies captured variable
        return count
    }
}

c := counter()
fmt.Println(c())  // 1
fmt.Println(c())  // 2
fmt.Println(c())  // 3

// Common gotcha: loop variable capture
// BAD (pre-Go 1.22):
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // all might print 5!
    }()
}

// GOOD (Go 1.22+): loop variable is per-iteration
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // prints 0,1,2,3,4 (any order)
    }()
}

// GOOD (pre-Go 1.22): pass as parameter
for i := 0; i < 5; i++ {
    go func(n int) {
        fmt.Println(n)  // prints 0,1,2,3,4
    }(i)
}
```

#### defer In Depth

```go
// defer executes AFTER surrounding function returns
// Used for cleanup: closing files, unlocking mutexes, etc.

func readFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()  // guaranteed to run when function exits

    // use f...
    return nil
}

// Multiple defers run in LIFO (stack) order
func example() {
    defer fmt.Println("first")
    defer fmt.Println("second")
    defer fmt.Println("third")
    // Output: third, second, first
}

// defer with arguments evaluated immediately
func example() {
    x := 10
    defer fmt.Println(x)  // prints 10, not 11
    x++
}

// defer with named return values (can modify return value)
func double(x int) (result int) {
    defer func() {
        result *= 2  // modifies return value after function body
    }()
    return x  // returns x, then defer doubles it
}
// double(5) returns 10

// panic and recover
func safeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r)
        }
    }()
    return a / b, nil
}

// Panic example
func mustParse(s string) int {
    n, err := strconv.Atoi(s)
    if err != nil {
        panic(err)  // crashes unless recovered
    }
    return n
}

// Recover only works in deferred functions
func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered:", r)
        }
    }()
    panic("something bad")  // caught by recover
}
```

**Interview Q:** What is the order of execution for multiple `defer` statements? **Answer:** LIFO (stack) order - last deferred runs first.

**Interview Q:** What happens to named return values when you `defer`? **Answer:** `defer` can modify named return values. `return` sets them, `defer` can read/write.

**Interview Q:** When should you use `recover()`? **Answer:** Only in deferred functions, only when you can handle the panic gracefully (e.g., HTTP handlers, goroutines). Don't use as flow control.

---

### 1.4 Error Handling

- [x] `error` as a value (not exception)

- [x] `errors.New()`, `fmt.Errorf()`

- [x] Error wrapping with `%w`

- [x] `errors.Is()` for error chain checking

- [x] `errors.As()` for typed errors

- [x] Sentinel errors (`var ErrNotFound = errors.New(...)`)

- [x] Custom error types (implementing `Error()` method)

#### Error as Value (Go's philosophy)

```go
// Go uses explicit error returns, not exceptions
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Always check errors!
result, err := divide(10, 0)
if err != nil {
    log.Printf("error: %v", err)
    return
}
fmt.Println(result)
```

#### Creating Errors

```go
import (
    "errors"
    "fmt"
)

// errors.New for simple errors
var ErrNotFound = errors.New("not found")
var ErrUnauthorized = errors.New("unauthorized")

// fmt.Errorf for formatted errors
err := fmt.Errorf("user %d not found", userID)

// Error wrapping with %w (creates error chain)
func getUser(id int) (*User, error) {
    err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user)
    if err != nil {
        return nil, fmt.Errorf("getUser: %w", err)  // wraps underlying error
    }
    return &user, nil
}
```

#### errors.Is for Error Chain Checking

```go
// errors.Is checks if error chain contains target error
var ErrNotFound = errors.New("not found")

func findUser(id int) error {
    return fmt.Errorf("user service: %w", ErrNotFound)
}

err := findUser(1)
if errors.Is(err, ErrNotFound) {
    // true - works even though error is wrapped
    fmt.Println("user not found")
}

// Also works with == for simple comparison
if err == ErrNotFound {
    // only works if err is EXACTLY ErrNotFound (not wrapped)
}
```

#### errors.As for Typed Errors

```go
// errors.As extracts typed error from chain
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validateUser(u *User) error {
    if u.Name == "" {
        return &ValidationError{
            Field:   "name",
            Message: "is required",
        }
    }
    return nil
}

// Usage
err := validateUser(&User{})
if err != nil {
    var valErr *ValidationError
    if errors.As(err, &valErr) {
        fmt.Printf("Validation failed on field '%s': %s\n",
            valErr.Field, valErr.Message)
    }
}
```

#### Sentinel Errors

```go
// Package-level error variables (sentinels)
var (
    ErrNotFound     = errors.New("not found")
    ErrAlreadyExists = errors.New("already exists")
    ErrPermission   = errors.New("permission denied")
    ErrTimeout      = errors.New("operation timed out")
)

// Usage in code
func GetByID(id string) (*Item, error) {
    item, err := db.Find(id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, fmt.Errorf("GetByID: %w", ErrNotFound)
        }
        return nil, fmt.Errorf("GetByID: %w", err)
    }
    return item, nil
}
```

#### Custom Error Types

```go
// Custom error with structured data
type APIError struct {
    StatusCode int
    Message    string
    Err        error  // wraps underlying error
}

func (e *APIError) Error() string {
    return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Message)
}

func (e *APIError) Unwrap() error {
    return e.Err  // enables errors.Is and errors.As
}

// Usage
func handleRequest() error {
    return &APIError{
        StatusCode: 404,
        Message:    "resource not found",
        Err:        ErrNotFound,
    }
}

// Check
err := handleRequest()
if errors.Is(err, ErrNotFound) {
    // true - Unwrap() chain works
}

var apiErr *APIError
if errors.As(err, &apiErr) {
    fmt.Println(apiErr.StatusCode)  // 404
}
```

#### Common Error Patterns

```go
// Pattern 1: Error wrapping with context
func processOrder(id string) error {
    order, err := getOrder(id)
    if err != nil {
        return fmt.Errorf("processOrder: failed to get order %s: %w", id, err)
    }
    // ...
}

// Pattern 2: Custom error checking
func IsNotFound(err error) bool {
    return errors.Is(err, ErrNotFound) ||
           (errors.Is(err, sql.ErrNoRows))
}

// Pattern 3: Error aggregation
type MultiError struct {
    Errors []error
}

func (e *MultiError) Error() string {
    msgs := make([]string, len(e.Errors))
    for i, err := range e.Errors {
        msgs[i] = err.Error()
    }
    return strings.Join(msgs, "; ")
}

// Pattern 4: Ignore specific errors
rows, err := db.Query(query)
if err != nil {
    return fmt.Errorf("query: %w", err)
}
defer rows.Close()
// rows.Close() error is typically ignored
```

**Interview Q:** What is the difference between `errors.Is` and `errors.As`? **Answer:** `Is` checks value equality in chain. `As` extracts typed error from chain.

**Interview Q:** Should you use `==` or `errors.Is()` to compare errors? **Answer:** Always use `errors.Is()` - it works with wrapped errors. `==` only works for unwrapped errors.

---

## 2. Data Structures

### 2.1 Arrays & Slices

- [x] Arrays (fixed size, value type)

- [x] Slices (dynamic, reference type)

- [x] `make([]int, len, cap)` - length vs capacity

- [x] `append()` - capacity doubling behavior

- [x] `copy()` function

- [x] Slice expression `slice[low:high]`

- [x] Slice internals (pointer, length, capacity)

- [x] Nil slice vs empty slice

- [x] Slice of slices (2D)

#### Arrays (Fixed Size, Value Type)

```go
// Array declaration - size is part of the type!
var arr [5]int           // [0, 0, 0, 0, 0]
arr := [5]int{1, 2, 3, 4, 5}
arr := [...]int{1, 2, 3}  // compiler counts: [3]int

// Arrays are VALUE types (copied!)
original := [3]int{1, 2, 3}
copy := original
copy[0] = 99
fmt.Println(original)  // [1, 2, 3] - unchanged!
fmt.Println(copy)      // [99, 2, 3]

// Fixed size - cannot append
arr := [3]int{1, 2, 3}
// arr = append(arr, 4)  // ERROR!

// Comparison
a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
fmt.Println(a == b)  // true - arrays are comparable
```

#### Slices (Dynamic, Reference Type)

```go
// Slice declaration
var s []int              // nil slice
s := []int{1, 2, 3}     // literal
s := make([]int, 5)     // len=5, cap=5
s := make([]int, 3, 5)  // len=3, cap=5

// Slices are REFERENCE types (shared backing array!)
original := []int{1, 2, 3, 4, 5}
slice := original[1:3]  // [2, 3] - shares memory!
slice[0] = 99
fmt.Println(original)  // [1, 99, 3, 4, 5] - CHANGED!

// Length vs Capacity
s := make([]int, 3, 5)
fmt.Println(len(s))  // 3 - elements stored
fmt.Println(cap(s))  // 5 - space in backing array
```

#### Slice Internals (Runtime SliceHeader)

```go
// A slice is a 3-word struct (runtime.sliceHeader):
// 1. Pointer - to underlying array
// 2. Length  - number of elements
// 3. Capacity - total space in backing array

// Visual representation:
// s := make([]int, 3, 5)
//
// sliceHeader: {ptr: 0xc0000b2000, len: 3, cap: 5}
//
// backing array: [0, 0, 0, ?, ?]
//                 ^-- len=3 --^  ^-- cap=5 --^

// This is why slices share memory when sliced!
original := [5]int{10, 20, 30, 40, 50}
s1 := original[1:3]  // points to same array
s2 := original[2:4]  // points to same array
// s1 = [20, 30], s2 = [30, 40]
// changing s2[0] affects s1[1] and original[2]
```

#### append() - Capacity Doubling

```go
// append adds elements and grows slice if needed
s := make([]int, 0, 2)  // len=0, cap=2
fmt.Println(cap(s))     // 2

s = append(s, 1)
s = append(s, 2)
fmt.Println(cap(s))     // 2 (still fits)

s = append(s, 3)  // triggers reallocation!
fmt.Println(cap(s))     // 4 (doubled)

// Append returns new slice header (may point to new array)
s1 := []int{1, 2, 3}
s2 := append(s1, 4)  // s2 may point to new array
s1[0] = 99            // may or may not affect s2!

// Append multiple elements
s = append(s, 4, 5, 6)

// Append slice to slice
s1 := []int{1, 2}
s2 := []int{3, 4}
s3 := append(s1, s2...)  // [1, 2, 3, 4]

// Pre-allocate for performance
s := make([]int, 0, 1000)  // avoid repeated reallocations
for i := 0; i < 1000; i++ {
    s = append(s, i)
}
```

#### copy() Function

```go
// copy copies elements from src to dst
dst := make([]int, 3)
src := []int{1, 2, 3, 4, 5}
n := copy(dst, src)  // n = 3 (number copied)
// dst = [1, 2, 3]

// Copy only copies up to len(dst)
dst := make([]int, 2)
src := []int{1, 2, 3}
copy(dst, src)  // dst = [1, 2] (only 2 elements)

// Copy to avoid shared memory
original := []int{1, 2, 3}
copySlice := make([]int, len(original))
copy(copySlice, original)
copySlice[0] = 99  // doesn't affect original

// Copy slice of bytes
src := []byte("hello")
dst := make([]byte, len(src))
copy(dst, src)
```

#### Slice Expressions

```go
arr := [5]int{10, 20, 30, 40, 50}

// Basic slice: arr[low:high]
arr[1:3]   // [20, 30] - from index 1 to 2 (high exclusive)
arr[0:]    // [10, 20, 30, 40, 50] - from start
arr[:3]    // [10, 20, 30] - up to index 2
arr[:]     // [10, 20, 30, 40, 50] - full slice

// Slice with capacity: arr[low:high:max]
arr[1:3:4]  // [20, 30], cap = 3 (max-low = 4-1)
// This limits capacity, preventing accidental shared memory

// Reslicing
s := []int{1, 2, 3, 4, 5}
s = s[1:3]   // [2, 3]
s = s[:1]    // [2] - can reslice within capacity
// s = s[:10] // PANIC: out of range
```

#### Nil Slice vs Empty Slice

```go
// Nil slice - no backing array
var s []int
fmt.Println(s == nil)  // true
fmt.Println(len(s))    // 0
fmt.Println(cap(s))    // 0

// Empty slice - has backing array but no elements
s := []int{}
fmt.Println(s == nil)  // false
fmt.Println(len(s))    // 0
fmt.Println(cap(s))    // 0

// Make empty slice
s := make([]int, 0)
s := *new([]int)

// JSON marshaling difference
var nilSlice []int
emptySlice := []int{}

json.Marshal(nilSlice)   // null
json.Marshal(emptySlice) // []

// Best practice: return empty slice, not nil
func getItems() []int {
    items, err := db.Query(...)
    if err != nil {
        return []int{}  // return empty, not nil
    }
    return items
}
```

#### Slice of Slices (2D)

```go
// 2D slice
matrix := make([][]int, 3)  // 3 rows
for i := range matrix {
    matrix[i] = make([]int, 4)  // 4 columns
}

// Literal 2D slice
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Access
matrix[1][2]  // 6

// Iterate
for i, row := range matrix {
    for j, val := range row {
        fmt.Printf("matrix[%d][%d] = %d\n", i, j, val)
    }
}
```

#### Common Slice Operations

```go
// Remove element at index i (order changes)
s := []int{1, 2, 3, 4, 5}
i := 2  // remove 3
s = append(s[:i], s[i+1:]...)  // [1, 2, 4, 5]

// Remove element (preserve order, O(n))
s = append(s[:i], s[i+1:]...)

// Insert at index i
s = append(s[:i], append([]int{newElement}, s[i:]...)...)

// Contains
func contains(s []int, v int) bool {
    for _, item := range s {
        if item == v {
            return true
        }
    }
    return false
}

// Index of
func indexOf(s []int, v int) int {
    for i, item := range s {
        if item == v {
            return i
        }
    }
    return -1
}

// Filter
func filter(s []int, f func(int) bool) []int {
    result := []int{}
    for _, v := range s {
        if f(v) {
            result = append(result, v)
        }
    }
    return result
}

// Map
func mapInts(s []int, f func(int) int) []int {
    result := make([]int, len(s))
    for i, v := range s {
        result[i] = f(v)
    }
    return result
}
```

**Interview Q:** What is the difference between `len()` and `cap()` of a slice? **Answer:** `len` = number of elements currently stored. `cap` = total space in underlying array before reallocation.

**Interview Q:** What happens when you `append` beyond capacity? **Answer:** New underlying array allocated, typically double the capacity. Original slice not modified.

**Interview Q:** What is the difference between nil slice and empty slice? **Answer:** Both have len=0, cap=0. Nil slice has no backing array (pointer=nil), empty slice has backing array. JSON: nil-&gt;null, empty-&gt;\[\].

---

### 2.2 Maps

- [x] Creation (literal, `make`)

- [x] `delete()` function

- [x] Comma-ok pattern `val, ok := m[key]`

- [x] Iterate with `range`

- [x] Map is NOT goroutine-safe (race condition risk)

- [x] Map internals (hash table, buckets)

- [x] Ordered iteration myth (maps are unordered)

#### Map Creation

```go
// Literal initialization
m := map[string]int{
    "alice": 30,
    "bob":   25,
}

// make with initial capacity
m := make(map[string]int, 100)

// Empty map
m := make(map[string]int)
m := map[string]int{}  // literal

// Nil map (cannot write to!)
var m map[string]int
// m["key"] = 1  // PANIC: assignment to entry in nil map
```

#### Basic Operations

```go
m := make(map[string]int)

// Insert/Update
m["alice"] = 30
m["bob"] = 25

// Read
age := m["alice"]  // 30

// Delete
delete(m, "alice")

// Check existence (comma-ok)
age, ok := m["alice"]
if ok {
    fmt.Println("alice's age:", age)
} else {
    fmt.Println("alice not found")
}

// Length
fmt.Println(len(m))  // number of key-value pairs
```

#### Map Iteration with range

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}

// Iterate all
for k, v := range m {
    fmt.Printf("%s = %d\n", k, v)
}

// Keys only
for k := range m {
    fmt.Println(k)
}

// Values only
for _, v := range m {
    fmt.Println(v)
}

// Sorted keys (manual - map iteration order is random)
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
    fmt.Printf("%s = %d\n", k, m[k])
}
```

#### Map is NOT Goroutine-Safe

```go
// Race condition example (UNSAFE!)
m := make(map[string]int)
go func() {
    m["alice"] = 30  // concurrent write
}()
go func() {
    m["bob"] = 25  // concurrent write
}()
// FATAL: concurrent map writes

// Solution 1: Use sync.Mutex
var mu sync.Mutex
var m = make(map[string]int)

func safeWrite(key string, value int) {
    mu.Lock()
    defer mu.Unlock()
    m[key] = value
}

func safeRead(key string) (int, bool) {
    mu.RLock()
    defer mu.RUnlock()
    v, ok := m[key]
    return v, ok
}

// Solution 2: Use sync.Map (better for read-heavy)
var m sync.Map

m.Store("alice", 30)
v, ok := m.Load("alice")

// Solution 3: Channel-based (share memory by communicating)
type operation struct {
    key   string
    value int
    op    string  // "read", "write", "delete"
    result chan int
}
```

#### Map Internals (Hash Table)

```go
// Map is implemented as hash table:
// - Hash function on key
// - Buckets (8 entries each)
// - Overflow buckets when needed
//
// Performance:
// - Average: O(1) for read/write/delete
// - Worst case: O(n) if many hash collisions
//
// Key requirements:
// - Keys must be comparable (==, !=)
// - Slices, maps, functions cannot be keys
// - Struct keys OK if all fields comparable

// Good keys
map[string]int{}     // string
map[int]bool{}       // int
map[User]int{}       // struct with comparable fields

// Bad keys
map[[]int]int{}      // ERROR: slice not comparable
map[map[string]int]int{}  // ERROR: map not comparable
```

#### Ordered Iteration Myth

```go
// Maps are UNORDERED - iteration order is randomized
m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}

// Each iteration may give different order
for k, v := range m {
    fmt.Println(k, v)
}
// Run multiple times - different order each time!

// If you need order, use sorted keys
func sortedKeys(m map[string]int) []string {
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    return keys
}
```

#### Common Map Patterns

```go
// Set (map[string]struct{})
set := make(map[string]struct{})
set["alice"] = struct{}{}  // add
delete(set, "alice")       // remove
if _, ok := set["alice"]; ok {
    fmt.Println("exists")
}

// Default value (if key missing)
func getDefault(m map[string]int, key, defaultVal int) int {
    if v, ok := m[key]; ok {
        return v
    }
    return defaultVal
}

// Map with default on access (auto-vivification)
func increment(m map[string]int, key string) {
    m[key]++  // automatically creates key with value 0
}

// Invert map (swap keys and values)
func invert(m map[string]int) map[int]string {
    inverted := make(map[int]string, len(m))
    for k, v := range m {
        inverted[v] = k
    }
    return inverted
}

// Can you take address of map element?
m := map[string]int{"a": 1}
// &m["a"]  // ERROR: cannot take address
// Must copy to variable first
v := m["a"]
p := &v
```

**Interview Q:** Can you take the address of a map element? **Answer:** No. `&m["key"]` is invalid. Maps may relocate elements when growing.

---

### 2.3 Structs

- [x] Struct creation, partial initialization

- [x] Value receiver vs pointer receiver

- [x] Struct tags (`json:"name"`, `bson:"_id"`)

- [x] Embedded structs (composition over inheritance)

- [x] Anonymous structs

- [x] Struct comparison (value type - comparable)

#### Struct Creation

```go
// Named struct
type User struct {
    Name string
    Age  int
    Email string
}

// Initialization styles
u1 := User{"Alice", 30, "alice@example.com"}  // positional (order matters!)
u2 := User{Name: "Bob", Age: 25}              // named (partial init)
u3 := new(User)                                // pointer, zero values
u4 := &User{"Charlie", 35, ""}                // pointer literal
u5 := &User{Name: "Dave", Age: 28}           // pointer named

// Zero value struct
var u User  // User{Name: "", Age: 0, Email: ""}

// Struct with pointer field
type Node struct {
    Value int
    Next  *Node  // pointer to next node
}

// Struct with slice/map fields
type User struct {
    Name   string
    Hobbies []string         // nil until initialized
    Scores  map[string]int   // nil until initialized
}
```

#### Value Receiver vs Pointer Receiver

```go
type User struct {
    Name string
    Age  int
}

// Value receiver - works on copy, no mutation
func (u User) GetAge() int {
    return u.Age
}

// Pointer receiver - works on original, can mutate
func (u *User) SetAge(age int) {
    u.Age = age
}

// When to use which:
// Value receiver: small structs, no mutation needed, immutability
// Pointer receiver: large structs (avoid copy), need mutation, interface satisfaction

// Example: Both methods work on pointer
u := User{Name: "Alice", Age: 30}
u.GetAge()   // works: (&u).GetAge() auto-dereferenced
u.SetAge(31) // works

// Example: Value receiver can't mutate original
func (u User) BumpAge() {
    u.Age++  // only modifies copy!
}

u := User{Age: 30}
u.BumpAge()
fmt.Println(u.Age)  // still 30!

// Pointer receiver CAN mutate
func (u *User) BumpAge() {
    u.Age++  // modifies original!
}

u := User{Age: 30}
u.BumpAge()
fmt.Println(u.Age)  // 31
```

#### Struct Tags

```go
// Struct tags are metadata strings used by reflection
type User struct {
    ID    int    `json:"id" db:"user_id"`
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    Pass  string `json:"-"`  // omit from JSON
    Age   int    `json:"age,omitempty"`  // omit if zero
}

// Common tags
// json:"name"        - JSON field name
// json:"-"           - skip field
// json:",omitempty"  - omit if zero value
// bson:"_id"         - MongoDB field name
// db:"column_name"   - SQL column name
// validate:"required" - validation rules

// Multiple tags
type Product struct {
    ID    string  `json:"id" bson:"_id" db:"product_id" validate:"required"`
    Price float64 `json:"price" db:"price" validate:"required,gt=0"`
}

// Read tags at runtime
import "reflect"

t := reflect.TypeOf(User{})
field, _ := t.FieldByName("ID")
tag := field.Tag.Get("json")  // "id"
```

#### Embedded Structs (Composition)

```go
// Embedding - promotes inner struct's methods
type Address struct {
    City    string
    Country string
}

func (a Address) FullAddress() string {
    return fmt.Sprintf("%s, %s", a.City, a.Country)
}

type User struct {
    Name    string
    Address  // embedded (no field name!)
}

u := User{
    Name: "Alice",
    Address: Address{City: "NYC", Country: "USA"},
}

// Access embedded fields directly
fmt.Println(u.City)       // "NYC" (promoted)
fmt.Println(u.FullAddress())  // "NYC, USA" (promoted method)

// Or access via embedded type
fmt.Println(u.Address.City)  // "NYC"

// Multiple embedding
type Contact struct {
    Email string
    Phone string
}

type Employee struct {
    User
    Contact
    EmployeeID int
}

e := Employee{
    User:    User{Name: "Bob", Address: Address{City: "LA"}},
    Contact: Contact{Email: "bob@example.com"},
}

fmt.Println(e.Name)      // "Bob" (from User)
fmt.Println(e.Email)     // "bob@example.com" (from Contact)
// Conflict: if both User and Contact have same field, must specify:
// fmt.Println(e.User.Name)  // explicit
```

#### Anonymous Structs

```go
// One-time struct without naming
point := struct {
    X, Y int
}{10, 20}

// In functions
func process() {
    data := struct {
        Name string
        Age  int
    }{"Alice", 30}
    fmt.Println(data)
}

// In HTTP handlers
func handler(w http.ResponseWriter, r *http.Request) {
    response := struct {
        Status  string `json:"status"`
        Message string `json:"message"`
    }{
        Status:  "ok",
        Message: "success",
    }
    json.NewEncoder(w).Encode(response)
}

// Slice of anonymous structs
users := []struct {
    Name string
    Age  int
}{
    {"Alice", 30},
    {"Bob", 25},
}
```

#### Struct Comparison

```go
// Structs are comparable if all fields are comparable
type Point struct {
    X, Y int
}

p1 := Point{1, 2}
p2 := Point{1, 2}
p3 := Point{3, 4}

fmt.Println(p1 == p2)  // true
fmt.Println(p1 == p3)  // false
 NOT comparable
type User struct {
    Name    string
    Hobbies []string  // slice makes struct incomparable
}

u1 := User{"Alice", []string{"reading"}}
u2 := User{"Alice", []string{"reading"}}
// Structs with slice fields are
// fmt.Println(u1 == u2)  // ERROR: operator == not defined

// Use reflect.DeepEqual for complex comparison
if reflect.DeepEqual(u1, u2) {
    fmt.Println("equal")
}

// Or implement custom Equals method
func (u User) Equals(other User) bool {
    return u.Name == other.Name &&
           slices.Equal(u.Hobbies, other.Hobbies)
}
```

#### Factory Pattern with Structs

```go
// Constructor pattern (recommended)
type User struct {
    name string
    age  int
}

// Unexported fields + exported constructor
func NewUser(name string, age int) *User {
    return &User{name: name, age: age}
}

// Options pattern (for many optional parameters)
type Server struct {
    host    string
    port    int
    timeout time.Duration
    maxConn int
}

type Option func(*Server)

func WithPort(port int) Option {
    return func(s *Server) { s.port = port }
}

func WithTimeout(t time.Duration) Option {
    return func(s *Server) { s.timeout = t }
}

func NewServer(host string, opts ...Option) *Server {
    s := &Server{
        host:    host,
        port:    8080,           // default
        timeout: 30 * time.Second,
        maxConn: 100,
    }
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// Usage
s := NewServer("localhost", WithPort(9090), WithTimeout(10*time.Second))
```

**Interview Q:** When do you use pointer receiver vs value receiver? **Answer:** Pointer when you need to modify the receiver or avoid copying large structs. Value when the struct is small and you don't need mutation.

**Interview Q:** What is the difference between embedding and inheritance? **Answer:** Go doesn't have inheritance. Embedding is composition with promoted methods. No polymorphism, no virtual methods. Prefer composition over inheritance.

---

## 3. Pointers

- [x] `&` (address-of), `*` (dereference)

- [x] Pointer to struct

- [x] Pointer receivers (mutate original)

- [x] `new()` vs `make()`

- [x] Pointer to slice / map

- [x] Memory management (stack vs heap, escape analysis)

#### Pointer Basics

```go
// Pointer holds memory address of a value
x := 10
p := &x      // p is *int, points to x's address
fmt.Println(*p)   // 10 (dereference - get value)
fmt.Println(&x)   // 0xc0000b2008 (address)

// Modify via pointer
*p = 20
fmt.Println(x)   // 20 (x changed!)

// Pointer declaration
var p *int        // nil pointer (zero value for pointers)
p = &x            // now points to x

// Pointer to pointer
pp := &p          // **int
**pp = 30         // modify x through pp
```

#### Pointer to Struct

```go
type User struct {
    Name string
    Age  int
}

// Pointer to struct (common pattern)
u := &User{Name: "Alice", Age: 30}
// Equivalent to:
u := new(User)
u.Name = "Alice"
u.Age = 30

// Access fields (auto-dereferenced)
fmt.Println(u.Name)  // "Alice" (no need for (*u).Name)

// Nil pointer check
var u *User
if u != nil {
    fmt.Println(u.Name)
}

// Struct literal pointer
u := &User{"Bob", 25}       // positional
u := &User{Name: "Bob"}     // named fields

// Return pointer from function (safe in Go!)
func newUser(name string) *User {
    return &User{Name: name}  // Go escapes to heap
}

u := newUser("Charlie")
fmt.Println(u.Name)
```

#### Pointer Receivers (Mutate Original)

```go
type Counter struct {
    value int
}

// Pointer receiver - modifies original
func (c *Counter) Increment() {
    c.value++
}

// Value receiver - works on copy
func (c Counter) Value() int {
    return c.value
}

// Usage
c := &Counter{value: 0}
c.Increment()
c.Increment()
fmt.Println(c.Value())  // 2

// Method chaining with pointer receivers
func (c *Counter) Add(n int) *Counter {
    c.value += n
    return c  // return self for chaining
}

c.Add(5).Add(3).Increment()  // fluent API
```

#### new() vs make()

```go
// new(T) - allocates zero-value, returns *T
p := new(int)     // *int pointing to 0
s := new([]int)   // *[]int pointing to nil slice

fmt.Println(*p)   // 0
fmt.Println(*s)   // [] (nil)

// make(T, args) - initializes internal fields, returns T
s := make([]int, 5)    // initialized slice
m := make(map[string]int)  // initialized map
ch := make(chan int, 10)    // initialized channel

// Key difference:
// new: returns pointer to zero-value
// make: returns initialized value (slice, map, or channel)

// When to use:
// - new: any type when you need pointer
// - make: only for slice, map, channel (reference types)

// Examples
p := new(int)         // *int, value 0
s := make([]int, 0)   // []int, len 0, cap 0
m := make(map[string]int)  // empty map

// Cannot use make with structs
// u := make(User)  // ERROR: make not supported for User
u := new(User)       // OK: returns *User
```

#### Pointer to Slice / Map

```go
// Slice is already a reference (header: ptr, len, cap)
// You rarely need pointer to slice

// Slice header is 3 words: {ptr, len, cap}
// When you pass slice to function, you copy this header
// If you append and header grows, original caller still has old header

// Example: why slice append might not work as expected
func appendItem(s []int, v int) {
    s = append(s, v)  // modifies local copy of header!
}

s := []int{1, 2, 3}
appendItem(s, 4)
fmt.Println(len(s))  // 3 (unchanged!)

// Solution 1: return new slice
func appendItem(s []int, v int) []int {
    return append(s, v)
}
s = appendItem(s, 4)  // must use return value

// Solution 2: pointer to slice (rare, not idiomatic)
func appendItem(s *[]int, v int) {
    *s = append(*s, v)
}
appendItem(&s, 4)

// Map is already a pointer (to hash table)
func addItem(m map[string]int, k string, v int) {
    m[k] = v  // modifies original map!
}

// No need for pointer to map
m := map[string]int{}
addItem(m, "key", 1)  // works directly
```

#### Memory Management (Stack vs Heap)

```go
// Stack: automatic allocation, fast, limited size
// Heap: dynamic allocation, GC managed, slower

// Go compiler decides where to allocate (escape analysis)
// If pointer escapes function -> heap allocation

// Escape to heap examples:

// 1. Returning pointer
func newUser() *User {
    u := User{Name: "Alice"}  // escapes to heap
    return &u
}

// 2. Sending pointer to channel
func send(ch chan *User) {
    u := &User{}  // escapes to heap
    ch <- u
}

// 3. Pointer captured by closure
func process() func() {
    x := 10
    return func() {  // x escapes to heap
        fmt.Println(x)
    }
}

// 4. Interface value (pointer stored in interface)
func getValue() interface{} {
    x := 42
    return x  // x escapes to heap
}

// Stack allocation (no escape):
func process() {
    x := 10  // stays on stack
    y := x + 5  // stays on stack
    fmt.Println(y)
}

// Check escape analysis:
// go build -gcflags="-m" main.go
// Output shows what escapes to heap
```

#### Nil Pointer Safety

```go
// Always check nil before dereferencing
func processUser(u *User) {
    if u == nil {
        return
    }
    fmt.Println(u.Name)
}

// Safe method calls (Go auto-dereferences)
func (u *User) Name() string {
    if u == nil {
        return ""
    }
    return u.Name
}

var u *User
fmt.Println(u.Name())  // "" (safe, no panic)

// Nil pointer dereference PANIC
var p *int
// fmt.Println(*p)  // PANIC: runtime error: invalid memory address

// Safe pattern
if p != nil {
    fmt.Println(*p)
}
```

**Interview Q:** What is the difference between `new()` and `make()`? **Answer:** `new(T)` allocates zero-value T, returns `*T`. `make(T, ...)` initializes slice/map/channel internals, returns `T` (not pointer).

**Interview Q:** When does a variable escape to heap? **Answer:** When pointer escapes function scope: returned from function, sent on channel, captured by closure, stored in interface.

**Interview Q:** Is Go pass-by-value or pass-by-reference? **Answer:** Always pass-by-value. Pointers, slices, maps, channels contain reference data (pointer to underlying array/hash table), but the reference wrapper itself is copied.

---

## 4. Interfaces

- [x] Interface definition, implicit satisfaction

- [x] Interface satisfaction at compile time

- [x] Type assertion `v.(Type)`

- [x] Type switch

- [x] Empty interface `interface{}`

- [x] Interface composition (embedding interfaces)

- [x] Interface values (type + value pair)

- [x] Nil interface trap

- [x] Stringer interface (`String() string`)

- [x] io.Reader, io.Writer, io.Closer

#### Interface Definition and Implicit Satisfaction

```go
// Interface: collection of method signatures
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Implicit satisfaction - NO explicit "implements" keyword!
type Circle struct {
    radius float64
}

// Circle satisfies Shape by implementing both methods
func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.radius
}

// Compile-time check (optional but recommended)
var _ Shape = (*Circle)(nil)  // fails if Circle doesn't implement Shape

// Usage
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

c := Circle{radius: 5}
printShapeInfo(c)  // Circle implicitly satisfies Shape
```

#### Interface Values (Type + Value Pair)

```go
// Interface value = (type, value) pair
var s Shape
fmt.Println(s == nil)  // true (both type and value nil)

s = Circle{radius: 5}
fmt.Println(s == nil)  // false (has value)

// Interface with nil value (NOT nil interface!)
var c *Circle = nil
s = c  // s has type (*Circle) but value nil
fmt.Println(s == nil)  // FALSE! Interface has type info

// Check both type and value
func describe(s Shape) {
    if s == nil {
        fmt.Println("nil interface")
        return
    }
    // s != nil but might have nil value
    fmt.Printf("type: %T, value: %v\n", s, s)
}
```

#### Type Assertion

```go
// Type assertion extracts concrete type from interface
var iface interface{} = "hello"

// Safe assertion with comma-ok
str, ok := iface.(string)
if ok {
    fmt.Println(str)  // "hello"
}

// Unsafe assertion (panics if wrong type)
str := iface.(string)  // works
// num := iface.(int)  // PANIC: interface conversion: interface {} is string, not int

// Type switch (better for multiple types)
switch v := iface.(type) {
case string:
    fmt.Println("string:", v)
case int:
    fmt.Println("int:", v)
case bool:
    fmt.Println("bool:", v)
case nil:
    fmt.Println("nil")
default:
    fmt.Printf("unknown type: %T\n", v)
}
```

#### Interface Composition (Embedding)

```go
// Combine small interfaces into larger ones
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Compose interfaces
type ReadWriter interface {
    Reader
    Writer
}

// Equivalent to:
type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
}

// Real-world: io package
// io.Reader, io.Writer, io.Closer, io.Seeker
// io.ReadWriter = Reader + Writer
// io.ReadCloser = Reader + Closer
// io.ReadWriteCloser = Reader + Writer + Closer
```

#### Empty Interface (interface{})

```go
// Empty interface accepts any type
func printAny(v interface{}) {
    fmt.Printf("type: %T, value: %v\n", v, v)
}

printAny(42)        // type: int, value: 42
printAny("hello")   // type: string, value: hello
printAny(Circle{})  // type: main.Circle, value: {5}

// Go 1.18+: use `any` instead of `interface{}`
func printAny(v any) {
    fmt.Printf("type: %T, value: %v\n", v, v)
}

// Type assertion on empty interface
func process(v any) {
    switch val := v.(type) {
    case string:
        fmt.Println("string:", strings.ToUpper(val))
    case int:
        fmt.Println("int:", val*2)
    case []int:
        fmt.Println("slice:", len(val))
    }
}

// When to use interface{}:
// - JSON encoding/decoding
// - Container types (maps with mixed values)
// - Legacy code compatibility
// - Generic code (pre-1.18)
```

#### Nil Interface Trap

```go
// This is a common gotcha!

type MyError struct {
    Msg string
}

func (e *MyError) Error() string {
    return e.Msg
}

func getError() error {
    var err *MyError = nil
    return err  // returns error with type (*MyError), value nil
}

// WRONG: interface is NOT nil
err := getError()
if err != nil {
    fmt.Println("error exists")  // THIS PRINTS!
    // err has type (*MyError), so != nil
}

// CORRECT: check with errors.Is or type assertion
func getError() error {
    var err *MyError = nil
    if err == nil {
        return nil  // return untyped nil
    }
    return err
}

// Or check differently
func isError(err error) bool {
    if err == nil {
        return false
    }
    // Check if underlying value is nil
    v := reflect.ValueOf(err)
    if v.Kind() == reflect.Ptr {
        return !v.IsNil()
    }
    return true
}
```

#### Standard Library Interfaces

```go
// Stringer - most common custom interface
type Stringer interface {
    String() string
}

type User struct {
    Name string
    Age  int
}

func (u User) String() string {
    return fmt.Sprintf("%s (age %d)", u.Name, u.Age)
}

u := User{"Alice", 30}
fmt.Println(u)  // "Alice (age 30)" - auto-calls String()

// io.Reader and io.Writer
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Anything that implements these can be used with io functions
// os.File implements both io.Reader and io.Writer
// bytes.Buffer implements both
// http.Response.Body implements io.Reader

// io.Closer
type Closer interface {
    Close() error
}

// Defer Close pattern
func process() error {
    f, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer f.Close()  // Close() exists because File implements io.Closer

    // process file...
    return nil
}

// error interface
type error interface {
    Error() string
}

// Any type with Error() string method satisfies error
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

#### Interface Best Practices

```go
// 1. Keep interfaces small (1-3 methods)
type Storer interface {
    Store(key string, value interface{}) error
    Retrieve(key string) (interface{}, error)
}

// 2. Define interfaces where they're consumed, not implemented
// BAD: package defines interface it implements
// GOOD: package that USES the interface defines it

// 3. Accept interfaces, return structs
func Process(r io.Reader) ([]byte, error) {
    return io.ReadAll(r)  // accepts any reader
}

// Returns concrete type, not interface
func NewFile(name string) (*os.File, error) {
    return os.Open(name)
}

// 4. Use compile-time checks
var _ io.Reader = (*MyReader)(nil)

// 5. Avoid interface{} when possible (use generics in Go 1.18+)
// BAD: func Process(v interface{}) interface{}
// GOOD: func Process[T any](v T) T
```

**Interview Q:** What is the "nil interface" problem? **Answer:** A nil pointer stored in an interface is NOT nil interface. `var p *MyStruct = nil; var i interface{} = p` -&gt; `i != nil` because interface has type info.

**Interview Q:** When should you use `interface{}` vs generics? **Answer:** `interface{}` for runtime flexibility. Generics (Go 1.18+) for compile-time type safety with type parameters.

**Interview Q:** What is the empty interface good for? **Answer:** Accepting any type: JSON encoding, maps with mixed values, legacy code, generic containers. Use `any` (Go 1.18+).

---

## 5. Concurrency (CRITICAL for 2026 Interviews)

### 5.1 Goroutines

- [x] `go` keyword

- [x] Lightweight threads (2KB stack)

- [x] Fork-join model

- [x] Goroutine scheduling (GMP model: Goroutine, OS Thread, Processor)

- [x] `runtime.GOMAXPROCS()`

- [x] Goroutine leak (goroutine that never exits)

#### Goroutine Basics

```go
// Goroutine: lightweight concurrent function execution
func main() {
    go sayHello()  // non-blocking, starts new goroutine

    // Wait for goroutine to finish (simple way)
    time.Sleep(time.Second)
}

func sayHello() {
    fmt.Println("Hello from goroutine!")
}

// Better: use WaitGroup
func main() {
    var wg sync.WaitGroup
    wg.Add(1)

    go func() {
        defer wg.Done()
        fmt.Println("Hello from goroutine!")
    }()

    wg.Wait()  // blocks until goroutine finishes
}
```

#### GMP Scheduling Model

```
GMP Model:
- G: Goroutine (2KB stack, grows/shrinks dynamically)
- M: Machine (OS thread)
- P: Processor (logical processor, holds local run queue)

runtime.GOMAXPROCS(n)  // set number of P's (default = CPU cores)

┌─────────────────────────────────────┐
│           Global Run Queue           │
└─────────────────────────────────────┘
        ↑              ↑
   ┌────┴────┐    ┌────┴────┐
   │    P1   │    │    P2   │
   │ ┌─────┐ │    │ ┌─────┐ │
   │ │ G1  │ │    │ │ G3  │ │
   │ │ G2  │ │    │ │ G4  │ │
   │ └─────┘ │    │ └─────┘ │
   └────┬────┘    └────┬────┘
        ↑              ↑
   ┌────┴────┐    ┌────┴────┐
   │   M1    │    │   M2    │
   │ (OS     │    │ (OS     │
   │ Thread) │    │ Thread) │
   └─────────┘    └─────────┘
```

```go
// Control parallelism
runtime.GOMAXPROCS(4)  // use 4 CPU cores

// Number of CPUs
fmt.Println(runtime.NumCPU())  // e.g., 8

// Number of goroutines
fmt.Println(runtime.NumGoroutine())  // e.g., 5

// Yield to other goroutines
runtime.Gosched()

// Bring goroutine to calling thread
runtime.LockOSThread()

// Debugging: detect race conditions
// go run -race main.go
```

#### Goroutine Leak

```go
// Goroutine leak: goroutine blocks forever, never exits

// LEAK: goroutine waiting on channel that nobody sends to
func leak() {
    ch := make(chan int)
    go func() {
        val := <-ch  // blocks forever if nobody sends
        fmt.Println(val)
    }()
    // function returns, goroutine still alive!
}

// LEAK: goroutine waiting for condition that never happens
func leak() {
    done := make(chan bool)
    go func() {
        for {
            select {
            case <-done:
                return
            default:
                // work...
            }
        }
    }()
    // done is never closed, goroutine runs forever
}

// PREVENT: use context or done channel
func safe() {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    go func() {
        select {
        case <-ctx.Done():
            fmt.Println("cancelled")
            return
        }
    }()
}

// DETECT: runtime.NumGoroutine()
// If count keeps growing, you have a leak
```

---

### 5.2 Channels

- [x] Unbuffered channels (synchronous)

- [x] Buffered channels (asynchronous until full)

- [x] Directional channels (`chan<-`, `<-chan`)

- [x] Closing channels (`close(ch)`)

- [x] Ranging over channels

- [x] Channel is a reference type

- [x] Channel internals (hchan struct)

- [x] `struct{}` for signaling (zero memory)

#### Channel Basics

```go
// Channel: typed conduit for goroutine communication
ch := make(chan int)      // unbuffered
ch <- 42                  // send
val := <-ch              // receive
close(ch)                 // close

// Buffered channel
ch := make(chan int, 5)   // buffer of 5
ch <- 10                  // doesn't block until buffer full
val := <-ch               // doesn't block until empty

// Unbuffered: sender blocks until receiver is ready
// Buffered: sender blocks when buffer full

// Channel is reference type
func producer(ch chan<- int) {  // send-only
    ch <- 42
}

func consumer(ch <-chan int) {  // receive-only
    val := <-ch
    fmt.Println(val)
}

ch := make(chan int)
go producer(ch)
go consumer(ch)
```

#### Closing Channels

```go
// Close: signal no more values
ch := make(chan int, 5)
ch <- 1
ch <- 2
ch <- 3
close(ch)

// Receive from closed channel: returns zero value
val, ok := <-ch  // ok is false when closed and empty

// Range over closed channel
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
close(ch)

for val := range ch {  // exits when channel closed
    fmt.Println(val)
}

// Send on closed channel: PANIC!
// close(ch)  // OK
// ch <- 42   // PANIC: send on closed channel

// Only sender should close
// Multiple senders: use sync.WaitGroup or atomic counter
```

#### Directional Channels

```go
// Send-only
func producer(ch chan<- int) {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)  // sender can close
}

// Receive-only
func consumer(ch <-chan int) {
    for val := range ch {
        fmt.Println(val)
    }
}

// Bidirectional
func process(ch chan int) {
    ch <- 42
    val := <-ch
}
```

#### struct{} for Signaling (Zero Memory)

```go
// Empty struct has zero size - use for pure signaling
done := make(chan struct{})

go func() {
    // do work...
    close(done)  // signal completion
}()

<-done  // wait for completion

// Better than bool channel (0 bytes vs 1 byte)
doneBool := make(chan bool)   // 1 byte per signal
doneStruct := make(chan struct{})  // 0 bytes per signal

// In select
select {
case <-done:
    return
case <-time.After(time.Second):
    return
}
```

#### Channel Patterns

```go
// Fan-out: one goroutine, multiple workers
func fanOut(jobs <-chan Job, workers int) <-chan Result {
    results := make(chan Result)
    var wg sync.WaitGroup

    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs {
                results <- process(job)
            }
        }()
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    return results
}

// Fan-in: multiple channels into one
func fanIn(channels ...<-chan Result) <-chan Result {
    var wg sync.WaitGroup
    merged := make(chan Result)

    for _, ch := range channels {
        wg.Add(1)
        go func(ch <-chan Result) {
            defer wg.Done()
            for val := range ch {
                merged <- val
            }
        }(ch)
    }

    go func() {
        wg.Wait()
        close(merged)
    }()

    return merged
}
```

---

### 5.3 Select Statement

- [x] Multiple channel operations

- [x] Timeout with `time.After`

- [x] Default case (non-blocking)

- [x] Random selection when multiple ready

- [x] `select{}` blocks forever (keep-alive)

#### Select Basics

```go
// Select: wait on multiple channel operations
select {
case msg := <-ch1:
    fmt.Println("from ch1:", msg)
case msg := <-ch2:
    fmt.Println("from ch2:", msg)
case ch3 <- value:
    fmt.Println("sent to ch3")
case <-time.After(5 * time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("no ready channels")
}

// Random selection when multiple ready
ch1 := make(chan string, 1)
ch2 := make(chan string, 1)
ch1 <- "one"
ch2 <- "two"

select {
case msg := <-ch1:
    fmt.Println(msg)  // randomly "one" or "two"
case msg := <-ch2:
    fmt.Println(msg)
}

// Timeout pattern
select {
case result := <-resultCh:
    fmt.Println("result:", result)
case <-time.After(3 * time.Second):
    fmt.Println("timeout!")
}

// Non-blocking operations
select {
case msg := <-ch:
    fmt.Println("received:", msg)
default:
    fmt.Println("nothing available")
}

// Block forever (keep process alive)
select {}  // blocks forever
```

---

### 5.4 Sync Package

- [x] `sync.Mutex` (mutual exclusion)

- [x] `sync.WaitGroup` (wait for goroutines)

- [x] `sync.RWMutex` (read-write lock)

- [x] `sync.Once` (run once, thread-safe)

- [x] `sync.Map` (concurrent map)

- [x] `sync.Pool` (object pool for GC reduction)

- [x] `sync.Cond` (condition variable)

#### sync.Mutex

```go
// Mutex: protect shared state from concurrent access
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) Get() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

// Common pattern: lock in defer
func (c *SafeCounter) Update(n int) {
    c.mu.Lock()
    defer c.mu.Unlock()  // guaranteed unlock
    c.count += n
}
```

#### sync.RWMutex

```go
// RWMutex: allows concurrent reads, exclusive writes
type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

func (c *Cache) Get(key string) (string, bool) {
    c.mu.RLock()         // multiple goroutines can read
    defer c.mu.RUnlock()
    val, ok := c.data[key]
    return val, ok
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()          // exclusive access
    defer c.mu.Unlock()
    c.data[key] = value
}
```

#### sync.Once

```go
// Once: execute function exactly once, even with multiple goroutines
var once sync.Once
var instance *Database

func GetDB() *Database {
    once.Do(func() {
        instance = connectDB()  // called only once
    })
    return instance
}

// Thread-safe initialization
func initConfig() *Config {
    var once sync.Once
    var config *Config

    once.Do(func() {
        config = loadConfig()
    })
    return config
}
```

#### sync.Pool

```go
// Pool: reuse objects to reduce GC pressure
var bufPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func process(data []byte) {
    buf := bufPool.Get().(*bytes.Buffer)
    defer bufPool.Put(buf)

    buf.Reset()
    buf.Write(data)
    // use buf...
}
```

#### sync.Map

```go
// Map: concurrent-safe map (better for read-heavy)
var m sync.Map

// Store
m.Store("alice", 30)

// Load
val, ok := m.Load("alice")
if ok {
    fmt.Println(val.(int))
}

// Delete
m.Delete("alice")

// Range
m.Range(func(key, value interface{}) bool {
    fmt.Printf("%s: %v\n", key, value)
    return true  // continue iteration
})

// LoadOrStore: load existing or store new
actual, loaded := m.LoadOrStore("bob", 25)
// If "bob" exists: actual=existing, loaded=true
// If not: actual=25, loaded=false
```

---

### 5.5 Concurrency Patterns (MUST KNOW)

- [x] Worker Pool

- [x] Fan-Out / Fan-In

- [x] Pipeline

- [x] Producer-Consumer

- [x] Fan-out/Fan-in with error group (`errgroup`)

- [x] Rate limiting (token bucket)

- [x] Circuit breaker pattern

#### Worker Pool Pattern

```go
// Worker Pool: fixed number of workers processing jobs
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("worker %d processing job %d\n", id, job)
        time.Sleep(time.Second)  // simulate work
        results <- job * 2
    }
}

func main() {
    numWorkers := 3
    numJobs := 10

    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // Start workers
    for w := 0; w < numWorkers; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 0; j < numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for r := 0; r < numJobs; r++ {
        fmt.Println(<-results)
    }
}
```

#### Fan-Out / Fan-In

```go
// Fan-out: distribute work to multiple goroutines
// Fan-in: merge results into single channel

func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func fanIn(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    merged := make(chan int)

    for _, ch := range channels {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for val := range ch {
                merged <- val
            }
        }(ch)
    }

    go func() {
        wg.Wait()
        close(merged)
    }()

    return merged
}

// Usage
ch1 := square(generate(1, 2, 3))
ch2 := square(generate(4, 5, 6))
merged := fanIn(ch1, ch2)

for val := range merged {
    fmt.Println(val)
}
```

#### Pipeline Pattern

```go
// Pipeline: chain of stages connected by channels

func stage1(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * 2  // stage 1: double
        }
        close(out)
    }()
    return out
}

func stage2(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n + 10  // stage 2: add 10
        }
        close(out)
    }()
    return out
}

// Chain stages
func pipeline() <-chan int {
    ch := generate(1, 2, 3, 4, 5)
    ch = stage1(ch)    // double: 2, 4, 6, 8, 10
    ch = stage2(ch)    // add 10: 12, 14, 16, 18, 20
    return ch
}
```

#### errgroup Pattern

```go
// errgroup: goroutine with error handling
import "golang.org/x/sync/errgroup"

func processItems(items []Item) error {
    g, ctx := errgroup.WithContext(context.Background())

    for _, item := range items {
        item := item  // capture loop variable
        g.Go(func() error {
            if err := ctx.Err(); err != nil {
                return err  // respect cancellation
            }
            return processItem(item)
        })
    }

    return g.Wait()  // returns first non-nil error
}

// With concurrency limit
func processWithLimit(items []Item) error {
    g, ctx := errgroup.WithContext(context.Background())
    g.SetLimit(10)  // max 10 concurrent goroutines

    for _, item := range items {
        item := item
        g.Go(func() error {
            return processItem(item)
        })
    }

    return g.Wait()
}
```

#### Rate Limiting

```go
// Token bucket algorithm
import "golang.org/x/time/rate"

// Create limiter: 10 events/sec, burst of 20
limiter := rate.NewLimiter(10, 20)

func handleRequest() error {
    if !limiter.Allow() {
        return fmt.Errorf("rate limit exceeded")
    }
    // process request...
    return nil
}

// Per-request rate limiting
func processWithRateLimit(requests []Request) {
    limiter := rate.NewLimiter(rate.Limit(1), 1)  // 1 per second

    for _, req := range requests {
        limiter.Wait(context.Background())  // block until allowed
        process(req)
    }
}

// Channel-based rate limiter
func rateLimiter(rps int) <-chan time.Time {
    throbber := time.NewTicker(time.Second / time.Duration(rps))
    return throbber.C
}
```

#### Circuit Breaker Pattern

```go
// Circuit Breaker: prevent cascading failures
type CircuitBreaker struct {
    mu           sync.Mutex
    failures     int
    maxFailures  int
    resetTimeout time.Duration
    lastFailure  time.Time
    state        string  // "closed", "open", "half-open"
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mu.Lock()
    if cb.state == "open" {
        if time.Since(cb.lastFailure) > cb.resetTimeout {
            cb.state = "half-open"
        } else {
            cb.mu.Unlock()
            return fmt.Errorf("circuit breaker is open")
        }
    }
    cb.mu.Unlock()

    err := fn()

    cb.mu.Lock()
    defer cb.mu.Unlock()

    if err != nil {
        cb.failures++
        cb.lastFailure = time.Now()
        if cb.failures >= cb.maxFailures {
            cb.state = "open"
        }
        return err
    }

    cb.failures = 0
    cb.state = "closed"
    return nil
}

// Usage
cb := &CircuitBreaker{
    maxFailures:  3,
    resetTimeout: 30 * time.Second,
    state:        "closed",
}

err := cb.Call(func() error {
    return callExternalService()
})
```

---

## 6. Context

- [x] `context.Background()`

- [x] `context.WithCancel()`

- [x] `context.WithTimeout()`

- [x] `context.WithDeadline()`

- [x] `context.WithValue()` (request-scoped values)

- [x] Context propagation

- [x] Context cancellation patterns

- [x] `context.TODO()` vs `context.Background()`

#### Context Basics

```go
// Context: carries deadlines, cancellation signals, and request-scoped values

// Background: root context, no deadline, no values
ctx := context.Background()

// TODO: placeholder when context not available
ctx := context.TODO()

// WithCancel: cancel manually
ctx, cancel := context.WithCancel(context.Background())
defer cancel()  // always defer cancel

go func() {
    select {
    case <-ctx.Done():
        fmt.Println("cancelled:", ctx.Err())
        return
    }
}()

cancel()  // trigger cancellation
```

#### WithTimeout and WithDeadline

```go
// WithTimeout: automatic cancellation after duration
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// WithDeadline: cancel at specific time
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()

// Check if context is done
select {
case <-ctx.Done():
    fmt.Println("context cancelled:", ctx.Err())
default:
    fmt.Println("context still active")
}

// Use with database operations
func getUser(ctx context.Context, id string) (*User, error) {
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    var user User
    err := db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", id).Scan(&user)
    return &user, err
}
```

#### Context Propagation

```go
// Pass context through function calls
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()  // get context from request

    user, err := service.GetUser(ctx, id)
    if err != nil {
        // context cancelled means client disconnected
        if ctx.Err() == context.Canceled {
            return  // client gone, don't respond
        }
        http.Error(w, err.Error(), 500)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func (s *Service) GetUser(ctx context.Context, id string) (*User, error) {
    // Add timeout if not already set
    if _, ok := ctx.Deadline(); !ok {
        var cancel context.CancelFunc
        ctx, cancel = context.WithTimeout(ctx, 3*time.Second)
        defer cancel()
    }

    return s.repo.FindByID(ctx, id)
}
```

#### WithValue (Request-Scoped Values)

```go
// WithValue: store request-scoped data (tracing, auth, etc.)
type contextKey string

const (
    userIDKey contextKey = "userID"
    traceIDKey contextKey = "traceID"
)

// Set values
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userID := extractUserID(r)
        ctx := context.WithValue(r.Context(), userIDKey, userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Get values
func handler(w http.ResponseWriter, r *http.Request) {
    userID, ok := r.Context().Value(userIDKey).(string)
    if !ok {
        http.Error(w, "unauthorized", 401)
        return
    }
    fmt.Println("user:", userID)
}

// When to use WithValue:
// ✓ Request ID / Trace ID
// ✓ Authenticated user ID
// ✓ Request metadata
// ✗ Function parameters (don't pass optional args via context)
// ✗ Database connection strings (use config instead)
```

#### Context Best Practices

```go
// 1. Always pass context as first parameter
func DoWork(ctx context.Context, data []byte) error { ... }

// 2. Don't store context in structs
// BAD
type Service struct {
    ctx context.Context  // don't do this
}

// GOOD
func (s *Service) DoWork(ctx context.Context) error { ... }

// 3. Don't pass nil context
// BAD
DoWork(nil, data)

// GOOD
DoWork(context.Background(), data)

// 4. Check context cancellation
func process(ctx context.Context) error {
    for _, item := range items {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
        }
        processItem(item)
    }
    return nil
}

// 5. Context values should be immutable
// BAD: modify existing context
ctx = context.WithValue(ctx, key, newVal)

// GOOD: create new context
newCtx := context.WithValue(ctx, key, newVal)
```

**Interview Q:** When should you NOT use `context.WithValue()`? **Answer:** For function parameters, not request metadata. Don't pass optional parameters via context. Use for tracing, auth, cancellation only.

**Interview Q:** What is the difference between `context.Background()` and `context.TODO()`? **Answer:** Functionally identical. `Background()` for main/init/test. `TODO()` as placeholder when context not yet available.

**Interview Q:** How do you detect if context is cancelled? **Answer:** `select { case <-ctx.Done(): ... }` or check `ctx.Err()` for error type.

---

## 7. Project Structure

### 7.1 Standard Layout

- [x] `cmd/` - entry points

- [x] `internal/` - private packages

- [x] `pkg/` - public packages (optional)

- [x] `go.mod` - module definition

- [x] `go.sum` - dependency checksums

- [x] `go mod tidy`, `go mod vendor`

- [x] Monorepo vs polyrepo

#### Standard Layout

```
project/
├── cmd/
│   └── api/
│       └── main.go           # entry point
├── internal/
│   ├── config/
│   │   └── config.go         # private config
│   ├── db/
│   │   └── mongo.go          # private DB
│   └── notes/
│       ├── handler.go
│       ├── repo.go
│       └── model.go          # private feature
├── pkg/
│   └── logger/
│       └── logger.go         # public package
├── migrations/
│   └── 001_init.up.sql
├── api/
│   └── openapi.yaml
├── docker/
│   └── Dockerfile
├── go.mod
├── go.sum
├── docker-compose.yaml
└── README.md
```

#### go.mod and Module System

```go
// go.mod: module definition
module github.com/user/project

go 1.22

require (
    github.com/gin-gonic/gin v1.9.1
    go.mongodb.org/mongo-driver v1.13.1
)

// require: direct dependencies
// require indirect: transitive dependencies
// go.sum: checksums for verification

// Common commands
// go mod init <module>       - initialize module
// go mod tidy                - add missing, remove unused
// go get <package>@version   - add/update dependency
// go mod download            - download dependencies
// go mod vendor              - copy dependencies to vendor/
// go mod graph               - show dependency graph
// go mod verify              - verify checksums
```

#### internal/ Directory

```go
// internal/: Go enforces encapsulation
// Packages inside internal/ cannot be imported by code
// outside the parent module

// Structure
project/
├── internal/
│   ├── auth/      # only project code can import
│   └── db/        # only project code can import
├── pkg/
│   └── logger/    # anyone can import
└── cmd/
    └── api/
        └── main.go  # can import internal/

// Benefit: prevents external code from using private packages
// Use for: business logic, database, config
// Don't use for: reusable libraries
```

#### Monorepo vs Polyrepo

```go
// Monorepo: multiple services in one repository
repo/
├── services/
│   ├── auth/
│   ├── booking/
│   └── payment/
├── libs/
│   └── shared/
├── go.work           # Go workspace (1.18+)
└── go.work.sum

// go.work: allows multiple modules in one repo
go 1.22

use (
    ./services/auth
    ./services/booking
    ./libs/shared
)

// Polyrepo: separate repository per service
// auth-service/     (separate repo)
// booking-service/  (separate repo)
// shared-lib/       (separate repo)
```

---

### 7.2 Package Design

- [x] Package naming (lowercase, single word)

- [x] Exported (capital) vs unexported (lowercase)

- [x] Constructor pattern (`NewHandler`, `NewRepo`)

- [x] Package as namespace vs package as abstraction

- [x] Avoiding circular dependencies

- [x] Package-level init

#### Package Naming

```go
// Package names: lowercase, single word, no underscores
package user        // GOOD
package userutil    // GOOD
package user_util   // BAD

// Import path matches package name
import "github.com/user/project/internal/user"

// Aliases when needed
import (
    "fmt"
    mylog "github.com/user/project/pkg/logger"
)

// Package comments: package-level documentation
// Package user provides user management functionality.
package user
```

#### Exported vs Unexported

```go
package user

// Exported (capital): accessible outside package
type User struct {
    Name  string   // exported field
    Email string   // exported field
    age   int      // unexported field
}

func (u *User) GetName() string {  // exported method
    return u.Name
}

func (u *User) validate() {  // unexported method
    // internal validation
}

// Exported constructor
func NewUser(name, email string) *User {
    u := &User{Name: name, Email: email}
    u.validate()
    return u
}

// Unexported helper
func hashPassword(pw string) string {
    // internal implementation
}
```

#### Constructor Pattern

```go
// Always use New* constructors
type Handler struct {
    repo   *Repo
    logger *Logger
}

// Constructor: initializes and validates
func NewHandler(repo *Repo, logger *Logger) *Handler {
    if repo == nil {
        panic("repo cannot be nil")
    }
    return &Handler{repo: repo, logger: logger}
}

// Options pattern for many parameters
type Option func(*Handler)

func WithLogger(l *Logger) Option {
    return func(h *Handler) { h.logger = l }
}

func WithRepo(r *Repo) Option {
    return func(h *Handler) { h.repo = r }
}

func NewHandler(repo *Repo, opts ...Option) *Handler {
    h := &Handler{repo: repo, logger: defaultLogger}
    for _, opt := range opts {
        opt(h)
    }
    return h
}
```

#### Avoiding Circular Dependencies

```go
// BAD: circular dependency
// package A imports B
// package B imports A
// Compile error!

// Solution 1: Extract common interface to third package
// package interfaces
type Storer interface {
    Store(key string, val interface{}) error
    Retrieve(key string) (interface{}, error)
}

// package A
import "interfaces"
type Service struct {
    store interfaces.Storer
}

// package B
import "interfaces"
type RedisStore struct {
    // implements interfaces.Storer
}

// Solution 2: Dependency injection
// package handler
type Handler struct {
    store Store  // interface, not concrete type
}

// package main (composition root)
handler := NewHandler(redisStore)  // inject dependency
```

#### Package-level init

```go
// init(): runs at package initialization (before main)
package config

var AppConfig *Config

func init() {
    AppConfig = loadConfig()
}

// Multiple init functions allowed
func init() {
    // first init
}

func init() {
    // second init (runs after first)
}

// Init order:
// 1. Import dependencies (recursively)
// 2. Allocate package-level variables
// 3. Run init() functions (in import order)
// 4. Run main()

// Warning: init() makes testing harder
// Prefer explicit initialization in main or New* functions
```

---

## 8. HTTP & Web Development

### 8.1 net/http Package

- [x] `http.HandleFunc`, `http.ListenAndServe`

- [x] Request (`r.Method`, `r.URL.Query()`, `r.Body`)

- [x] Response (`w.Write()`, `w.WriteHeader()`, `w.Header().Set()`)

- [x] HTTP methods (GET, POST, PUT, PATCH, DELETE)

- [x] Status codes (200, 201, 400, 401, 403, 404, 500)

- [x] Middleware pattern

#### Basic HTTP Server

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Handler functions
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/users", usersHandler)

    // Start server
    fmt.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "World"
    }
    fmt.Fprintf(w, "Hello, %s!", name)
}
```

#### Request Handling

```go
func usersHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handleGetUsers(w, r)
    case http.MethodPost:
        handleCreateUser(w, r)
    case http.MethodPut:
        handleUpdateUser(w, r)
    case http.MethodDelete:
        handleDeleteUser(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
    // Query parameters
    name := r.URL.Query().Get("name")
    age := r.URL.Query().Get("age")

    // Path parameters (use mux or router for this)
    id := r.URL.Path[len("/users/"):]

    // Headers
    authHeader := r.Header.Get("Authorization")

    // Response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(users)
}
```

#### Middleware Pattern

```go
// Middleware: function that wraps handlers
type Middleware func(http.Handler) http.Handler

// Logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("Started %s %s", r.Method, r.URL.Path)

        next.ServeHTTP(w, r)

        log.Printf("Completed in %v", time.Since(start))
    })
}

// Authentication middleware
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        userID, err := validateToken(token)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Add user ID to context
        ctx := context.WithValue(r.Context(), userIDKey, userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Chain middleware
func chainMiddleware(handler http.Handler, middlewares ...Middleware) http.Handler {
    for i := len(middlewares) - 1; i >= 0; i-- {
        handler = middlewares[i](handler)
    }
    return handler
}

// Usage
handler := chainMiddleware(userHandler, loggingMiddleware, authMiddleware)
```

---

### 8.2 JSON Handling

- [x] `json.Marshal()` / `json.Unmarshal()`

- [x] `json.NewEncoder(w).Encode()`

- [x] `json.NewDecoder(r.Body).Decode()`

- [x] Struct tags (`json:"name"`, `binding:"required"`)

- [x] `omitempty` tag

- [x] `json:"-"` to skip field

- [x] Custom marshaler (`MarshalJSON()`)

- [x] `json.RawMessage` for raw JSON

#### JSON Encoding/Decoding

```go
import "encoding/json"

// Marshal: struct to JSON
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age,omitempty"`
    Pass  string `json:"-"`  // skip in JSON
}

user := User{Name: "Alice", Email: "alice@example.com"}
jsonBytes, err := json.Marshal(user)
// {"name":"Alice","email":"alice@example.com"}

// Marshal with indentation
jsonBytes, err := json.MarshalIndent(user, "", "  ")

// Unmarshal: JSON to struct
var u User
err := json.Unmarshal(jsonBytes, &u)

// Decode from request body
func createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    // use user...
}

// Encode to response
func respond(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
```

#### Custom Marshaler

```go
// Custom JSON marshaling
type Time struct {
    time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02"))), nil
}

// Custom unmarshaling
func (t *Time) UnmarshalJSON(data []byte) error {
    str := strings.Trim(string(data), `"`)
    parsed, err := time.Parse("2006-01-02", str)
    if err != nil {
        return err
    }
    t.Time = parsed
    return nil
}

// json.RawMessage: delay parsing
type Event struct {
    Type string          `json:"type"`
    Data json.RawMessage `json:"data"`  // don't parse yet
}

func processEvent(event Event) {
    switch event.Type {
    case "user.created":
        var user User
        json.Unmarshal(event.Data, &user)
    case "order.placed":
        var order Order
        json.Unmarshal(event.Data, &order)
    }
}
```

---

### 8.3 Gin Framework (Production)

- [x] `gin.Default()`, `gin.New()`

- [x] Middleware (`r.Use()`)

- [x] Route groups (`r.Group("/notes")`)

- [x] JSON binding (`c.ShouldBindJSON()`)

- [x] Path parameters (`c.Param("id")`)

- [x] Query parameters (`c.Query("name")`)

- [x] Custom middleware

- [x] CORS middleware

- [x] Rate limiting middleware

#### Gin Basics

```go
import "github.com/gin-gonic/gin"

func main() {
    // Default includes Logger and Recovery middleware
    r := gin.Default()

    // Or custom without defaults
    r := gin.New()
    r.Use(gin.Logger(), gin.Recovery())

    // Routes
    r.GET("/", homeHandler)
    r.GET("/users/:id", getUser)
    r.POST("/users", createUser)
    r.PUT("/users/:id", updateUser)
    r.DELETE("/users/:id", deleteUser)

    r.Run(":8080")
}

func getUser(c *gin.Context) {
    id := c.Param("id")
    c.JSON(200, gin.H{"id": id})
}

func createUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(201, user)
}
```

#### Route Groups and Middleware

```go
func main() {
    r := gin.Default()

    // Public routes
    public := r.Group("/api")
    {
        public.GET("/health", healthCheck)
        public.POST("/login", login)
    }

    // Protected routes
    protected := r.Group("/api")
    protected.Use(authMiddleware())
    {
        protected.GET("/users", getUsers)
        protected.POST("/users", createUser)
    }

    // Admin routes
    admin := r.Group("/admin")
    admin.Use(authMiddleware(), adminMiddleware())
    {
        admin.DELETE("/users/:id", deleteUser)
    }
}

func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
            return
        }
        // validate token...
        c.Next()  // continue to handler
    }
}
```

---

## 9. Database (MongoDB)

- [x] MongoDB connection with context

- [x] CRUD operations (InsertOne, Find, FindOneAndUpdate, DeleteOne)

- [x] BSON tags

- [x] Cursor handling (`cursor.Close()`)

- [x] Connection pooling

- [x] Transactions

- [x] Aggregation pipeline

- [x] Indexes

#### MongoDB Connection

```go
import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context) (*mongo.Client, error) {
    uri := "mongodb://localhost:27017"

    clientOpts := options.Client().
        ApplyURI(uri).
        SetMaxPoolSize(100).
        SetMinPoolSize(10)

    client, err := mongo.Connect(ctx, clientOpts)
    if err != nil {
        return nil, fmt.Errorf("mongo connect: %w", err)
    }

    // Verify connection
    if err := client.Ping(ctx, nil); err != nil {
        return nil, fmt.Errorf("mongo ping: %w", err)
    }

    return client, nil
}

// Graceful disconnect
func Disconnect(ctx context.Context, client *mongo.Client) error {
    return client.Disconnect(ctx)
}
```

#### CRUD Operations

```go
import (
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type User struct {
    ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name  string             `bson:"name" json:"name"`
    Email string             `bson:"email" json:"email"`
}

// Insert
func (r *Repo) Insert(ctx context.Context, user *User) error {
    user.ID = primitive.NewObjectID()
    _, err := r.coll.InsertOne(ctx, user)
    return err
}

// Find one
func (r *Repo) FindByID(ctx context.Context, id primitive.ObjectID) (*User, error) {
    var user User
    err := r.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
    if err == mongo.ErrNoDocuments {
        return nil, fmt.Errorf("user not found")
    }
    return &user, err
}

// Find all
func (r *Repo) FindAll(ctx context.Context) ([]User, error) {
    cursor, err := r.coll.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var users []User
    if err := cursor.All(ctx, &users); err != nil {
        return nil, err
    }
    return users, nil
}

// Update
func (r *Repo) Update(ctx context.Context, id primitive.ObjectID, update bson.M) error {
    _, err := r.coll.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
    return err
}

// Delete
func (r *Repo) Delete(ctx context.Context, id primitive.ObjectID) error {
    _, err := r.coll.DeleteOne(ctx, bson.M{"_id": id})
    return err
}
```

#### Aggregation Pipeline

```go
// Aggregation: complex data transformations
func (r *Repo) GetOrderStats(ctx context.Context) ([]bson.M, error) {
    pipeline := mongo.Pipeline{
        {{Key: "$match", Value: bson.M{"status": "completed"}}},
        {{Key: "$group", Value: bson.M{
            "_id":   "$userId",
            "total": bson.M{"$sum": "$amount"},
            "count": bson.M{"$sum": 1},
        }}},
        {{Key: "$sort", Value: bson.M{"total": -1}}},
        {{Key: "$limit", Value: 10}},
    }

    cursor, err := r.coll.Aggregate(ctx, pipeline)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var results []bson.M
    if err := cursor.All(ctx, &results); err != nil {
        return nil, err
    }
    return results, nil
}
```

#### Indexes

```go
// Create indexes for better query performance
func (r *Repo) EnsureIndexes(ctx context.Context) error {
    // Single field index
    indexModel := mongo.IndexModel{
        Keys:    bson.D{{Key: "email", Value: 1}},  // ascending
        Options: options.Index().SetUnique(true),
    }
    _, err := r.coll.Indexes().CreateOne(ctx, indexModel)

    // Compound index
    compoundIndex := mongo.IndexModel{
        Keys: bson.D{
            {Key: "status", Value: 1},
            {Key: "createdAt", Value: -1},
        },
    }
    _, err = r.coll.Indexes().CreateOne(ctx, compoundIndex)

    // Text index for search
    textIndex := mongo.IndexModel{
        Keys: bson.D{{Key: "name", Value: "text"}},
    }
    _, err = r.coll.Indexes().CreateOne(ctx, textIndex)

    return err
}
```

#### Transactions

```go
// Transactions: atomic operations across multiple collections
func (r *Repo) Transfer(ctx context.Context, from, to string, amount float64) error {
    session, err := r.client.StartSession()
    if err != nil {
        return err
    }
    defer session.EndSession(ctx)

    _, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
        // Debit sender
        _, err := r.accounts.UpdateOne(sessCtx,
            bson.M{"_id": from},
            bson.M{"$inc": bson.M{"balance": -amount}},
        )
        if err != nil {
            return nil, err
        }

        // Credit receiver
        _, err = r.accounts.UpdateOne(sessCtx,
            bson.M{"_id": to},
            bson.M{"$inc": bson.M{"balance": amount}},
        )
        if err != nil {
            return nil, err
        }

        return nil, nil
    })

    return err
}
```

**Interview Q:** Why use `context.WithTimeout` for database operations? **Answer:** Prevents hanging forever if DB is slow/down. Ensures request-scoped cancellation propagation.

---

## 10. Testing (MUST for 2026)

### 10.1 Unit Testing

- [x] `_test.go` files, `TestXxx(t *testing.T)`

- [x] `t.Errorf()` vs `t.Fatalf()`

- [x] Table-driven tests (Go convention)

- [x] Subtests `t.Run("name", func(t *testing.T))`

- [x] `go test ./...`, `-v`, `-run`, `-count`

- [x] Test coverage `go test -cover`

#### Basic Test Structure

```go
// math.go
func Add(a, b int) int {
    return a + b
}

// math_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}

// Run tests
// go test ./...
// go test -v ./...
// go test -run TestAdd ./...
// go test -count=1 ./...  // disable caching
```

#### Table-Driven Tests (Go Convention)

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"zero", 0, 0, 0},
        {"mixed", -1, 1, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d",
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

#### Subtests

```go
func TestUser(t *testing.T) {
    t.Run("creation", func(t *testing.T) {
        user := NewUser("Alice", 30)
        if user.Name != "Alice" {
            t.Errorf("Name = %s; want Alice", user.Name)
        }
    })

    t.Run("validation", func(t *testing.T) {
        user := NewUser("", 30)
        err := user.Validate()
        if err == nil {
            t.Error("expected error for empty name")
        }
    })

    t.Run("age validation", func(t *testing.T) {
        user := NewUser("Alice", -1)
        err := user.Validate()
        if err == nil {
            t.Error("expected error for negative age")
        }
    })
}

// Run specific subtest
// go test -run TestUser/validation
```

#### Test Helpers and Cleanup

```go
func setupTestDB(t *testing.T) *mongo.Database {
    t.Helper()  // marks as test helper (error reported at caller)
    t.Cleanup(func() {
        // cleanup after test
        client.Disconnect(context.Background())
    })

    db, err := ConnectTestDB()
    if err != nil {
        t.Fatal(err)
    }
    return db
}

func TestCreateUser(t *testing.T) {
    db := setupTestDB(t)
    // test using db...
}
```

#### Test Coverage

```go
// Run with coverage
// go test -cover ./...
// go test -coverprofile=coverage.out ./...
// go tool cover -html=coverage.out  // view in browser

// Coverage report
// go tool cover -func=coverage.out

// Example output:
// math.go:10:    Add      100.0%
// math.go:15:    Multiply 75.0%
```

---

### 10.2 Benchmarking

- [x] `BenchmarkXxx(b *testing.B)`

- [x] `b.N` loop

- [x] `go test -bench=.`

#### Basic Benchmark

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}

// Run benchmarks
// go test -bench=.
// go test -bench=BenchmarkAdd
// go test -bench=. -benchmem  // show memory allocations
// go test -bench=. -benchtime=5s  // run for 5 seconds

// Benchmark with setup
func BenchmarkProcess(b *testing.B) {
    data := generateTestData(1000)
    b.ResetTimer()  // exclude setup time

    for i := 0; i < b.N; i++ {
        Process(data)
    }
}

// Parallel benchmark
func BenchmarkParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            Add(2, 3)
        }
    })
}
```

---

### 10.3 Mocking & Test Doubles

- [x] Interface-based mocking

- [x] `testify/mock` library

- [x] Table-driven test with mocks

- [x] `httptest` package for HTTP testing

#### Interface-Based Mocking

```go
// Define interface
type UserRepository interface {
    FindByID(id string) (*User, error)
    Save(user *User) error
}

// Mock implementation
type MockUserRepo struct {
    users map[string]*User
    err   error
}

func NewMockUserRepo() *MockUserRepo {
    return &MockUserRepo{users: make(map[string]*User)}
}

func (m *MockUserRepo) FindByID(id string) (*User, error) {
    if m.err != nil {
        return nil, m.err
    }
    user, ok := m.users[id]
    if !ok {
        return nil, fmt.Errorf("user not found")
    }
    return user, nil
}

func (m *MockUserRepo) Save(user *User) error {
    if m.err != nil {
        return m.err
    }
    m.users[user.ID] = user
    return nil
}

// Test with mock
func TestGetUser(t *testing.T) {
    mock := NewMockUserRepo()
    mock.users["123"] = &User{Name: "Alice"}

    service := NewUserService(mock)
    user, err := service.GetUser("123")
    if err != nil {
        t.Fatal(err)
    }
    if user.Name != "Alice" {
        t.Errorf("Name = %s; want Alice", user.Name)
    }
}
```

#### HTTP Testing

```go
import "net/http/httptest"

func TestHelloHandler(t *testing.T) {
    // Create request
    req := httptest.NewRequest("GET", "/hello?name=Alice", nil)

    // Create response recorder
    w := httptest.NewRecorder()

    // Call handler
    helloHandler(w, req)

    // Check response
    if w.Code != 200 {
        t.Errorf("status = %d; want 200", w.Code)
    }

    expected := "Hello, Alice!"
    if w.Body.String() != expected {
        t.Errorf("body = %s; want %s", w.Body.String(), expected)
    }
}

// Test with Gin
func TestCreateUser(t *testing.T) {
    router := setupRouter()

    body := `{"name":"Alice","email":"alice@example.com"}`
    req := httptest.NewRequest("POST", "/users", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != 201 {
        t.Errorf("status = %d; want 201", w.Code)
    }

    var user User
    json.Unmarshal(w.Body.Bytes(), &user)
    if user.Name != "Alice" {
        t.Errorf("name = %s; want Alice", user.Name)
    }
}
```

#### Testify Mock Library

```go
import "github.com/stretchr/testify/mock"

// Mock struct
type MockRepo struct {
    mock.Mock
}

func (m *MockRepo) FindByID(id string) (*User, error) {
    args := m.Called(id)
    return args.Get(0).(*User), args.Error(1)
}

// Test with testify
func TestGetUser(t *testing.T) {
    mock := new(MockRepo)
    mock.On("FindByID", "123").Return(&User{Name: "Alice"}, nil)

    service := NewUserService(mock)
    user, err := service.GetUser("123")

    assert.NoError(t, err)
    assert.Equal(t, "Alice", user.Name)
    mock.AssertExpectations(t)
}
```

---

### 10.4 Integration Testing

- [x] Test containers (testcontainers-go)

- [x] Database testing with Docker

#### Testcontainers

```go
import (
    "testing"
    "github.com/testcontainers/testcontainers-go"
)

func setupTestContainer(t *testing.T) *mongo.Client {
    ctx := context.Background()

    req := testcontainers.ContainerRequest{
        Image:        "mongo:latest",
        ExposedPorts: []string{"27017/tcp"},
        WaitingFor:   wait.ForListeningPort("27017/tcp"),
    }

    container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: req,
        Started:          true,
    })
    if err != nil {
        t.Fatal(err)
    }

    t.Cleanup(func() {
        container.Terminate(ctx)
    })

    host, _ := container.Host(ctx)
    port, _ := container.MappedPort(ctx, "27017")

    client, _ := mongo.Connect(ctx, options.Client().ApplyURI(
        fmt.Sprintf("mongodb://%s:%s", host, port.Port()),
    ))

    return client
}

func TestWithDatabase(t *testing.T) {
    client := setupTestContainer(t)
    // test with real database...
}
```

**Interview Q:** What is the difference between `t.Errorf` and `t.Fatalf`? **Answer:** `Errorf` logs error but continues test. `Fatalf` logs error and immediately stops the test (`t.FailNow()`).

---

## 11. Generics (Go 1.18+)

- [x] Type parameters `[T any]`

- [x] Type constraints (`comparable`, `any`, `~int`)

- [x] Custom constraint interfaces

- [x] Generic functions vs generic types

- [x] When to use generics vs interfaces

- [x] Common patterns: Map, Filter, Reduce

#### Generic Functions

```go
// Type parameter with any constraint (accepts any type)
func Print[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

// Multiple type parameters
func Map[T any, U any](s []T, f func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s {
        result[i] = f(v)
    }
    return result
}

// Usage
ints := []int{1, 2, 3}
strs := Map(ints, func(i int) string {
    return strconv.Itoa(i)
})
// strs = ["1", "2", "3"]
```

#### Type Constraints

```go
// Built-in constraints
// any     - any type (alias for interface{})
// comparable - supports == and !=

// Number constraint (Go 1.21+ has math/big)
type Number interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
    ~float32 | ~float64
}

func Sum[T Number](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

// Ordered constraint (comparable + ordered)
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
    ~float32 | ~float64 | ~string
}

func Max[T Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// ~int means "underlying type is int" (includes type definitions)
type MyInt int
Sum([]MyInt{1, 2, 3})  // works because ~int matches MyInt
```

#### Custom Constraint Interfaces

```go
// Interface as constraint
type Stringer interface {
    String() string
}

func Join[T Stringer](items []T, sep string) string {
    strs := make([]string, len(items))
    for i, item := range items {
        strs[i] = item.String()
    }
    return strings.Join(strs, sep)
}

// Complex constraint
type Container[T any] interface {
    Add(item T)
    Get() T
    Size() int
}

// Generic type with constraint
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    var zero T
    if len(s.items) == 0 {
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

// Usage
intStack := &Stack[int]{}
intStack.Push(1)
intStack.Push(2)

strStack := &Stack[string]{}
strStack.Push("hello")
```

#### Common Generic Patterns

```go
// Filter
func Filter[T any](s []T, f func(T) bool) []T {
    result := []T{}
    for _, v := range s {
        if f(v) {
            result = append(result, v)
        }
    }
    return result
}

// Reduce
func Reduce[T any, U any](s []T, init U, f func(U, T) U) U {
    acc := init
    for _, v := range s {
        acc = f(acc, v)
    }
    return acc
}

// Contains
func Contains[T comparable](s []T, v T) bool {
    for _, item := range s {
        if item == v {
            return true
        }
    }
    return false
}

// Index
func Index[T comparable](s []T, v T) int {
    for i, item := range s {
        if item == v {
            return i
        }
    }
    return -1
}

// Keys/Values for maps
func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

func Values[K comparable, V any](m map[K]V) []V {
    vals := make([]V, 0, len(m))
    for _, v := range m {
        vals = append(vals, v)
    }
    return vals
}
```

#### Generics vs Interfaces

```go
// Use generics: compile-time type safety, no type assertion
func Max[T Ordered](a, b T) T {
    if a > b { return a }
    return b
}
Max(1, 2)      // works
Max("a", "b")  // works
// Max(1, "a") // compile error - type safety!

// Use interfaces: runtime polymorphism, multiple implementations
type Shape interface {
    Area() float64
}

func PrintArea(s Shape) {
    fmt.Println(s.Area())
}
// Circle and Rectangle can both implement Shape

// When to use which:
// Generics: type-safe collections, algorithms on same types
// Interfaces: different types with same behavior (polymorphism)
```

**Interview Q:** When should you use generics instead of interfaces? **Answer:** Generics for compile-time type safety. Interfaces for runtime polymorphism. Generics avoid type assertion overhead.

---

## 12. Advanced Topics (2026 Interview Differentiators)

### 12.1 Memory & Performance

- [x] Escape analysis (`go build -gcflags="-m"`)

- [x] Stack vs heap allocation

- [x] `sync.Pool` for object reuse

- [x] String builder vs concatenation

- [x] Pre-allocate slices `make([]T, 0, cap)`

#### Escape Analysis

```go
// Check what escapes to heap
// go build -gcflags="-m" main.go

// Stack allocation (fast, no GC)
func process() {
    x := 10        // stays on stack
    y := x + 5     // stays on stack
    fmt.Println(y)
}

// Heap allocation (slower, GC managed)
func createUser() *User {
    u := User{Name: "Alice"}  // escapes to heap (returned)
    return &u
}

// Common escape patterns:
// 1. Returning pointer
// 2. Sending to channel
// 3. Captured by closure
// 4. Stored in interface
// 5. Large struct (stack too small)
```

#### sync.Pool for Object Reuse

```go
// Reduce GC pressure by reusing objects
var bufPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func process(data []byte) string {
    buf := bufPool.Get().(*bytes.Buffer)
    defer func() {
        buf.Reset()
        bufPool.Put(buf)
    }()

    buf.Write(data)
    return buf.String()
}
```

#### String Builder vs Concatenation

```go
// BAD: O(n²) string concatenation
result := ""
for i := 0; i < 10000; i++ {
    result += "a"  // creates new string each time
}

// GOOD: O(n) with strings.Builder
var builder strings.Builder
builder.Grow(10000)  // pre-allocate
for i := 0; i < 10000; i++ {
    builder.WriteString("a")
}
result := builder.String()

// Or use strings.Join
parts := []string{"a", "b", "c"}
result := strings.Join(parts, "")
```

---

### 12.2 Reflection

- [x] `reflect.TypeOf()`, `reflect.ValueOf()`

- [x] Struct tag inspection

- [x] Dynamic method invocation

- [x] Performance cost of reflection

#### Reflection Basics

```go
import "reflect"

// Type and value
x := 42
t := reflect.TypeOf(x)   // int
v := reflect.ValueOf(x)  // 42

// Struct reflection
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

u := User{"Alice", 30}
t = reflect.TypeOf(u)

// Iterate fields
for i := 0; i < t.NumField(); i++ {
    field := t.Field(i)
    fmt.Printf("Field: %s, Type: %s, Tag: %s\n",
        field.Name, field.Type, field.Tag.Get("json"))
}

// Get/set values
v = reflect.ValueOf(&u).Elem()  // need pointer to modify
v.FieldByName("Name").SetString("Bob")
fmt.Println(u.Name)  // "Bob"
```

#### Dynamic Method Invocation

```go
type Greeter interface {
    Greet() string
}

type EnglishGreeter struct{}
func (g EnglishGreeter) Greet() string { return "Hello!" }

type SpanishGreeter struct{}
func (g SpanishGreeter) Greet() string { return "Hola!" }

func callGreet(g interface{}) {
    v := reflect.ValueOf(g)
    method := v.MethodByName("Greet")
    results := method.Call(nil)
    fmt.Println(results[0].String())
}

callGreet(EnglishGreeter{})  // "Hello!"
callGreet(SpanishGreeter{})  // "Hola!"
```

---

### 12.3 Build & Tooling

- [x] `go build`, `go run`, `go install`

- [x] `go vet` (static analysis)

- [x] `golangci-lint` (linter aggregator)

- [x] `go generate` (code generation)

- [x] Build tags (`//go:build ignore`)

- [x] Cross-compilation (`GOOS=linux GOARCH=amd64`)

#### Build Commands

```bash
# Build binary
go build -o myapp .

# Run directly
go run main.go

# Install to $GOPATH/bin
go install github.com/user/cmd/myapp@latest

# Static analysis
go vet ./...

# Linting (install golangci-lint first)
golangci-lint run

# Cross-compile
GOOS=linux GOARCH=amd64 go build -o myapp-linux .
GOOS=windows GOARCH=amd64 go build -o myapp.exe .

# Build with race detector
go build -race .

# Build tags for conditional compilation
//go:build ignore
//go:build linux
//go:build cgo
```

#### go generate

```go
//go:generate stringer -type=Color

type Color int

const (
    Red Color = iota
    Green
    Blue
)

// Run: go generate ./...
// Creates color_string.go with String() method
```

---

### 12.4 Dependency Management

- [x] `go.mod` (module path, Go version)

- [x] `go.sum` (checksums)

- [x] `go mod tidy` (add missing, remove unused)

- [x] `go get`, `go mod download`

- [x] Version pinning

- [x] Private module proxies

#### Module Commands

```bash
# Initialize module
go mod init github.com/user/project

# Add/update dependencies
go get github.com/gin-gonic/gin@v1.9.1
go get -u ./...  # update all

# Clean up
go mod tidy  # add missing, remove unused

# Download all dependencies
go mod download

# Show dependency graph
go mod graph

# Vendor dependencies
go mod vendor

# Verify checksums
go mod verify

# Tidy and vendor (CI/CD)
go mod tidy && go mod vendor
```

---

### 12.5 Docker & Deployment

- [x] Multi-stage Docker builds

- [x] `docker-compose.yaml`

- [x] Scratch vs alpine base images

- [x] Health checks in Docker

- [x] Graceful shutdown (`signal.Notify`)

#### Dockerfile

```dockerfile
# Multi-stage build for Go
FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final image (scratch = empty)
FROM scratch
COPY --from=builder /app/main /main
EXPOSE 8080
CMD ["/main"]

# Or with alpine (has shell, debugging tools)
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main /main
EXPOSE 8080
CMD ["/main"]
```

#### Graceful Shutdown

```go
import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    srv := &http.Server{Addr: ":8080"}

    go func() {
        if err := srv.ListenAndServe(); err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()

    // Wait for interrupt signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Println("Shutting down server...")

    // Give outstanding requests a deadline
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %s", err)
    }

    log.Println("Server exiting")
}

// Or using signal.NotifyContext (Go 1.16+)
func main() {
    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    defer stop()

    srv := &http.Server{Addr: ":8080"}

    go func() {
        if err := srv.ListenAndServe(); err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()

    <-ctx.Done()

    shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(shutdownCtx); err != nil {
        log.Fatalf("Server forced to shutdown: %s", err)
    }
}
```

---

## 13. System Design Concepts for Go

### 13.1 API Design

- [x] REST API conventions

- [x] Status codes (200, 201, 400, 401, 404, 500)

- [x] Request validation

- [x] Pagination (cursor vs offset)

- [x] Rate limiting

- [x] Idempotency

#### REST API Conventions

```go
// Resource-based URLs
GET    /api/v1/users          // list users
GET    /api/v1/users/:id      // get user
POST   /api/v1/users          // create user
PUT    /api/v1/users/:id      // update user
DELETE /api/v1/users/:id      // delete user

// Nested resources
GET    /api/v1/users/:id/orders    // user's orders
POST   /api/v1/users/:id/orders    // create order for user

// Status codes
200 OK                    // success
201 Created               // resource created
204 No Content            // success, no body (DELETE)
400 Bad Request           // invalid input
401 Unauthorized          // authentication required
403 Forbidden             // authorization failed
404 Not Found             // resource not found
409 Conflict              // duplicate resource
422 Unprocessable Entity  // validation error
429 Too Many Requests     // rate limit
500 Internal Server Error // server error
```

#### Pagination

```go
// Offset-based pagination
// GET /api/users?page=2&limit=20

func paginate(w http.ResponseWriter, r *http.Request) {
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

    if page < 1 { page = 1 }
    if limit < 1 || limit > 100 { limit = 20 }

    offset := (page - 1) * limit

    users, total, err := repo.FindAll(offset, limit)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    json.NewEncoder(w).Encode(map[string]interface{}{
        "data":  users,
        "page":  page,
        "limit": limit,
        "total": total,
    })
}

// Cursor-based pagination (better for large datasets)
// GET /api/users?cursor=abc123&limit=20

func cursorPaginate(w http.ResponseWriter, r *http.Request) {
    cursor := r.URL.Query().Get("cursor")
    limit := 20

    users, nextCursor, err := repo.FindByCursor(cursor, limit)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    json.NewEncoder(w).Encode(map[string]interface{}{
        "data":        users,
        "next_cursor": nextCursor,
    })
}
```

#### Rate Limiting

```go
import "golang.org/x/time/rate"

// Per-IP rate limiter
var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

func getVisitor(ip string) *rate.Limiter {
    mu.Lock()
    defer mu.Unlock()

    limiter, exists := visitors[ip]
    if !exists {
        limiter = rate.NewLimiter(1, 3)  // 1 req/sec, burst 3
        visitors[ip] = limiter
    }
    return limiter
}

func rateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        limiter := getVisitor(r.RemoteAddr)
        if !limiter.Allow() {
            http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

---

### 13.2 Authentication

- [x] JWT (JSON Web Tokens)

- [x] OAuth 2.0

- [x] Session-based auth

- [x] Password hashing (bcrypt)

#### JWT Implementation

```go
import (
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

// Generate JWT
func generateToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secretKey))
}

// Validate JWT
func validateToken(tokenStr string) (string, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return []byte(secretKey), nil
    })

    if err != nil || !token.Valid {
        return "", fmt.Errorf("invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", fmt.Errorf("invalid claims")
    }

    userID, ok := claims["user_id"].(string)
    if !ok {
        return "", fmt.Errorf("invalid user_id")
    }

    return userID, nil
}

// Password hashing
func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func checkPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

---

### 13.3 Common Patterns

- [x] Repository pattern (your basicproject uses this)

- [x] Service layer

- [x] Middleware chain

- [x] Graceful shutdown

- [x] Circuit breaker

- [x] Retry with backoff

#### Repository Pattern

```go
// Interface defines data access contract
type UserRepository interface {
    FindByID(ctx context.Context, id string) (*User, error)
    FindByEmail(ctx context.Context, email string) (*User, error)
    Create(ctx context.Context, user *User) error
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id string) error
}

// Concrete implementation
type MongoUserRepo struct {
    coll *mongo.Collection
}

func (r *MongoUserRepo) FindByID(ctx context.Context, id string) (*User, error) {
    var user User
    err := r.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
    return &user, err
}

// Test with mock
type MockUserRepo struct {
    users map[string]*User
}

func (m *MockUserRepo) FindByID(ctx context.Context, id string) (*User, error) {
    user, ok := m.users[id]
    if !ok {
        return nil, fmt.Errorf("not found")
    }
    return user, nil
}
```

#### Service Layer

```go
// Service contains business logic
type UserService struct {
    repo   UserRepository
    mailer Mailer
}

func NewUserService(repo UserRepository, mailer Mailer) *UserService {
    return &UserService{repo: repo, mailer: mailer}
}

func (s *UserService) Register(ctx context.Context, req RegisterRequest) (*User, error) {
    // Validate
    if req.Email == "" {
        return nil, fmt.Errorf("email required")
    }

    // Check duplicate
    existing, _ := s.repo.FindByEmail(ctx, req.Email)
    if existing != nil {
        return nil, fmt.Errorf("email already exists")
    }

    // Hash password
    hash, err := hashPassword(req.Password)
    if err != nil {
        return nil, err
    }

    // Create user
    user := &User{
        Email:    req.Email,
        Password: hash,
        Name:     req.Name,
    }

    if err := s.repo.Create(ctx, user); err != nil {
        return nil, err
    }

    // Send welcome email
    s.mailer.Send(user.Email, "Welcome!", "Thanks for registering")

    return user, nil
}
```

#### Retry with Exponential Backoff

```go
func retryWithBackoff(maxRetries int, baseDelay time.Duration, fn func() error) error {
    var lastErr error

    for i := 0; i <= maxRetries; i++ {
        if err := fn(); err != nil {
            lastErr = err
            if i < maxRetries {
                delay := baseDelay * time.Duration(math.Pow(2, float64(i)))
                time.Sleep(delay)
            }
        } else {
            return nil
        }
    }

    return fmt.Errorf("max retries exceeded: %w", lastErr)
}

// Usage
err := retryWithBackoff(3, time.Second, func() error {
    return db.Ping(ctx)
})
```

---

## 14. Interview Cheat Sheet

### Quick Reference Table

| Topic | Your Coverage | Priority |
| --- | --- | --- |
| Variables & Types | Basics-1 | Must Know |
| Control Flow | Basics-1 | Must Know |
| Functions | Basics-1 | Must Know |
| Error Handling | Basics-1, Basics-4 | Must Know |
| Slices & Maps | Basics-1, Basics-4 | Must Know |
| Structs | Basics-2, Basics-3 | Must Know |
| Pointers | Basics-2 | Must Know |
| Interfaces | Basics-4 | Must Know |
| Goroutines | Basics-5, goheart | Critical |
| Channels | Basics-5, goheart | Critical |
| Select | goheart | Critical |
| Sync Package | goheart | Important |
| Context | basicproject | Important |
| HTTP/JSON | basicHttpServer | Important |
| Project Structure | basicproject | Important |
| Testing | Basics-5 (notes) | Must Know |
| Generics | Basics-5 (notes) | Important |
| MongoDB | basicproject | Project-specific |
| Gin Framework | basicproject | Project-specific |

---

## 15. Top 30 Fresher Interview Questions (Go 2026)

 1. What is Go? Why is it fast?
 2. What are goroutines and how do they differ from threads?
 3. What are channels and how do they work?
 4. What is the difference between buffered and unbuffered channels?
 5. What is `select` statement in Go?
 6. What is `defer` and what is the execution order?
 7. What is the difference between `Mutex` and `Channel`?
 8. What are interfaces? How does implicit satisfaction work?
 9. What is the nil interface problem?
10. What is the difference between `new()` and `make()`?
11. What is the difference between `slice` and `array`?
12. How does map handle concurrency? Is it safe?
13. What is the difference between value receiver and pointer receiver?
14. What is context in Go and why is it important?
15. How do you handle errors in Go?
16. What is error wrapping with `%w`?
17. What are the common concurrency patterns in Go?
18. What is a worker pool pattern?
19. What is fan-out/fan-in?
20. What is a goroutine leak?
21. How do you test in Go?
22. What is a table-driven test?
23. What are generics in Go 1.18+?
24. When to use generics vs interfaces?
25. What is escape analysis?
26. What is the `internal` package in Go?
27. What is `go mod tidy`?
28. How do you implement graceful shutdown?
29. What is the difference between `io.Reader` and `io.Writer`?
30. What is the Go memory model?

---

## 16. Your Learning Roadmap (2026)

```
Week 1-2: Basics (Variables, Control Flow, Functions, Error Handling)
          -> Your Basics-1 covers this

Week 3-4: Data Structures (Slices, Maps, Structs, Pointers)
          -> Your Basics-2, Basics-3 covers this

Week 5-6: Interfaces + Concurrency (Goroutines, Channels, Select)
          -> Your Basics-4, Basics-5 covers this

Week 7-8: HTTP + Web (net/http, JSON, Gin Framework)
          -> Your basicHttpServer covers this

Week 9-10: Project Structure + Database (MongoDB, CRUD)
           -> Your basicproject covers this

Week 11-12: Advanced (Testing, Generics, Sync, Context Patterns)
           -> Your goheart covers this

Week 13-14: System Design + Mock Interviews
           -> Practice with goauth project

Week 15-16: Docker, Deployment, Open Source Contributions
           -> Your docker-compose.yaml is a start
```

---

## 17. Resources for 2026

### Books

- "The Go Programming Language" (Donovan & Kernighan)
- "Concurrency in Go" (Katherine Cox-Buday)
- "Go with Tests" (Ellie Burchard - free online)

### Practice

- [Exercism Go Track](https://exercism.org/tracks/go)
- [LeetCode in Go](https://leetcode.com)
- [Go by Example](https://gobyexample.com)
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests)

### Your Project Progress

```
✅ Basics-1 to Basics-5 (all fundamentals covered)
✅ basicHttpServer (net/http + JSON)
✅ basicproject (full CRUD API with MongoDB + Gin)
✅ goauth (auth project - in progress)
✅ goheart (deep dive concurrency)
✅ ProjectStructure (standard layout)

Next: Add tests to basicproject, implement goauth fully,
      add Docker, add CI/CD pipeline
```

---

*Last updated: 2026-06-13Based on: codebase at C:\\Users\\aspha\\Desktop\\GO\\GO*

---

## 18. Service Layer Pattern (From Your Projects)

**Source:** Booking-Platform/AuthService/services/

#### Service Layer Architecture

```go
// Service layer contains business logic
// Separates HTTP handlers from data access

// Service interface (for testing)
type UserService interface {
    Register(ctx context.Context, req RegisterRequest) (*User, error)
    Login(ctx context.Context, email, password string) (string, error)
    GetByID(ctx context.Context, id string) (*User, error)
    Update(ctx context.Context, id string, req UpdateRequest) error
    Delete(ctx context.Context, id string) error
}

// Concrete implementation
type userService struct {
    repo   UserRepository
    mailer Mailer
    hasher PasswordHasher
}

// Constructor
func NewUserService(repo UserRepository, mailer Mailer, hasher PasswordHasher) UserService {
    return &userService{
        repo:   repo,
        mailer: mailer,
        hasher: hasher,
    }
}

// Business logic methods
func (s *userService) Register(ctx context.Context, req RegisterRequest) (*User, error) {
    // 1. Validate input
    if err := req.Validate(); err != nil {
        return nil, fmt.Errorf("validation: %w", err)
    }

    // 2. Check duplicate
    existing, _ := s.repo.FindByEmail(ctx, req.Email)
    if existing != nil {
        return nil, ErrEmailExists
    }

    // 3. Hash password
    hash, err := s.hasher.Hash(req.Password)
    if err != nil {
        return nil, fmt.Errorf("hash password: %w", err)
    }

    // 4. Create user
    user := &User{
        Email:    req.Email,
        Password: hash,
        Name:     req.Name,
    }

    if err := s.repo.Create(ctx, user); err != nil {
        return nil, fmt.Errorf("create user: %w", err)
    }

    // 5. Send welcome email (async)
    go s.mailer.Send(user.Email, "Welcome!", "Thanks for registering")

    return user, nil
}
```

#### Service with Dependency Injection

```go
// Dependencies injected via constructor
type OrderService struct {
    orderRepo    OrderRepository
    userRepo     UserRepository
    paymentSvc   PaymentService
    inventorySvc InventoryService
    logger       *slog.Logger
}

func NewOrderService(
    orderRepo OrderRepository,
    userRepo UserRepository,
    paymentSvc PaymentService,
    inventorySvc InventoryService,
    logger *slog.Logger,
) *OrderService {
    return &OrderService{
        orderRepo:    orderRepo,
        userRepo:     userRepo,
        paymentSvc:   paymentSvc,
        inventorySvc: inventorySvc,
        logger:       logger,
    }
}

func (s *OrderService) CreateOrder(ctx context.Context, req CreateOrderRequest) (*Order, error) {
    // Business rule: check inventory
    available, err := s.inventorySvc.CheckStock(ctx, req.ProductID, req.Quantity)
    if err != nil {
        return nil, fmt.Errorf("check stock: %w", err)
    }
    if !available {
        return nil, ErrInsufficientStock
    }

    // Business rule: process payment
    payment, err := s.paymentSvc.Charge(ctx, req.UserID, req.Amount)
    if err != nil {
        return nil, fmt.Errorf("payment: %w", err)
    }

    // Create order
    order := &Order{
        UserID:    req.UserID,
        ProductID: req.ProductID,
        Quantity:  req.Quantity,
        PaymentID: payment.ID,
        Status:    "created",
    }

    if err := s.orderRepo.Create(ctx, order); err != nil {
        // Rollback payment
        s.paymentSvc.Refund(ctx, payment.ID)
        return nil, fmt.Errorf("create order: %w", err)
    }

    s.logger.Info("order created",
        "order_id", order.ID,
        "user_id", req.UserID,
    )

    return order, nil
}
```

#### Testing Services

```go
// Mock repository for testing
type MockOrderRepo struct {
    orders map[string]*Order
}

func (m *MockOrderRepo) Create(ctx context.Context, order *Order) error {
    m.orders[order.ID] = order
    return nil
}

// Unit test
func TestCreateOrder(t *testing.T) {
    // Arrange
    mockOrderRepo := &MockOrderRepo{orders: make(map[string]*Order)}
    mockUserRepo := &MockUserRepo{}
    mockPayment := &MockPaymentService{}
    mockInventory := &MockInventoryService{inStock: true}
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    svc := NewOrderService(mockOrderRepo, mockUserRepo, mockPayment, mockInventory, logger)

    req := CreateOrderRequest{
        UserID:    "user123",
        ProductID: "prod456",
        Quantity:  2,
        Amount:    50.00,
    }

    // Act
    order, err := svc.CreateOrder(context.Background(), req)

    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, order)
    assert.Equal(t, "user123", order.UserID)
    assert.Equal(t, 2, order.Quantity)
}
```

---

## 19. DTO Pattern (From Your Projects)

**Source:** Booking-Platform/AuthService/dto/

#### Data Transfer Objects

```go
// DTOs separate API request/response from domain models

// Request DTOs
type RegisterRequest struct {
    Name     string `json:"name" validate:"required,min=2,max=50"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
    Name  string `json:"name,omitempty" validate:"omitempty,min=2,max=50"`
    Email string `json:"email,omitempty" validate:"omitempty,email"`
}

// Response DTOs
type UserResponse struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Role  string `json:"role"`
}

type AuthResponse struct {
    Token string       `json:"token"`
    User  UserResponse `json:"user"`
}

type ErrorResponse struct {
    Error   string `json:"error"`
    Code    string `json:"code"`
    Details any    `json:"details,omitempty"`
}

// Pagination DTOs
type PaginationRequest struct {
    Page  int `json:"page" validate:"min=1"`
    Limit int `json:"limit" validate:"min=1,max=100"`
}

type PaginatedResponse[T any] struct {
    Data       []T   `json:"data"`
    Page       int   `json:"page"`
    Limit      int   `json:"limit"`
    Total      int64 `json:"total"`
    TotalPages int   `json:"total_pages"`
}

// Conversion functions
func UserToResponse(user *User) UserResponse {
    return UserResponse{
        ID:    user.ID.Hex(),
        Name:  user.Name,
        Email: user.Email,
        Role:  user.Role,
    }
}

func UsersToResponse(users []*User) []UserResponse {
    responses := make([]UserResponse, len(users))
    for i, user := range users {
        responses[i] = UserToResponse(user)
    }
    return responses
}
```

#### Validation with DTOs

```go
import "github.com/go-playground/validator/v10"

var validate = validator.New()

func (r RegisterRequest) Validate() error {
    return validate.Struct(r)
}

// Custom validation
func init() {
    validate.RegisterValidation("password_strength", func(fl validator.FieldLevel) bool {
        password := fl.Field().String()
        hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
        hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
        hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
        return hasUpper && hasLower && hasNumber && len(password) >= 8
    })
}

type StrongPasswordRequest struct {
    Password string `json:"password" validate:"required,password_strength"`
}
```

---

## 20. Middleware Pattern (Detailed)

**Source:** Booking-Platform/AuthService/middlewares/

#### HTTP Middleware

```go
// Middleware function type
type Middleware func(http.Handler) http.Handler

// Chain middleware
func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
    for i := len(middlewares) - 1; i >= 0; i-- {
        handler = middlewares[i](handler)
    }
    return handler
}

// Logging middleware
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Wrap response writer to capture status code
        wrapped := &statusResponseWriter{ResponseWriter: w, statusCode: 200}

        next.ServeHTTP(wrapped, r)

        slog.Info("request",
            "method", r.Method,
            "path", r.URL.Path,
            "status", wrapped.statusCode,
            "duration", time.Since(start).String(),
            "remote_addr", r.RemoteAddr,
        )
    })
}

type statusResponseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (w *statusResponseWriter) WriteHeader(code int) {
    w.statusCode = code
    w.ResponseWriter.WriteHeader(code)
}

// Auth middleware
func AuthMiddleware(tokenSecret string) Middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            token := r.Header.Get("Authorization")
            if token == "" {
                http.Error(w, "unauthorized", http.StatusUnauthorized)
                return
            }

            userID, err := validateToken(token, tokenSecret)
            if err != nil {
                http.Error(w, "invalid token", http.StatusUnauthorized)
                return
            }

            ctx := context.WithValue(r.Context(), userIDKey, userID)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

// CORS middleware
func CORSMiddleware(allowedOrigins []string) Middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            origin := r.Header.Get("Origin")
            for _, allowed := range allowedOrigins {
                if origin == allowed {
                    w.Header().Set("Access-Control-Allow-Origin", origin)
                    break
                }
            }
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

            if r.Method == "OPTIONS" {
                w.WriteHeader(http.StatusOK)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}

// Recovery middleware (from gin.Recovery equivalent)
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                slog.Error("panic recovered",
                    "error", err,
                    "path", r.URL.Path,
                )
                http.Error(w, "internal server error", http.StatusInternalServerError)
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

---

## 21. Configuration Management

**Source:** Booking-Platform/AuthService/config/

#### Configuration with Viper

```go
import "github.com/spf13/viper"

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Auth     AuthConfig     `mapstructure:"auth"`
    Logger   LoggerConfig   `mapstructure:"logger"`
}

type ServerConfig struct {
    Host string `mapstructure:"host"`
    Port int    `mapstructure:"port"`
}

type DatabaseConfig struct {
    URI      string `mapstructure:"uri"`
    Name     string `mapstructure:"name"`
    Timeout  int    `mapstructure:"timeout"`
}

type AuthConfig struct {
    JWTSecret     string `mapstructure:"jwt_secret"`
    JWTExpiration int    `mapstructure:"jwt_expiration"`
    BcryptCost    int    `mapstructure:"bcrypt_cost"`
}

type LoggerConfig struct {
    Level  string `mapstructure:"level"`
    Output string `mapstructure:"output"`
}

func LoadConfig(path string) (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(path)
    viper.AddConfigPath(".")

    // Environment variables
    viper.AutomaticEnv()
    viper.SetEnvPrefix("APP")

    // Defaults
    viper.SetDefault("server.host", "localhost")
    viper.SetDefault("server.port", 8080)
    viper.SetDefault("auth.jwt_expiration", 24)
    viper.SetDefault("auth.bcrypt_cost", 12)
    viper.SetDefault("logger.level", "info")

    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("read config: %w", err)
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, fmt.Errorf("unmarshal config: %w", err)
    }

    return &config, nil
}
```

#### Environment Variables

```go
// .env file (use godotenv or viper)
// DB_URI=mongodb://localhost:27017
// JWT_SECRET=my-secret-key
// APP_PORT=8080

// Load with godotenv
import "github.com/joho/godotenv"

func init() {
    if err := godotenv.Load(); err != nil {
        log.Printf("no .env file found")
    }
}

// Access
dbURI := os.Getenv("DB_URI")
jwtSecret := os.Getenv("JWT_SECRET")

// Or with viper
viper.SetEnvPrefix("APP")
viper.BindEnv("server.port", "APP_PORT")
```

---

## 22. Structured Logging

**Source:** Booking-Platform/AuthService/pkg/logger/

#### slog (Standard Library)

```go
import "log/slog"

// Create logger
logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelInfo,
}))

// Or with options
logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level:       slog.LevelDebug,
    AddSource:   true,
    ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
        if a.Key == slog.TimeKey {
            a.Value = slog.StringValue(a.Value.Time().Format(time.RFC3339))
        }
        return a
    },
}))

// Structured logging
logger.Info("server started",
    "host", "localhost",
    "port", 8080,
    "env", "production",
)

logger.Error("database connection failed",
    "error", err,
    "uri", dbURI,
    "attempts", 3,
)

// With context
logger.Info("request handled",
    "method", r.Method,
    "path", r.URL.Path,
    "status", 200,
    "duration", time.Since(start),
    "user_id", userID,
)

// Create child logger with common fields
requestLogger := logger.With(
    "request_id", requestID,
    "user_id", userID,
)

requestLogger.Info("processing order")
requestLogger.Info("order created", "order_id", order.ID)
```

#### zerolog (High Performance)

```go
import "github.com/rs/zerolog"

// Create logger
output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
logger := zerolog.New(output).With().Timestamp().Logger()

// Structured logging
logger.Info().
    Str("host", "localhost").
    Int("port", 8080).
    Msg("server started")

// With error
logger.Error().
    Err(err).
    Str("uri", dbURI).
    Msg("database connection failed")

// Sub-logger with context
requestLogger := logger.With().
    Str("request_id", requestID).
    Str("user_id", userID).
    Logger()

requestLogger.Info().Msg("processing order")
```

---

## 23. Database Migrations

**Source:** sql_sharding_2/migrations/

#### golang-migrate

```go
import (
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigrations(dbURL string) error {
    m, err := migrate.New(
        "file://migrations",
        dbURL,
    )
    if err != nil {
        return fmt.Errorf("create migrate instance: %w", err)
    }
    defer m.Close()

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("run migrations: %w", err)
    }

    return nil
}

// Specific version
m.Migrate(2)  // migrate to version 2

// Rollback
m.Rollback()

// Force version
m.Force(3)  // set version without running migrations
```

#### Migration Files

```sql
-- migrations/001_create_users.up.sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);

-- migrations/001_create_users.down.sql
DROP TABLE IF EXISTS users;
```

---

## 24. Graceful Shutdown (Detailed)

**Source:** Booking-Platform/AuthService/main.go

#### Production Graceful Shutdown

```go
import (
    "context"
    "log/slog"
    "os"
    "os/signal"
    "syscall"
    "time"
)

type Application struct {
    server *http.Server
    db     *mongo.Client
    logger *slog.Logger
}

func (a *Application) Start(ctx context.Context) error {
    // Start server
    go func() {
        a.logger.Info("server starting", "addr", a.server.Addr)
        if err := a.server.ListenAndServe(); err != http.ErrServerClosed {
            a.logger.Error("server failed", "error", err)
            os.Exit(1)
        }
    }()

    // Wait for shutdown signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    sig := <-quit

    a.logger.Info("shutdown signal received", "signal", sig.String())

    // Create shutdown context with timeout
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    // Shutdown HTTP server
    a.logger.Info("shutting down HTTP server")
    if err := a.server.Shutdown(shutdownCtx); err != nil {
        a.logger.Error("HTTP server shutdown failed", "error", err)
    }

    // Disconnect database
    a.logger.Info("disconnecting database")
    if err := a.db.Disconnect(shutdownCtx); err != nil {
        a.logger.Error("database disconnect failed", "error", err)
    }

    a.logger.Info("server stopped gracefully")
    return nil
}

// Usage
func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    config, err := LoadConfig(".")
    if err != nil {
        logger.Error("failed to load config", "error", err)
        os.Exit(1)
    }

    db, err := ConnectDB(config.Database.URI)
    if err != nil {
        logger.Error("failed to connect to database", "error", err)
        os.Exit(1)
    }

    app := &Application{
        server: &http.Server{
            Addr:         fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
            Handler:      setupRouter(db, config),
            ReadTimeout:  15 * time.Second,
            WriteTimeout: 15 * time.Second,
            IdleTimeout:  60 * time.Second,
        },
        db:     db,
        logger: logger,
    }

    if err := app.Start(context.Background()); err != nil {
        logger.Error("application failed", "error", err)
        os.Exit(1)
    }
}
```

---

## 25. Dependency Injection

#### Constructor Injection

```go
// Dependencies injected via constructor
type Handler struct {
    service UserService
    logger  *slog.Logger
}

func NewHandler(service UserService, logger *slog.Logger) *Handler {
    return &Handler{service: service, logger: logger}
}

// Wire up dependencies in main
func main() {
    // Create dependencies
    db := connectDB()
    repo := NewUserRepo(db)
    hasher := NewBcryptHasher(12)
    mailer := NewSMTPMailer(config.SMTP)
    service := NewUserService(repo, mailer, hasher)
    handler := NewHandler(service, logger)

    // Wire up router
    router := setupRouter(handler)
}
```

#### Interface-Based Injection

```go
// Define interfaces where consumed
type UserRepository interface {
    FindByID(ctx context.Context, id string) (*User, error)
    Create(ctx context.Context, user *User) error
}

// Service depends on interface, not implementation
type UserService struct {
    repo UserRepository  // interface
}

// Test with mock
func TestUserService(t *testing.T) {
    mockRepo := &MockUserRepo{}
    svc := NewUserService(mockRepo)
    // test...
}
```

#### Wire (Google's Dependency Injection)

```go
//go:build wireinject

import "github.com/google/wire"

func InitializeApp() (*Application, error) {
    wire.Build(
        provideConfig,
        provideDatabase,
        provideUserRepo,
        provideUserService,
        provideHandler,
        wire.Bind(new(UserRepository), new(*MongoUserRepo)),
        NewApplication,
    )
    return nil, nil
}
```

---

## 26. Resilience Patterns

#### Error Groups (errgroup)

```go
import "golang.org/x/sync/errgroup"

func processInParallel(items []Item) error {
    g, ctx := errgroup.WithContext(context.Background())
    g.SetLimit(10)  // max 10 concurrent

    for _, item := range items {
        item := item
        g.Go(func() error {
            if ctx.Err() != nil {
                return ctx.Err()
            }
            return processItem(ctx, item)
        })
    }

    return g.Wait()
}
```

#### Rate Limiting (Token Bucket)

```go
import "golang.org/x/time/rate"

// Per-client rate limiter
type RateLimiter struct {
    limiters map[string]*rate.Limiter
    mu       sync.RWMutex
}

func NewRateLimiter(rps rate.Limit, burst int) *RateLimiter {
    return &RateLimiter{
        limiters: make(map[string]*rate.Limiter),
    }
}

func (rl *RateLimiter) GetLimiter(key string) *rate.Limiter {
    rl.mu.RLock()
    limiter, exists := rl.limiters[key]
    rl.mu.RUnlock()

    if !exists {
        rl.mu.Lock()
        limiter = rate.NewLimiter(10, 20)  // 10 req/s, burst 20
        rl.limiters[key] = limiter
        rl.mu.Unlock()
    }

    return limiter
}
```

#### Circuit Breaker

```go
type CircuitBreaker struct {
    mu            sync.Mutex
    state         string  // "closed", "open", "half-open"
    failures      int
    successCount  int
    threshold     int
    resetTimeout  time.Duration
    lastFailure   time.Time
}

func NewCircuitBreaker(threshold int, resetTimeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        state:        "closed",
        threshold:    threshold,
        resetTimeout: resetTimeout,
    }
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
    cb.mu.Lock()
    if cb.state == "open" {
        if time.Since(cb.lastFailure) > cb.resetTimeout {
            cb.state = "half-open"
            cb.successCount = 0
        } else {
            cb.mu.Unlock()
            return fmt.Errorf("circuit breaker is open")
        }
    }
    cb.mu.Unlock()

    err := fn()

    cb.mu.Lock()
    defer cb.mu.Unlock()

    if err != nil {
        cb.failures++
        cb.lastFailure = time.Now()
        if cb.failures >= cb.threshold {
            cb.state = "open"
        }
        return err
    }

    if cb.state == "half-open" {
        cb.successCount++
        if cb.successCount >= 3 {
            cb.state = "closed"
            cb.failures = 0
        }
    }

    return nil
}
```

#### Retry with Exponential Backoff

```go
func Retry(maxRetries int, baseDelay time.Duration, fn func() error) error {
    var lastErr error

    for i := 0; i <= maxRetries; i++ {
        if err := fn(); err != nil {
            lastErr = err
            if i < maxRetries {
                delay := baseDelay * time.Duration(math.Pow(2, float64(i)))
                time.Sleep(delay)
            }
        } else {
            return nil
        }
    }

    return fmt.Errorf("max retries exceeded: %w", lastErr)
}

// Usage
err := Retry(3, time.Second, func() error {
    return db.Ping(ctx)
})
```

---

## 27. Caching & Message Queues

#### Redis Caching

```go
import "github.com/redis/go-redis/v9"

type Cache struct {
    client *redis.Client
    ttl    time.Duration
}

func NewCache(addr string, ttl time.Duration) *Cache {
    return &Cache{
        client: redis.NewClient(&redis.Options{Addr: addr}),
        ttl:    ttl,
    }
}

func (c *Cache) Get(ctx context.Context, key string, dest interface{}) error {
    data, err := c.client.Get(ctx, key).Bytes()
    if err == redis.Nil {
        return fmt.Errorf("cache miss")
    }
    if err != nil {
        return fmt.Errorf("redis get: %w", err)
    }
    return json.Unmarshal(data, dest)
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}) error {
    data, err := json.Marshal(value)
    if err != nil {
        return fmt.Errorf("json marshal: %w", err)
    }
    return c.client.Set(ctx, key, data, c.ttl).Err()
}

// Cache-aside pattern
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    // Check cache first
    var user User
    if err := s.cache.Get(ctx, "user:"+id, &user); err == nil {
        return &user, nil
    }

    // Cache miss - fetch from DB
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }

    // Populate cache
    s.cache.Set(ctx, "user:"+id, user)

    return user, nil
}
```

#### Message Queue (NATS)

```go
import "github.com/nats-io/nats.go"

type EventBus struct {
    nc *nats.Conn
}

func NewEventBus(url string) (*EventBus, error) {
    nc, err := nats.Connect(url)
    if err != nil {
        return nil, err
    }
    return &EventBus{nc: nc}, nil
}

func (eb *EventBus) Publish(ctx context.Context, subject string, data interface{}) error {
    payload, err := json.Marshal(data)
    if err != nil {
        return err
    }
    return eb.nc.Publish(subject, payload)
}

func (eb *EventBus) Subscribe(ctx context.Context, subject string, handler func([]byte)) error {
    _, err := eb.nc.Subscribe(subject, func(msg *nats.Msg) {
        handler(msg.Data)
    })
    return err
}

// Usage
eb, _ := NewEventBus("nats://localhost:4222")

// Publish
eb.Publish(ctx, "user.created", UserCreatedEvent{UserID: "123"})

// Subscribe
eb.Subscribe(ctx, "user.created", func(data []byte) {
    var event UserCreatedEvent
    json.Unmarshal(data, &event)
    // handle event...
})
```

---

## 28. gRPC & WebSockets

#### gRPC Basics

```protobuf
// proto/user.proto
syntax = "proto3";

package user;

service UserService {
    rpc GetUser(GetUserRequest) returns (UserResponse);
    rpc CreateUser(CreateUserRequest) returns (UserResponse);
    rpc ListUsers(ListUsersRequest) returns (stream UserResponse);
}

message GetUserRequest {
    string id = 1;
}

message CreateUserRequest {
    string name = 1;
    string email = 2;
}

message UserResponse {
    string id = 1;
    string name = 2;
    string email = 3;
}
```

```go
// Server implementation
type userServer struct {
    pb.UnimplementedUserServiceServer
    svc UserService
}

func (s *userServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
    user, err := s.svc.GetByID(ctx, req.Id)
    if err != nil {
        return nil, status.Errorf(codes.NotFound, "user not found")
    }
    return &pb.UserResponse{
        Id:    user.ID,
        Name:  user.Name,
        Email: user.Email,
    }, nil
}

// Client usage
conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
client := pb.NewUserServiceClient(conn)

user, _ := client.GetUser(ctx, &pb.GetUserRequest{Id: "123"})
```

#### WebSockets

```go
import "github.com/gorilla/websocket"

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("upgrade:", err)
        return
    }
    defer conn.Close()

    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("read:", err)
            break
        }

        // Echo message
        if err := conn.WriteMessage(messageType, message); err != nil {
            log.Println("write:", err)
            break
        }
    }
}
```

---

## 29. Common Go Idioms & Gotchas

#### Range Variable Capture

```go
// BAD: loop variable shared (pre-Go 1.22)
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // all print 5!
    }()
}

// GOOD: Go 1.22+ (loop variable per-iteration)
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // prints 0,1,2,3,4
    }()
}

// GOOD: pre-1.22, pass as parameter
for i := 0; i < 5; i++ {
    go func(n int) {
        fmt.Println(n)  // prints 0,1,2,3,4
    }(i)
}
```

#### Defer Arguments Evaluated Immediately

```go
x := 10
defer fmt.Println(x)  // prints 10, not 11
x++

// Named returns can be modified by defer
func double(x int) (result int) {
    defer func() { result *= 2 }()
    return x  // returns x, defer doubles it
}
// double(5) returns 10
```

#### Nil Interface Gotcha

```go
var p *MyStruct = nil
var i interface{} = p

// i is NOT nil!
// i has type *MyStruct, value nil
if i != nil {
    fmt.Println("i is not nil")  // this prints!
}

// Solution: return untyped nil
func getError() error {
    var err *MyStruct = nil
    if err == nil {
        return nil  // return untyped nil
    }
    return err
}
```

#### String Conversion Gotcha

```go
// This gives Unicode code point, NOT string representation
s := string(65)  // "A", not "65"

// To convert number to string representation
s := strconv.Itoa(65)    // "65"
s := fmt.Sprintf("%d", 65)  // "65"
```

#### Map Iteration Order

```go
m := map[int]string{1: "a", 2: "b", 3: "c"}
// Iteration order is RANDOM
// Don't rely on order!

// If you need order, sort keys first
keys := make([]int, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Ints(keys)
for _, k := range keys {
    fmt.Println(k, m[k])
}
```

#### Nil Slice vs Empty Slice

```go
var s []int         // nil slice
s := []int{}        // empty slice

// JSON marshaling
json.Marshal(nil)    // null
json.Marshal([]int{}) // []

// Best practice: return empty, not nil
func getItems() []int {
    return []int{}  // better than nil
}
```

#### Pointer Receiver Consistency

```go
// If any method needs pointer receiver, ALL methods should use pointer receiver
type User struct { Name string }

// GOOD: all pointer receivers
func (u *User) GetName() string { return u.Name }
func (u *User) SetName(name string) { u.Name = name }

// BAD: mixed receivers (confusing)
func (u User) GetName() string { return u.Name }
func (u *User) SetName(name string) { u.Name = name }
```

#### Error Handling Patterns

```go
// Always check errors
result, err := doSomething()
if err != nil {
    return fmt.Errorf("doSomething: %w", err)
}

// Don't ignore errors
result, _ := doSomething()  // BAD!

// Use errors.Is/errors.As, not ==
if errors.Is(err, sql.ErrNoRows) {  // GOOD
if err == sql.ErrNoRows {  // BAD (doesn't work with wrapping)
```

#### init() Function Gotchas

```go
// init() runs before main, can cause issues
func init() {
    // Runs at package initialization
    // Hard to test
    // Execution order unclear
    // Prefer explicit initialization
}

// Better: explicit initialization
func Init() error {
    // Can return error
    // Can be called explicitly
    // Testable
}
```

---

## 30. Missing Topics Checklist

### Essential Topics (Must Know)

| Topic | Section | Status |
| --- | --- | --- |
| Zero values | 1.1 | ✅ Complete |
| Type conversion vs assertion | 1.1 | ✅ Complete |
| defer in depth | 1.3 | ✅ Complete |
| errors.As | 1.4 | ✅ Complete |
| Slice internals | 2.1 | ✅ Complete |
| Map concurrency | 2.2 | ✅ Complete |
| Embedded structs | 2.3 | ✅ Complete |
| new() vs make() | 3 | ✅ Complete |
| Escape analysis | 3 | ✅ Complete |
| Nil interface trap | 4 | ✅ Complete |
| Goroutine scheduling | 5.1 | ✅ Complete |
| Channel patterns | 5.2 | ✅ Complete |
| sync.Once, sync.Map | 5.4 | ✅ Complete |
| errgroup | 5.5 | ✅ Complete |
| Context patterns | 6 | ✅ Complete |
| Table-driven tests | 10.1 | ✅ Complete |
| Benchmarking | 10.2 | ✅ Complete |
| Interface mocking | 10.3 | ✅ Complete |
| Generics constraints | 11 | ✅ Complete |

### Project Patterns (From Your Repos)

| Pattern | Source | Section |
| --- | --- | --- |
| Service layer | AuthService | 18 |
| DTO pattern | AuthService | 19 |
| Middleware | AuthService | 20 |
| Config management | AuthService | 21 |
| Structured logging | AuthService | 22 |
| DB migrations | sql_sharding_2 | 23 |
| Graceful shutdown | AuthService | 24 |
| Dependency injection | Both | 25 |
| Resilience patterns | General | 26 |
| Caching | General | 27 |
| gRPC/WebSockets | General | 28 |
| Go idioms/gotchas | General | 29 |
