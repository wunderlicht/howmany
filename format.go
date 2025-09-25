package main

import (
	"fmt"
	"strings"
)

// formats the results for output. Changes to the appearance are made here.
func formatHistogram(occurrences map[int]int, scenarios int, confidence float64) string {
	const (
		header = "#iterations occurrence probability cumulative\n"
		row    = "%11d%11d%9.2f%11.2f" //row format
		marker = " <-- %.1f%% confidence"
	)
	var b strings.Builder
	b.WriteString(header)

	scenarioCounted := 0
	cumulative := 0.0
	markerStillToPrint := true
	if confidence == 0.0 { //don't print the marker if confidence is 0.0
		markerStillToPrint = false
	}
	for iter := 1; scenarioCounted < scenarios; iter++ {
		scenarioCounted += occurrences[iter]
		prob := percent(occurrences[iter], scenarios)
		cumulative += prob

		b.WriteString(fmt.Sprintf(row, iter, occurrences[iter], prob, cumulative))
		if cumulative >= confidence && markerStillToPrint {
			markerStillToPrint = false //we reached the threshold
			b.WriteString(fmt.Sprintf(marker, confidence))
		}
		b.WriteString("\n")
	}
	return b.String()
}

// what would it be if we just took the average
func formatPredictionOnAverage(history []int, goal int) string {
	//guard
	if goal == 0 || len(history) == 0 || history == nil {
		return ""
	}
	const format = "Average: %.2f\nIterations based on average: %.1f\n"

	avg := average(history)
	iterations := float64(goal) / avg
	return fmt.Sprintf(format, avg, iterations)
}
