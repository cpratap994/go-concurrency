package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	inputChannel := make(chan string)
	exit := make(chan bool)

	go printNumbers(100, exit, inputChannel)

	for {
		//Select statement are blocking
		select {
		case msg := <-inputChannel:
			fmt.Println("received message " + msg)

		case <-exit:
			fmt.Println("goroutine completed, exiting")
			os.Exit(1)

		case <-time.After(5 * time.Second):
			fmt.Println("go nothing on channels, exiting")
			os.Exit(1)

		default:
			fmt.Println("this is default case, if nothing matches to case, this will run and make it unblocking")

		}
	}

}

func printNumbers(max int, exit chan bool, input chan string) {

	for i := 0; i < max; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond)

	}

	time.Sleep(time.Second)

	exit <- true
}
