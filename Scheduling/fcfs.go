package scheduling

import "fmt"

func FCFS(arrivalTime []int, burstTime []int) {
	n := len(arrivalTime)
	completionTime := make([]int, n)
	turnaroundTime := make([]int, n)
	waitingTime := make([]int, n)

	completionTime[0] = burstTime[0]
	for i := 1; i < n; i++ {
		completionTime[i] = completionTime[i-1] + burstTime[i]
	}

	for i := 0; i < n; i++ {
		turnaroundTime[i] = completionTime[i] - arrivalTime[i]
		waitingTime[i] = turnaroundTime[i] - burstTime[i]
	}

	fmt.Println("FCFS Scheduling:")
	fmt.Println("Process\tArrival Time\tBurst Time\tCompletion Time\tTurnaround Time\tWaiting Time")
	for i := 0; i < n; i++ {
		fmt.Printf("%d\t%d\t\t%d\t\t%d\t\t%d\t\t%d\n", i+1, arrivalTime[i], burstTime[i], completionTime[i], turnaroundTime[i], waitingTime[i])
	}
}
