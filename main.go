package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	hd := []int{2, 6, 1, 3, 4, 8}
	fmt.Println(scenario(hd, 6))
}

// scenario picks randomly from the historical data. How many picks are necessary to match the target?
// Example
// historical = [2, 4, 3, 2], target = 10
// possible picks: 4, 3, 2, 3	summed up: 12, needed iterations: 4
// possible picks: 4, 3, 3		summed up: 10, needed iterations: 3
func scenario(historicalData []int, target int) (iterations int) {
	sum := 0
	for iterations = 0; sum < target; iterations++ {
		pick := rand.IntN(len(historicalData))
		sum += historicalData[pick]
	}
	return
}

// formats the results for output. Changes to the appearance are made here.
func formatHistogram(occurrences map[int]int, scenarios int) string {
	const (
		header = "#iterations probably cumulative occurrence"
	)
	return header
}
