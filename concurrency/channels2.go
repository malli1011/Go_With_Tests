package concurrency

import (
	"fmt"
	"time"
)

var greetings = []string{"Hello", "Ciao!", "Hola!", "Hej!", "Salut!", "Vannakam"}
var goodbyes = []string{"Goodbye!", "Arrivederci", "Adios!", "Hej Hej", "Bye!", "Nandri"}

func greet2(greetings []string, ch chan<- string) {
	fmt.Print("Greeter ready!")
	for _, g := range greetings {
		ch <- g
	}
	//Close the channel to notify the receiver that no more greeting will be sent.
	//if we don't close it, the receiver will keep getting zeor value from the channal
	//for each call. comment the below line and see the output.
	close(ch)
	fmt.Println("Greeter completed!")
}

func Greeting() {
	ch := make(chan string)

	go greet2(greetings, ch)

	time.Sleep(1 * time.Second)
	fmt.Println("Main ready!")

	for g := range ch {
		fmt.Println("Greeting received! ", g)

	}

}

func Greeting2() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go greet2(greetings, ch1)
	go greet2(goodbyes, ch2)
	time.Sleep(2 * time.Second)
	fmt.Println("Main ready!")

	for {
		select {
		//we can use the second return type (bool) of the channel
		//to validate if the channel has closed.
		case msg1, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			time.Sleep(500 * time.Millisecond)
			fmt.Println("received from ch1: ", msg1)
		case msg2, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			time.Sleep(1 * time.Second)
			fmt.Println("received from ch2: ", msg2)
		default:
			fmt.Println("Nothing received")
			return
		}
	}

}
