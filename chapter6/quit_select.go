package main

import (
	"fmt"
	"math/rand"
)

//Channels as a handle on a service, previous example was still lock step,
//	here let the person talk if he is ready not to wait for another
func main() {
	quit := make(chan string) // Create a quit channel
	c := generate("Joe", quit)
	for i := rand.Intn(12); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Received from Joe: %q ,Ended\n", <-quit)
}

func generate(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i): //Doing nothing here
			case <-quit:
				cleanup()
				quit <- "See you!"
				return
			}
		}
	}()
	return c
}

func cleanup() {
	fmt.Println("wait! Cleaning up big time... :)")
}
