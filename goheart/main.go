package main
import ( 
	"fmt"
	// "sync"
	"time"
)






//article : https://medium.com/@bRutTe_forc3/concurrency-in-go-a-deep-dive-into-goroutines-channels-waitgroups-and-more-b75475290847


//eg -1 

//communicate go routines by using shared memory and mutexes to synchronize access to that memory
// var mu sync.Mutex
// var availableSeat = 2

//using sync.WaitGroup to wait for all goroutines to finish before exiting the main function
// it has .Add  := to set the number of goroutines to wait for
// .Done := to signal that a goroutine has finished
// .Wait := to block until all goroutines have finished
//to keep the main func alive and wait for all routines to finish before exiting the program
// func bookSeat(user string,wg *sync.WaitGroup){
// 	defer wg.Done()
// 	mu.Lock()
// 	defer mu.Unlock()
// 	if availableSeat > 0 {
// 		fmt.Printf("Seat booked for %s\n", user)
// 		availableSeat--
// 	} else {
// 		fmt.Printf("No seats available for %s\n", user)
// 	}
// }

// func main () { 
// 	var wg sync.WaitGroup
// 	users := []string{"Alice", "Bob", "Charlie", "Dave"}
// 	for _ , user := range users {
// 		wg.Add(1)
// 		go bookSeat(user, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Printf("Final available seats %d\n",availableSeat)
// }



//eg -2 


//sharing data by communciating through channels instead of shared memory


// A channel is:
// A pipe between goroutines to send & receive data
/// channel operations 
// Send
// ch <- "hello"
// Receive
// msg := <-ch
// Close
// close(ch)



// func main () { 
// 	ch := make(chan string)
// 	go func(){
// 		ch <- "Hello from goroutine"
// 	}()
// 	msg := <-ch
// 	fmt.Println(msg)
// }
//steps
// Channel created
// Goroutine starts
// Goroutine tries ch <- "hello"
// Main is waiting <-ch
// They meet → data transferred
// Print happens
//  This is synchronization



//unbuffered channel eg : 
// Default channel = unbuffered
//Send and receive must happen at the same time

//deadlock 
//forever blocking => ch <- "hello"  // no one is receiving
//forever blocking => msg := <-ch  // no one is sending


//multiple goroutines sending and receiving on the same channel
// channel as a signal = >  
// done := make(chan bool)
//   done <- struct{}{} // signal completion


// for range orders {
//     <-done
// }
//  Meaning:
// “Wait for all signals”


// meaning of : 
// struct{}{}
// Empty value
// Zero memory
// Only used for signaling



//directional channels
// func producer(ch chan<- string) {} // send only 
//func consumer(ch <-chan string) {} // receive only

// closing channels
// close(ch) // signals no more values will be sent



//buffered channels
// ch := make(chan string, 2) // buffer size of 2
// can stores values without receiver , block s only when full 


// eg : - 

// func producer(ch chan<- string){
// 	for i := 0; i < 5; i++ {
// 		ch <- fmt.Sprintf("Producer-Message %d", i)
// 	}
// 	close(ch) // signal no more messages
// }
// func consumer(ch <-chan string){
// 	for msg := range ch { //range over channel until closed
// 		fmt.Println("Received:", msg)
// 	}
// }
// func main () { 
// 	ch := make(chan string)
// 	go producer(ch)
// 	consumer(ch) //blocking call to consume messages until channel is closed
// }

// for msg := range ch   
// is doign the following :
// for {
//     msg, ok := <-ch
//     if !ok {
//         break
//     }
//     fmt.Println(msg)
// }
//blocking operation



// Consumer:   wait ── receive ── print ── wait ── receive ── print
// Producer:        send ───────────── send ───────────── send
//  They sync at each step



//eg -3 

// func producer(ch chan<- string){
// 	for i := 0; i < 5; i++ {
// 		ch <- fmt.Sprintf("Producer-Message %d", i)
// 	}
// 	close(ch) // signal no more messages
// }
// func consumer(ch <-chan string){
// 	for { //block until channel is closed
// 		val, ok := <-ch
// 		if !ok {
// 			break // channel closed, exit loop
// 		}
// 		fmt.Println("Received:", val)
// 	}
// } 
// func main() { 
//  	ch := make(chan string)
// 	go producer(ch)
// 	consumer(ch)
// }



