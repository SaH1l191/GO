package main  

import (
	"fmt"
	"time"
)
//A goroutine is a lightweight function running concurrently.


//fork join process : main goroutine starts other goroutines and waits for them to finish before exiting.
// func sayHello() {
// 	fmt.Println("Hello from goroutine")
// }

// func main() {
// 	go sayHello()
// 	time.Sleep(time.Second)
// }


//channels : used for communication between goroutines. They allow you to send and receive values of a specified type.

// func worker(ch chan int) {
// 	ch <- 42 // send value to channel
// }
// func main() {
// 	ch := make(chan int)
// 	go worker(ch)
// 	value := <-ch // receive value from channel // blocks until value is available
// 	fmt.Println(value)
// }


// Unbuffered Channels
// Default channels are unbuffered.
// Meaning:
// Sender waits
// Receiver waits
// Both must be ready.


//Buffered Channels
// Buffered channels store values temporarily.
//The value moves immediately when both sides are ready.it doesn't store anything.like a pipe
// ch := make(chan int, 2) //	buffer size of 2
// ch <- 1
// ch <- 2
// fmt.Println(<-ch)
// fmt.Println(<-ch)

// ch := make(chan int)
// ch <- 5 // blocks forever (no receiver)
// fatal error: all goroutines are asleep - deadlock

// closing channels :
// Use close() to signal no more values will be sent in a channel.

// ch := make(chan int)
// go func() {
// 	ch <- 1
// 	ch <- 2
// 	close(ch)
// }()


//ranging over channels :
// ch := make(chan int)
// go func() {
// 	for i := 1; i <= 3; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }()
// for val := range ch {
// 	fmt.Println(val)
// }
//when ch <- i then the value is sent to channel and the loop continues until the channel is closed. Once closed, the loop exits.
//execution order -> go func for 1 -> for val :- range ch(1 size) prints , then go func for 2 & so no ...


//select waits on multiple channel operations.
// select {
// case msg := <-ch1:
// 	fmt.Println(msg)
// case msg := <-ch2:
// 	fmt.Println(msg)
// }
// Whichever channel is ready executes.


// Timeout with select
// select {
// case msg := <-ch:
// 	fmt.Println(msg)

// case <-time.After(2 * time.Second):
// 	fmt.Println("Timeout")
// }If channel doesn't respond in 2 seconds, timeout triggers.


//Directional Channels
// Channels can be restricted to send-only or receive-only.

// Send-only
// func sendData(ch chan<- int) {
// 	ch <- 10
// }

// Receive-only
// func receiveData(ch <-chan int) {
// 	val := <-ch
// 	fmt.Println(val)
// }


// real concurrent programming :
// func producer(ch chan int) {
// 	for i := 1; i <= 5; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }

// func consumer(ch <-chan int) {
// 	for val := range ch {
// 		fmt.Println("Received:", val)
// 	}
// }

// func main() {
// 	ch := make(chan int)

// 	go producer(ch)
// 	consumer(ch)
// }



//most common packages ; 
// fmt.Print()	Print without formatting
// fmt.Println()	Print with newline
// fmt.Printf()	Formatted printing
// fmt.Sprintf()	Format string and return it
// fmt.Errorf()	Create formatted error


// package main
// import "fmt"

// func main() {
//     name := "Rahul"
//     age := 25

//     fmt.Println("Hello", name)

//     fmt.Printf("Name: %s Age: %d\n", name, age)

//     msg := fmt.Sprintf("User %s is %d years old", name, age)
//     fmt.Println(msg)
// }


// os and io — Files & Input/Output
// data, err := os.ReadFile("file.txt")
// os.WriteFile("test.txt", []byte("Hello"), 0644)

// io : Defines interfaces for data streams.
// io.Copy(dst, src)


// net/http — HTTP Server & Client
// This package lets you build web servers and APIs.


//encoding/json :Used to convert Go structs ↔ JSON.
// Marshal (Go → JSON)
// type User struct {
//     Name string
//     Age  int
// }
// user := User{"Rahul", 25}
// data, _ := json.Marshal(user)
// fmt.Println(string(data))
// Output:
// {"Name":"Rahul","Age":25}


// Unmarshal (JSON → Go)
// jsonData := `{"Name":"Rahul","Age":25}`
// var u User
// json.Unmarshal([]byte(jsonData), &u)
// fmt.Println(u.Name)


//remaining : 
// sync package
// sync.Mutex and sync.RWMutex
// sync.WaitGroup
// sync.Once
// sync.Map
// sync.Cond (less common but exists)
// Common patterns

// Fan-out / fan-in
// Worker pools
// Done channel for cancellation (before context)
// Pipeline pattern

// Race conditions

// What a data race is
// go test -race and go run -race
// Why maps are not goroutine-safe

// 10. Testing

// _test.go files, TestXxx(t *testing.T) naming
// t.Errorf vs t.Fatalf — difference matters
// Table-driven tests — the standard Go pattern
// go test ./..., -v, -run flags
// Subtests: t.Run("name", func(t *testing.T))
// Benchmarks: BenchmarkXxx(b *testing.B)
// testify library (common in real projects)
// Mocking via interfaces — Go's way of doing dependency injection for tests


// 11. Generics (Go 1.18+)

// Type parameters: func Map[T, U any](s []T, f func(T) U) []U
// Type constraints: comparable, any, custom constraint interfaces
// When to use generics vs interfaces
// Common generic utility patterns








