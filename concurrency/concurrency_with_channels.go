package concurrency

import (
	"fmt"
	"time"
)

func Channels_ex() {
	//ch:= make(chan string)
	/*
		In the above line, we just created a channel of string with out passing any length,
		This will block the sender until a receiver is available, as the receiver/sender will not know how many messages
		to be sent in the channel.

		In the below we pass a lengh of 1, so the sender know only one message to be sent. so it will not block the
		sender until a receiver is active/available.

		Comment the below line and uncommnet the above and check the order of print statements.

	*/
	ch := make(chan string, 1)

	go greet(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("Main ready")

	greeting := <-ch
	time.Sleep(2 * time.Second)
	fmt.Println("Greeting received!")
	fmt.Println(greeting)

}

func greet(ch chan string) {
	fmt.Println("Greeter ready!\nGreeter waiting to send greeting...")
	//greet
	ch <- "Hello world!"
	fmt.Println("Greeter completed!")
}
