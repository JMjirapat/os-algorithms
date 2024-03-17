package scheduling

import "fmt"

func RR(arrivalTime []int, burstTime []int, quantum int) {
	n := len(arrivalTime)
	remainingBurst := make([]int, n)
	copy(remainingBurst, burstTime)

	currentTime := 0
	completionTime := make([]int, n)
	turnaroundTime := make([]int, n)
	waitingTime := make([]int, n)

	for {
		done := true
		for i := 0; i < n; i++ {
			if remainingBurst[i] > 0 {
				done = false
				if remainingBurst[i] > quantum {
					currentTime += quantum
					remainingBurst[i] -= quantum
				} else {
					currentTime += remainingBurst[i]
					waitingTime[i] = currentTime - burstTime[i] - arrivalTime[i]
					remainingBurst[i] = 0
					completionTime[i] = currentTime
				}
			}
		}
		if done {
			break
		}
	}

	for i := 0; i < n; i++ {
		turnaroundTime[i] = completionTime[i] - arrivalTime[i]
	}

	fmt.Println("\nRound Robin Scheduling (Quantum:", quantum, "):")
	fmt.Println("Process\tArrival Time\tBurst Time\tCompletion Time\tTurnaround Time\tWaiting Time")
	for i := 0; i < n; i++ {
		fmt.Printf("%d\t%d\t\t%d\t\t%d\t\t%d\t\t%d\n", i+1, arrivalTime[i], burstTime[i], completionTime[i], turnaroundTime[i], waitingTime[i])
	}
}
