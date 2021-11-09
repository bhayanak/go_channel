package main

import (
	"fmt"
	"time"
)

//Channels as a handle on a service, previous example was still lock step,
//	here let the person talk if he is ready not to wait for another
func main() {
	c := fanIn(generate("Joe"), generate("Ann"))
	for i := 0; i < 10; i++ {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("Ended")
			return
		}
	}
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func generate(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) //send string to channel
		}
	}()
	return c
}
