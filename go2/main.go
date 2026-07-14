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
