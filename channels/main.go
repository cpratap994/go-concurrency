package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//chnl := make(chan string)
	chnl := make(chan string, 2)

	go sender(chnl)
	go receiver(chnl)

	// sleep here to give some time to run the goroutine
	time.Sleep(100 * time.Second)
}

func sender(chnl chan string) {
	i := 0

	for {
		fmt.Println("sending " + strconv.Itoa(i))
		chnl <- strconv.Itoa(i)
		fmt.Println("sender was blocked " + strconv.Itoa(i))
		i++
	}
}

func receiver(chnl chan string) {
	for {
		fmt.Println("received " + <-chnl)
		time.Sleep(10 * time.Second)
	}
}
