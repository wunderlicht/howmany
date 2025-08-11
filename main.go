package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	const (
		filename  = "exampleHistory.csv"
		scenarios = 100_000
		target    = 56
	)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	hd, err := readHistoryCSV(f)
	if err != nil {
		log.Fatal(err)
	}
	occ := runSimulation(hd, target, scenarios)
	fmt.Print(formatHistogram(occ, scenarios))
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

// run # of scenarios and put the results in the respective bucket
// Example
// historical = [2, 4, 3, 2], target = 10, scenarios 2
// possible picks: 4, 3, 2, 3	summed up: 12, needed iterations: 4
// possible picks: 4, 3, 3		summed up: 10, needed iterations: 3
// would result in occurrences = map[int]int{3:1, 4:1}
func runSimulation(historicalData []int, target, scenarios int) (occurrences map[int]int) {
	occurrences = make(map[int]int)
	for i := 0; i < scenarios; i++ {
		occurrences[scenario(historicalData, target)]++
	}
	return
}

// returns an array of data read from the CVS. The reader is an open file.
func readHistoryCSV(r io.Reader) (history []int, err error) {

	history = make([]int, 0, 100) //allocate a bigger chunk to minimize reallocation
	rows, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, err
	}

	for i := 1; i < len(rows); i++ {
		const donePos = 1
		doneStr := rows[i][donePos]
		done, err := strconv.Atoi(doneStr)
		if err != nil {
			return nil, err
		}
		history = append(history, done)
	}
	return history, nil
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

	scenarioCounted := 0
	cumulative := 0.0
	belowThreshold := true
	for iter := 1; scenarioCounted < scenarios; iter++ {
		scenarioCounted += occurrences[iter]
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

func average(d []int) float64 {
	sum := 0.0
	if len(d) == 0 {
		return 0.0
	}
	for i := range d {
		sum += float64(d[i])
	}
	return sum / float64(len(d))
}