//select statement

// select is like a switch statement for channels. 
// It lets you wait on multiple channel operations at the same time 
// and proceed with whichever one is ready first.

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch1 <- "Message from channel 1"
// 	}()
// 	go func(){
// 		time.Sleep(1 * time.Second)
// 		ch2 <- "Message from channel 2"
// 	}()
// 	for i := 0; i < 2; i++ {
// 		select { // one who wins first will be printed 
// 		case msg1 := <-ch1:
// 			fmt.Println("Received:", msg1)
// 		case msg2 := <-ch2:
// 			fmt.Println("Received:", msg2)
// 		}
// 	}
// } 

//Buffered channels can hold that many values without blocking the sender.
//fan out (one input -> many workers )
//						receive only    send only channel
func worker(id int,jobs <- chan int , results chan <- int ){
	for j:= range jobs { //If the jobs channel is empty, the worker pauses and waits for a value. This is called blocking on receive.
		fmt.Printf("Worker %d processing job %d\n",id,j)
		time.Sleep(500 * time.Millisecond)//suimulate real work 
		results <- j * 2
		//Once a job arrives, it “wakes up,” processes it, and sends the result to the results channel.
		//So a worker does not finish immediately. It keeps listening to jobs until the channel is closed.
	}
}


//only sender closes 
func main(){ 
	const numWorkers = 3
	const numJobs = 5

	jobs := make(chan int,numWorkers)
	results := make(chan int, numJobs)
	
	//assinging workers to job channel to listen upon  
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}
	
	// send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close (jobs) // no more jobs will be sent

	// ... wait for results
	for a := 1; a <= numJobs; a++ {
		fmt.Println("Result:", <-results)
	}
}
//steps ; 
// Start main program.
// Create channels jobs and results.
// Launch 3 workers concurrently:
// Worker 1, Worker 2, Worker 3.
// Push jobs 1–5 into jobs channel.
// Workers pull jobs as soon as they are available:
// Worker 1 might get job 1
// Worker 2 might get job 2
// Worker 3 might get job 3
// Worker 1 might pull job 4 (after finishing job 1), etc.
// Each worker processes job for 500ms and sends j*2 to results.
// Main routine collects results as they come in (order may vary).


//flow of code : 
// Step 0: The setup

// You have:

// jobs channel (workers receive from this)
// results channel (workers send to this)
// worker() function, which loops over jobs channel

// And you launch goroutines:

// for w := 1; w <= numWorkers; w++ {
//     go worker(w, jobs, results)
// }
// Step 1: What a goroutine is
// A goroutine is like a lightweight thread.
// When you call go worker(...), it starts running concurrently with the main function.
// The main function does not wait for the goroutine to finish unless you explicitly coordinate (e.g., with channels, WaitGroup, etc.).
// So launching go worker(...) is like saying: “Run this function in the background while I continue my own work.”
// Step 2: How the worker behaves
// for j := range jobs {
//     fmt.Printf("Worker %d processing job %d\n", id, j)
//     time.Sleep(500 * time.Millisecond)
//     results <- j * 2
// }
// for j := range jobs → this loops over incoming jobs.
// If the jobs channel is empty, the worker pauses and waits for a value. This is called blocking on receive.
// Once a job arrives, it “wakes up,” processes it, and sends the result to the results channel.

// So a worker does not finish immediately. It keeps listening to jobs until the channel is closed.

// Step 3: How channels coordinate things

// Channels are synchronization points between goroutines.

// When a goroutine tries to send to a channel:
// If the channel is full, it blocks until there is space.
// If the channel has space (buffered), it can send immediately.
// When a goroutine tries to receive from a channel:
// If the channel is empty, it blocks until a value arrives.
// If there is a value, it receives immediately.

// This is how Go ensures safe communication without explicit locks.

// Step 4: How your worker pool flows

// Let’s walk through an example:

// Main goroutine creates jobs and results channels.

// Main goroutine launches 3 worker goroutines:

// Worker 1: waiting for job
// Worker 2: waiting for job
// Worker 3: waiting for job

// Main goroutine sends 5 jobs into jobs channel:

// jobs <- 1
// jobs <- 2
// jobs <- 3
// jobs <- 4
// jobs <- 5

// Workers pick jobs concurrently (one job per worker at a time):

