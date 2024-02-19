package main

import (
	"fmt"

	PFR "github.com/JMjirapat/OS-algorithms/PFR"
)

func main() {
	refStr := []int{1, 2, 3, 2, 1, 4, 3, 5, 6, 4, 3, 5, 3, 5, 6, 7, 2, 1, 5, 7}
	// refStr := []int{7, 0, 1, 2, 0, 3, 0, 4, 2, 3, 0, 3, 2, 1, 2, 0, 1, 7}
	fifo := PFR.NewOptimal(3, refStr)
	fmt.Println(fifo.Run())
}
