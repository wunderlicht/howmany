package main

import (
	"math/rand/v2"
)

// scenario picks randomly from the historical data. How many picks are
// necessary to match the goal? A goal of 0 returns 1 iteration
// Example
// historical = [2, 4, 3, 2], goal = 10
// possible picks: 4, 3, 2, 3	summed up: 12, needed iterations: 4
// possible picks: 4, 3, 3		summed up: 10, needed iterations: 3
func scenario(historicalData []int, goal int) (iterations int) {
	if goal == 0 { //guard
		return 1
	}
	sum := 0
	for iterations = 0; sum < goal; iterations++ {
		pick := rand.IntN(len(historicalData))
		sum += historicalData[pick]
	}
	return
}

// run # of scenarios and put the results in the respective bucket
// Example
// historical = [2, 4, 3, 2], goal = 10, scenarios 2
// possible picks: 4, 3, 2, 3	summed up: 12, needed iterations: 4
// possible picks: 4, 3, 3		summed up: 10, needed iterations: 3
// would result in occurrences = map[int]int{3:1, 4:1}
func runSimulation(historicalData []int, goal, scenarios int) (occurrences map[int]int) {
	occurrences = make(map[int]int)
	for i := 0; i < scenarios; i++ {
		iterationsNeeded := scenario(historicalData, goal)
		occurrences[iterationsNeeded]++
	}
	return
}
