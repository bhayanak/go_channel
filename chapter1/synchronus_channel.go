package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)     //Create a channel of string
	go do_stuff("Starting", c) //Launch go routine and pass channel along
	for i := 0; i < 5; i++ {
		fmt.Printf("Recieved :%q\n", <-c) //Print as and when something comes to channel back
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Ended")
}

func do_stuff(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) //send string to channel
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	}
}
