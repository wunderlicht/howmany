package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
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
		header = "#iterations probably cumulative occurrence\n"
		row    = "%11d%9.2f%11.2f%11d" //row format
		marker = " <-- 85% confidence"
	)
	var b strings.Builder
	b.WriteString(header)

	scensum := 0
	cumulative := 0.0
	belowThreshold := true
	for iter := 1; scensum < scenarios; iter++ {
		scensum += occurrences[iter]
		prob := percent(occurrences[iter], scenarios)
		cumulative += prob

		b.WriteString(fmt.Sprintf(row, iter, prob, cumulative, occurrences[iter]))
		if cumulative >= 85.0 && belowThreshold {
			belowThreshold = false //we reached the threshold
			b.WriteString(marker)
		}
		b.WriteString("\n")
	}
	return b.String()
}

// could have been a generic but creat is better than clever.
// will return +Inf when total is 0
func percent(value, total int) float64 {
	return float64(value) / float64(total) * 100
}
