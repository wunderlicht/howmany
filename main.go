package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

const (
	defaultScenarios = 100_000
	envFile          = "HMFILE"
	envScenarios     = "HMSCENARIOS"
	envConfidence    = "HMCONFIDENCE"
	envAverage       = "HMAVERAGE"
)

type parameter struct {
	filename   string
	scenarios  int
	goal       int
	confidence float64
	average    bool
}

func main() {
	var param parameter
	populateParameter(&param)

	f, err := os.Open(param.filename)
	if err != nil {
		log.Fatal(err)
	}
	defer CloseIgnoreError(f)

	hd, err := readHistoryCSV(f)
	if err != nil {
		log.Fatal(err)
	}
	occ := runSimulation(hd, param.goal, param.scenarios)
	fmt.Print(formatHistogram(occ, param.scenarios, param.confidence))
	if param.average {
		fmt.Print(formatPredictionOnAverage(hd, param.goal))
	}
}

func populateParameter(p *parameter) {
	//goal can't be omitted or set by environment variable
	flag.IntVar(&p.goal, "goal", 0,
		"goal to meet for a scenario (mandatory)")
	flag.IntVar(&p.goal, "g", 0,
		"short for -goal")

	flag.StringVar(&p.filename, "file",
		getEnvOrDefaultString(envFile, ""),
		"filename of CSV with historic data (mandatory or env)")
	flag.StringVar(&p.filename, "f",
		getEnvOrDefaultString(envFile, ""),
		"short for -file")

	flag.IntVar(&p.scenarios, "scenarios",
		getEnvOrDefaultInt(envScenarios, defaultScenarios),
		"number of scenarios")
	flag.IntVar(&p.scenarios, "s",
		getEnvOrDefaultInt(envScenarios, defaultScenarios),
		"short for -scenarios")

	flag.Float64Var(&p.confidence, "confidence",
		getEnvOrDefaultFloat(envConfidence, 0.0),
		"set marker to # of iterations that meets confidence level")
	flag.Float64Var(&p.confidence, "c",
		getEnvOrDefaultFloat(envConfidence, 0.0),
		"short for -confidence")

	flag.BoolVar(&p.average, "average",
		getEnvOrDefaultBool(envAverage, false),
		"also print estimation based on average done")
	flag.BoolVar(&p.average, "a",
		getEnvOrDefaultBool(envAverage, false),
		"short for -average")
	flag.Parse()
}

// returns env's value when set otherwise fallback.
func getEnvOrDefaultString(key string, fallback string) string {
	val, present := os.LookupEnv(key)
	if !present {
		return fallback
	}
	return val
}

// returns env's value when set and well-formed otherwise fallback
func getEnvOrDefaultInt(key string, fallback int) int {
	val, present := os.LookupEnv(key)
	if !present {
		return fallback
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return i
}

// returns env's value when set and well-formed otherwise fallback
func getEnvOrDefaultFloat(key string, fallback float64) float64 {
	val, present := os.LookupEnv(key)
	if !present {
		return fallback
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return fallback
	}
	return f
}

// returns env's value when set and well-formed otherwise fallback
func getEnvOrDefaultBool(key string, fallback bool) bool {
	val, present := os.LookupEnv(key)
	if !present {
		return fallback
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		return fallback
	}
	return b
}

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
	iters := float64(goal) / avg
	return fmt.Sprintf(format, avg, iters)
}

// could have been a generic but clear is better than clever.
// will return +Inf when total is 0
func percent(value, total int) float64 {
	return float64(value) / float64(total) * 100
}

// could have been a generic but clear is better than clever.
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

func CloseIgnoreError(f *os.File) {
	_ = f.Close()
}