// Worker 1: job 1
// Worker 2: job 2
// Worker 3: job 3

// Workers process the jobs (sleep 500ms) and then send results:

// results <- 2
// results <- 4
// results <- 6
// Remaining jobs 4 and 5 are picked by the next available worker.

// Main goroutine receives results from results channel:

// r := <-results  // blocks until a result is available

// ✅ Key insight: goroutines coordinate through channels, not by directly “waiting” or “locking” each other.

// Step 5: How the program ends
// Worker loops end only when jobs channel is closed.

// Main goroutine can close(jobs) after sending all jobs:

// close(jobs)
// This signals workers: “No more jobs coming.”
// Then the for j := range jobs loop in each worker finishes, and the worker goroutine exits.
// Step 6: Summary of flow
// Main goroutine: prepares channels, launches workers, sends jobs, collects results.
// Worker goroutines: wait for jobs, process, send results.
// Channels: act as both a queue and a synchronization tool.
// Program finishes when all jobs are processed, all results are collected, and channels are closed.


//simple eg ;
// package main
// import "fmt"


// func processOrders(id int,done chan struct{}){
//   fmt.Printf("Processing order #%d\n", id)
//   done <- struct{}{} ; //doublt struct means sfirst {} used wiht struct and second {} means empty value of struct type //empty struct means signal 0 stop 
// }

// func main() {
// 	done := make(chan struct{})
// 	orders := []int{101,102,103}
// 	for _,id := range orders{
// 	  go processOrders(id,done)
// 	}
// 	for range orders {
// 	  <-done 
// 	}
// 	fmt.Println("all processed")
// }


//eg of 
// Producer-Consumer Example in Go
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func producer(ch chan<- int, count int) {
// 	for i := 1; i <= count; i++ {
// 		fmt.Println("Producing:", i)
// 		ch <- i // send item to channel
// 		time.Sleep(200 * time.Millisecond) // simulate work
// 	}
// 	close(ch) // signal no more items
// }

// func consumer(ch <-chan int) {
// 	for item := range ch { // receive items until channel is closed
// 		fmt.Println("Consuming:", item)
// 		time.Sleep(500 * time.Millisecond) // simulate processing
// 	}
// }

// func main() {
// 	ch := make(chan int, 5) // buffered channel

// 	// Start producer and consumer as goroutines
// 	go producer(ch, 10)
// 	go consumer(ch)

// 	// Wait for the consumer to finish (simple sleep)
// 	time.Sleep(6 * time.Second)
// 	fmt.Println("All done!")
// }
// How it works
// Channels:
// ch is a buffered channel of integers.
// Producer sends values to ch.
// Consumer receives values from ch.
// Producer:
// Sends integers 1 to 10 into the channel.
// Closes the channel when done to signal the consumer no more data is coming.
// Consumer:
// Loops over the channel using for item := range ch.
// Automatically exits when the channel is closed.
// Goroutines:
// Both producer and consumer run concurrently with go.
// The main goroutine sleeps to let them finish (simple way; more robust would be sync.WaitGroup).




//eg  := 

// with done context + for select 
// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// // Producer sends numbers into ch until context is canceled
// func producer(ctx context.Context, ch chan<- int) {
// 	i := 1
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Producer stopping")
// 			close(ch) // close channel to signal consumer
// 			return
// 		case ch <- i:
// 			fmt.Println("Producing:", i)
// 			i++
// 			time.Sleep(200 * time.Millisecond)
// 		}
// 	}
// }

// // Consumer receives numbers from ch until channel is closed or context is canceled
// func consumer(ctx context.Context, ch <-chan int) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Consumer stopping")
// 			return
// 		case item, ok := <-ch:
// 			if !ok {
// 				fmt.Println("Channel closed, consumer exiting")
// 				return
// 			}
// 			fmt.Println("Consuming:", item)
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 	}
// }

// func main() {
// 	ch := make(chan int, 5)
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	go producer(ctx, ch)
// 	go consumer(ctx, ch)

// 	// Wait enough for context to expire
// 	time.Sleep(4 * time.Second)
// 	fmt.Println("Main finished")
// }






// start from below: 
//fan in remaining 
// Pipeline pattern

// package main

// import (
// 	"fmt"
// 	"time"
// )

