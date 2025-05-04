package main

//channels ( and how to use them with go routines)
//channels are communication pipes that let goroutines exchange data safely without explicit locks or race conditions.
// channels can receive any type of data, like int, float, slice, struct etc...
// to define a channel u do like this : ch:= make(chan int)
// to send a value in the channel you use the arrow syntax : ch <- value
// to recieve from a channel you do this : value <- ch
// channels need closing so the program knows there are no more data coming, else you get an infinite loop ( deadlock error)
// to close it you call: close(ch)
// Directionality: Can specify if a channel is send-only (chan<- int) or receive-only (<-chan int)

//Real-World Example: Restaurant Kitchen System
// Imagine a restaurant with chefs and servers. The kitchen needs to coordinate food preparation and delivery without chaos.
// Go channels model this perfectly:

import (
	"fmt"
	"time"
)

type Order struct {
	ID    int
	Item  string
	Table int
}

func main6() {
	// Create a channel for food orders
	orders := make(chan Order)
	// Channel for completed orders
	completedOrders := make(chan Order)

	// Start chef goroutine
	go chef(orders, completedOrders)

	// Start server goroutine
	go server(completedOrders)

	// Send some orders to the kitchen
	orders <- Order{1, "Pasta", 3}
	orders <- Order{2, "Salad", 5}
	orders <- Order{3, "Steak", 2}
	
	// Wait a bit to see the results
	time.Sleep(10 * time.Second)
}

func chef(orders <-chan Order, completed chan<- Order) {
	for order := range orders {
		fmt.Printf("Chef: Preparing %s for table %d\n", order.Item, order.Table)
		
		// Simulate cooking time
		time.Sleep(2 * time.Second)
		
		fmt.Printf("Chef: Finished %s for table %d\n", order.Item, order.Table)
		completed <- order
	}
}

func server(completed <-chan Order) {
	for order := range completed {
		fmt.Printf("Server: Delivering %s to table %d\n", order.Item, order.Table)
		time.Sleep(1 * time.Second)
		fmt.Printf("Server: Delivered %s to table %d\n", order.Item, order.Table)
	}
}