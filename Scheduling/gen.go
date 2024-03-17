package scheduling

import (
	"math"
	"math/rand"
)

type Process struct {
	ArrivalTime int
	BurstTime   int
}

func GenerateProcesses(n int, short int, medium int, long int) []Process {
	var processes []Process
	if short+medium+long != 100 {
		panic("Short, medium, and long must sum to 100")
	}

	for i := 0; i < n; i++ {
		burstTime := generateBurstTime(short, medium)
		processes = append(processes, Process{
			ArrivalTime: rand.Intn(n * int(math.Log2(float64(n/10)))),
			BurstTime:   burstTime,
		})
	}

	return processes
}

func generateBurstTime(short int, medium int) int {
	randNum := rand.Intn(100)

	switch {
	case randNum < short: // 70% chance
		return rand.Intn(7) + 2 // 2 to 8 milliseconds
	case randNum < medium: // 20% chance (70% + 20% = 90%)
		return rand.Intn(11) + 20 // 20 to 30 milliseconds
	default: // 10% chance
		return rand.Intn(6) + 35 // 35 to 40 milliseconds
	}
}

func ShuffleProcesses(processes []Process) []Process {
	shuffled := make([]Process, len(processes))
	perm := rand.Perm(len(processes))
	for i, j := range perm {
		shuffled[i] = processes[j]
	}
	return shuffled
}