// // Producer sends numbers into its own channel
// func producer(id int, ch chan<- int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- id*10 + i // produce unique numbers per producer
// 		time.Sleep(100 * time.Millisecond)
// 	}
// 	close(ch)
// }

// // Fan-in merges multiple channels into a single channel
// func fanIn(ch1, ch2 <-chan int) <-chan int {
// 	out := make(chan int)

// 	go func() {
// 		for ch1 != nil || ch2 != nil {
// 			select {
// 			case v, ok := <-ch1:
// 				if !ok {
// 					ch1 = nil // stop reading from ch1
// 					continue
// 				}
// 				out <- v
// 			case v, ok := <-ch2:
// 				if !ok {
// 					ch2 = nil // stop reading from ch2
// 					continue
// 				}
// 				out <- v
// 			}
// 		}
// 		close(out)
// 	}()

// 	return out
// }

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go producer(1, ch1)
// 	go producer(2, ch2)

// 	// Merge both channels
// 	merged := fanIn(ch1, ch2)

// 	// Consume from the merged channel
// 	for v := range merged {
// 		fmt.Println("Received:", v)
// 	}

// 	fmt.Println("Done!")
// }


//Example 2: what happens when you exceed the buffer
// package main

// func main() {
//     ch := make(chan int, 2)
//     ch <- 1 // fine
//     ch <- 2 // fine, buffer full
//     ch <- 3 // BLOCKS - buffer full, no receiver anywhere
// }
// Output:

// fatal error: all goroutines are asleep - deadlock!

//Buffer full + no one receiving = sender blocks forever. With only main running and no goroutine to drain the channel, the program deadlocks.




eg : 
// package main

// import (
//     "fmt"
//     "sync"
// )
// func main() {
//     taskQueue := make(chan string, 4)
//     var wg sync.WaitGroup

//     // Queue up tasks and use WaitGroup to wait until all tasks are enqueued
//     tasks := []string{"resize image", "send invoice", "generate report", "backup database"}

//     for _, task := range tasks {
//         wg.Add(1)  // Increment the WaitGroup counter
//         go func(t string) {
//             defer wg.Done()  // Decrement the counter when the task is done
//             taskQueue <- t
//         }(task)
//     }

//     // Use a goroutine to close the channel after all tasks are queued
//     go func() {
//         wg.Wait()  // Wait until all tasks are enqueued
//         close(taskQueue)  // Now it's safe to close the channel
//     }()

//     // Use select with val, ok to process tasks
//     for {
//         select {
//         case task, ok := <-taskQueue:
//             if !ok {
//                 // Channel is closed and all tasks have been processed
//                 fmt.Println("No more tasks to process.")
//                 return
//             }
//             // Process the task
//             fmt.Println("Processing:", task)
//         }
//     }
// }
explanation : 
// Here's the flow of how synchronization occurs:

// Main Goroutine: Enqueues tasks into the channel (in either the main loop or within separate goroutines).
// sync.WaitGroup: The main goroutine (or any other goroutine responsible for enqueuing tasks) increments the wait group counter before sending the task and calls wg.Done() after sending each task.
// Goroutine to Close the Channel:
// The goroutine that calls wg.Wait() will block until the counter reaches zero (i.e., all tasks have been enqueued).
// Once the counter reaches zero, the goroutine safely closes the channel, ensuring that the channel isn’t closed prematurely.
// select Loop:
// In the main goroutine, the select loop keeps running, processing tasks from the channel.
// Once the channel is closed (detected by the ok variable becoming false), the loop exits gracefully.
// Flow in Detail
// Task Enqueuing:
// Each task is enqueued in a separate goroutine (go func() { taskQueue <- task; wg.Done() }), and each goroutine uses wg.Done() to signal that the task has been sent.
// Waiting for All Tasks to be Enqueued:
// In the main goroutine, wg.Wait() blocks until all tasks are enqueued. It ensures that the channel isn't closed until every task has been queued.
// Closing the Channel:
// After wg.Wait() (meaning all tasks are queued), the goroutine responsible for closing the channel calls close(taskQueue).
// Processing Tasks in the select Loop:
// The select loop waits for values from the taskQueue channel.
// When ok == false (because the channel is closed), the select loop terminates, and the program exits.





eg : = 
Select
select is like switch but for channels. It waits on multiple channels simultaneously and runs whichever case becomes ready first.

Example 1: first response wins

