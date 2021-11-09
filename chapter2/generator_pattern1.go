package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := generate("Starting") // Function returns a channel
	for i := 0; i < 5; i++ {
		fmt.Printf("Recieved :%q\n", <-c) //Print as and when something comes to channel back
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Ended")
}

func generate(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) //send string to channel
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()
	return c
}
