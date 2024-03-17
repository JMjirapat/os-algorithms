package main

import (
	"fmt"

	scheduling "github.com/JMjirapat/OS-algorithms/Scheduling"
)

func main() {
	// refStr := []int{1, 2, 3, 2, 1, 4, 3, 5, 6, 4, 3, 5, 3, 5, 6, 7, 2, 1, 5, 7}
	// refStr := []int{11, 12, 10, 3, 7, 14, 1, 12, 10, 11,
	// 	12, 10, 3, 14, 12, 10, 11, 12, 10, 11,
	// 	4, 5, 12, 14, 12, 10, 11, 12, 7, 12,
	// 	10, 14, 12, 10, 3, 14, 12, 10, 11, 12,
	// 	10, 11, 12, 7, 12, 10, 14, 12, 10, 11,
	// 	12, 14, 12, 10, 11, 12, 7, 12, 10, 14,
	// 	12, 10, 3, 14, 12, 10, 11, 12, 10, 11,
	// 	12, 7, 12, 10, 14, 12, 10, 11, 12, 10,
	// 	11, 12, 7, 12, 10, 14, 12, 10, 11}
	// refStr := []int{7, 0, 1, 2, 0, 3, 0, 4, 2, 3, 0, 3, 2, 1, 2, 0, 1, 7}
	// fifo := PFR.NewLRU(4, refStr)
	// fmt.Println(fifo.Run())
	totalProcesses := 30
	short_probability := 40
	medium_probability := 30
	long_probability := 30
	processes := scheduling.GenerateProcesses(totalProcesses, short_probability, medium_probability, long_probability)

	fmt.Println("Arrival Time\tBurst Time")
	for _, p := range processes {
		fmt.Printf("%d\t\t%d\n", p.ArrivalTime, p.BurstTime)
	}

	shuffledProcesses := scheduling.ShuffleProcesses(processes)
	fmt.Println("Arrival Time\tBurst Time")
	for _, p := range shuffledProcesses {
		fmt.Printf("%d\t\t%d\n", p.ArrivalTime, p.BurstTime)
	}

	arrivalTime := make([]int, totalProcesses)
	burstTime := make([]int, totalProcesses)
	for i, p := range shuffledProcesses {
		arrivalTime[i] = p.ArrivalTime
		burstTime[i] = p.BurstTime
	}
	scheduling.FCFS(arrivalTime, burstTime)
	scheduling.SJF(arrivalTime, burstTime)
	scheduling.RR(arrivalTime, burstTime, 10)
}