You have two data sources. You want whichever responds first:

package main

import (
    "fmt"
    "time"
)
func fetchFromCache(ch chan string) {
    time.Sleep(500 * time.Millisecond)
    ch <- "data from cache"
}
func fetchFromDB(ch chan string) {
    time.Sleep(2 * time.Second)
    ch <- "data from database"
}
func main() {
    cache := make(chan string)
    db := make(chan string)
    go fetchFromCache(cache)
    go fetchFromDB(db)
    select {
    case result := <-cache:
        fmt.Println("Got:", result)
    case result := <-db:
        fmt.Println("Got:", result)
    }
}
Output:

Got: data from cache
Cache responds in 500ms, DB in 2s. select picks the cache. If both were ready at the same time, Go picks randomly.

How select behaves:

One case ready → runs it
Multiple ready → picks randomly
None ready → blocks until one is ready
Example 2: request timeout

package main

import (
    "fmt"
    "time"
)
func fetchUserProfile(userID int, ch chan string) {
    time.Sleep(3 * time.Second) // slow external API
    ch <- fmt.Sprintf("profile for user %d", userID)
}
func main() {
    ch := make(chan string)
    go fetchUserProfile(42, ch)
    select {
    case profile := <-ch:
        fmt.Println("Got:", profile)
    case <-time.After(2 * time.Second):
        fmt.Println("Request timed out - returning cached version")
    }
}
Output:

Request timed out - returning cached version
time.After(d) returns a channel that receives after duration d. Timeout logic with no external libraries.




eg : 
Example 3: event loop

Real-world programs wrap select in a loop to keep listening:

package main

import (
    "fmt"
    "time"
)
func main() {
    orders := make(chan string)
    cancels := make(chan string)
    shutdown := make(chan struct{})
    go func() {
        time.Sleep(500 * time.Millisecond)
        orders <- "order #1001"
    }()
    go func() {
        time.Sleep(1 * time.Second)
        cancels <- "order #999"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        close(shutdown)
    }()
    for {
        select {
        case o := <-orders:
            fmt.Println("New order:", o)
        case c := <-cancels:
            fmt.Println("Cancelled:", c)
        case <-shutdown:
            fmt.Println("Shutting down order processor")
            return
        }
    }
}
Output:

New order: order #1001
Cancelled: order #999
Shutting down order processor


Default Selection
Add default and select becomes non-blocking. If no channel is ready, it runs default immediately.

Example 1: non-blocking check

package main

import "fmt"
func main() {
    notifications := make(chan string, 1)
    // check if there's a notification without waiting
    select {
    case msg := <-notifications:
        fmt.Println("Notification:", msg)
    default:
        fmt.Println("No notifications right now")
    }
}
Output:

No notifications right now
Example 2: try to send, skip if full

select {
case alertChannel <- "disk usage > 90%":
    fmt.Println("Alert sent")
default:
    fmt.Println("Alert channel busy, dropping alert")
}
Instead of blocking when the channel is full, you drop the alert and move on. Useful for fire-and-forget notifications where blocking is worse than occasionally missing one.





worker pool
Worker Pool Pattern
This is the most common production concurrency pattern in Go. Once you understand it, you’ll see it everywhere.

The problem: you need to resize 50,000 images that users have uploaded. Creating one goroutine per image means 50,000 goroutines at 2KB each, that’s 100MB just for goroutine stacks, plus the scheduler going haywire.

The solution: spin up a fixed number of worker goroutines. They all pull from a shared jobs channel. As soon as a worker finishes one image, it picks up the next. Job count can be millions; goroutine count stays fixed at numWorkers.

Press enter or click to view image in full size

image jobs channel → [Worker 1] →
                     [Worker 2] → results channel
                     [Worker 3] →
Example:

package main

