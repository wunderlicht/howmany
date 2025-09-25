package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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
	defer closeIgnoreError(f)

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

func closeIgnoreError(f *os.File) {
	_ = f.Close()
}
