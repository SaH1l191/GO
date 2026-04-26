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


//fan out (one input -> many workers )

func worker(id int,jobs <- chan int , results chan <- int ){
	for j:= range jobs {
		fmt.Printf("Worker %d processing job %d\n",id,j)
		time.Sleep(500 * time.Millisecond)//suimulate real work 
		results <- j * 2
	}
}

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


//fan in remaining 
// Pipeline pattern