import (
    "fmt"
    "sync"
    "time"
)
type ImageJob struct {
    filename string
    width    int
    height   int
}
type ImageResult struct {
    filename string
    status   string
}
func resizeWorker(id int, jobs <-chan ImageJob, results chan<- ImageResult, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d: resizing %s to %dx%d\n", id, job.filename, job.width, job.height)
        time.Sleep(500 * time.Millisecond) // simulate resize operation
        results <- ImageResult{
            filename: job.filename,
            status:   fmt.Sprintf("resized to %dx%d", job.width, job.height),
        }
    }
}
func main() {
    jobs := []ImageJob{
        {"avatar1.jpg", 128, 128},
        {"banner.png", 1920, 1080},
        {"thumbnail.jpg", 300, 300},
        {"cover.png", 800, 600},
        {"icon.jpg", 64, 64},
        {"hero.png", 1200, 400},
    }
    jobCh := make(chan ImageJob, len(jobs))
    resultCh := make(chan ImageResult, len(jobs))
    var wg sync.WaitGroup
    // Start 3 workers
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go resizeWorker(w, jobCh, resultCh, &wg)
    }
    // Send all jobs
    for _, job := range jobs {
        jobCh <- job
    }
    close(jobCh) // no more jobs
    // Close results once all workers are done
    go func() {
        wg.Wait()
        close(resultCh)
    }()
    // Collect results
    for result := range resultCh {
        fmt.Printf("Done: %s - %s\n", result.filename, result.status)
    }
    fmt.Println("All images processed")
}
Output (order varies):

Worker 1: resizing avatar1.jpg to 128x128
Worker 2: resizing banner.png to 1920x1080
Worker 3: resizing thumbnail.jpg to 300x300
Worker 1: resizing cover.png to 800x600
Worker 2: resizing icon.jpg to 64x64
Worker 3: resizing hero.png to 1200x400
Done: avatar1.jpg - resized to 128x128
Done: banner.png - resized to 1920x1080
Done: thumbnail.jpg - resized to 300x300
Done: cover.png - resized to 800x600
Done: icon.jpg - resized to 64x64
Done: hero.png - resized to 1200x400
All images processed




Big eg : 
Fan-Out / Fan-In Pattern
Worker Pool distributes jobs across workers. Fan-Out / Fan-In is a related but distinct idea:

Press enter or click to view image in full size

Fan-Out: one goroutine produces data, multiple goroutines consume it in parallel.

Fan-In: multiple goroutines produce data, one goroutine merges everything into a single stream.

A real scenario: you’re building a fraud detection system. Every transaction comes in through one stream. You want to run multiple independent checks on it simultaneously velocity check, location check, blacklist check then merge all their verdicts into one place for a final decision.

→ [velocity checker]  →
transaction stream   → [location checker]  → merged verdicts
                     → [blacklist checker] →

Example: parallel fraud checks

package main

import (
    "fmt"
    "sync"
    "time"
)
type Transaction struct {
    ID     int
    Amount float64
    UserID int
}
type Verdict struct {
    TransactionID int
    Check         string
    Passed        bool
}
// Fan-Out: each checker runs independently on the same transaction
func velocityCheck(transactions <-chan Transaction, out chan<- Verdict, wg *sync.WaitGroup) {
    defer wg.Done()
    for tx := range transactions {
        time.Sleep(100 * time.Millisecond) // simulate check
        out <- Verdict{tx.ID, "velocity", tx.Amount < 10000}
    }
}
func locationCheck(transactions <-chan Transaction, out chan<- Verdict, wg *sync.WaitGroup) {
    defer wg.Done()
    for tx := range transactions {
        time.Sleep(150 * time.Millisecond) // simulate check
        out <- Verdict{tx.ID, "location", tx.UserID != 999} // 999 = flagged user
    }
}
func blacklistCheck(transactions <-chan Transaction, out chan<- Verdict, wg *sync.WaitGroup) {
    defer wg.Done()
    for tx := range transactions {
        time.Sleep(80 * time.Millisecond) // simulate check
        out <- Verdict{tx.ID, "blacklist", tx.Amount != 666} // 666 = known fraud amount
    }
}
// Fan-In: merge all verdict channels into one
func mergeVerdicts(channels ...<-chan Verdict) <-chan Verdict {
    merged := make(chan Verdict, 10)
    var wg sync.WaitGroup
    forward := func(ch <-chan Verdict) {
        defer wg.Done()
        for v := range ch {
            merged <- v
        }
    }
    wg.Add(len(channels))
    for _, ch := range channels {
        go forward(ch)
    }
    go func() {
        wg.Wait()
        close(merged)
    }()
    return merged
}
func main() {
    transactions := []Transaction{
        {1, 500, 101},
        {2, 15000, 102}, // fails velocity check
        {3, 200, 999},   // fails location check
        {4, 666, 103},   // fails blacklist check
    }
    // Create one channel per checker (Fan-Out)
    velocityCh := make(chan Transaction, len(transactions))
    locationCh := make(chan Transaction, len(transactions))
    blacklistCh := make(chan Transaction, len(transactions))
    verdictV := make(chan Verdict, 10)
    verdictL := make(chan Verdict, 10)
    verdictB := make(chan Verdict, 10)
    var wg sync.WaitGroup
    wg.Add(3)
    go velocityCheck(velocityCh, verdictV, &wg)
    go locationCheck(locationCh, verdictL, &wg)
    go blacklistCheck(blacklistCh, verdictB, &wg)
    // Send each transaction to all checkers simultaneously
    for _, tx := range transactions {
        velocityCh <- tx
        locationCh <- tx
        blacklistCh <- tx
    }
    close(velocityCh)
    close(locationCh)
    close(blacklistCh)
    // Fan-In: merge all verdicts into one stream
    go func() {
        wg.Wait()
        close(verdictV)
        close(verdictL)
        close(verdictB)
    }()
    merged := mergeVerdicts(verdictV, verdictL, verdictB)
    for verdict := range merged {
        status := "✓ PASSED"
        if !verdict.Passed {
            status = "✗ FAILED"
        }
        fmt.Printf("Transaction #%d | %s check | %s\n",
            verdict.TransactionID, verdict.Check, status)
    }
}
Output (order of verdicts may vary):

