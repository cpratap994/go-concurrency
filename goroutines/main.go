//  package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	fmt.Println("Running goroutine..")
// 	go gorun("routine1")
// 	go gorun("routine2")

// 	// sleep here to give some time to run the goroutine
// 	time.Sleep(1 * time.Millisecond)
// }

// //values can not be returned from goroutines
// func gorun(name string) {
// 	for i := 0; i < 200; i++ {
// 		fmt.Println("running gorun " + name + " " + strconv.Itoa(i))
// 	}
// }

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	go gorun("runner1")

	time.Sleep(1 * time.Second)
}

func gorun(name string) {
	for i := 0; i < 200; i++ {
		fmt.Println("name " + name + " " + strconv.Itoa(i))
	}
}
