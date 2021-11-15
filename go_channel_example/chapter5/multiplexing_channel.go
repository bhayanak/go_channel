package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Channels as a handle on a service, previous example was still lock step,
//	here let the person talk if he is ready not to wait for another
func main() {
	c := fanIn(generate("Joe"), generate("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Ended")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func generate(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) //send string to channel
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e2)))
		}
	}()
	return c
}
