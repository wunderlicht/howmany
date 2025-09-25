package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
