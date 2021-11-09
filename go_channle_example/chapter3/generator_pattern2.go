package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Channels as a handle on a service
func main() {
	joe := generate("Joe")
	ann := generate("Ann")
	for i := 0; i < 5; i++ {
		fmt.Printf(<-joe)
		fmt.Printf(<-ann)
	}
	fmt.Println("Ended")
}

func generate(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d\n", msg, i) //send string to channel
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()
	return c
}
