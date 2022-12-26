package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	// ch := make(chan string, 3)
	// for i := 0; i < 10; i++ {
	// 	ch <- "hello"
	// 	fmt.Println("Pushed a message to the channel")
	// }

	ch := make(chan string, 3)
	for i := 0; i < 10; i++ {
		select {
		case ch <- "hello":
			fmt.Println("Pushed a message to the channel")
		default:
			fmt.Println("WARN: The channel is full")
		}
	}
}