Transaction #1 | velocity check | ✓ PASSED
Transaction #1 | location check | ✓ PASSED
Transaction #1 | blacklist check | ✓ PASSED
Transaction #2 | velocity check | ✗ FAILED
Transaction #3 | location check | ✗ FAILED
Transaction #4 | blacklist check | ✗ FAILED
The key function is mergeVerdicts. It takes multiple receive-only channels and forwards everything into one merged channel that's the Fan-In. The caller only has to range over one channel instead of three, regardless of how many checkers you add later.





eg : 

Example: log processing pipeline

package main

import (
    "fmt"
    "strings"
    "time"
)
type LogEvent struct {
    Level   string
    Message string
    Time    time.Time
}
// Stage 1: parse raw log strings into structured events
func parse(lines <-chan string) <-chan LogEvent {
    out := make(chan LogEvent)
    go func() {
        for line := range lines {
            parts := strings.SplitN(line, " ", 2)
            if len(parts) != 2 {
                continue
            }
            out <- LogEvent{
                Level:   parts[0],
                Message: parts[1],
                Time:    time.Now(),
            }
        }
        close(out)
    }()
    return out
}
// Stage 2: filter - only pass ERROR and CRITICAL events
func filterErrors(events <-chan LogEvent) <-chan LogEvent {
    out := make(chan LogEvent)
    go func() {
        for event := range events {
            if event.Level == "ERROR" || event.Level == "CRITICAL" {
                out <- event
            }
        }
        close(out)
    }()
    return out
}
// Stage 3: format events for storage
func format(events <-chan LogEvent) <-chan string {
    out := make(chan string)
    go func() {
        for event := range events {
            out <- fmt.Sprintf("[%s] %s at %s",
                event.Level,
                event.Message,
                event.Time.Format("15:04:05"),
            )
        }
        close(out)
    }()
    return out
}
func main() {
    // Raw log lines coming in
    rawLogs := []string{
        "INFO  User logged in",
        "DEBUG Cache miss for key user:42",
        "ERROR Database connection failed",
        "INFO  Request completed in 120ms",
        "CRITICAL Disk usage above 95%",
        "DEBUG Query took 2ms",
        "ERROR Payment gateway timeout",
        "INFO  Health check passed",
    }
    // Source: feed raw lines into first stage
    source := make(chan string)
    go func() {
        for _, line := range rawLogs {
            source <- line
        }
        close(source)
    }()
    // Build the pipeline
    parsed := parse(source)
    filtered := filterErrors(parsed)
    formatted := format(filtered)
    // Consume: store the formatted error events
    fmt.Println("=== Error Events for Storage ===")
    for entry := range formatted {
        fmt.Println(entry)
    }
}
Output:

=== Error Events for Storage ===
[ERROR] Database connection failed at 14:22:31
[CRITICAL] Disk usage above 95% at 14:22:31
[ERROR] Payment gateway timeout at 14:22:31



