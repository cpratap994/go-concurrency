package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {

	// go printStrings(10)
	// fmt.Println("exiting main..")

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go printStrings(wg, 10)

	fmt.Println("waiitng goroutine to finish")
	wg.Wait()
	fmt.Println("exiting main..")

}

func printStrings(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	for i := 0; i < num; i++ {
		fmt.Println("pritning " + strconv.Itoa(i))
	}
	//wg.Done()
}
