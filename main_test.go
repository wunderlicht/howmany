package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_scenario(t *testing.T) {
	type args struct {
		historicalData []int
		target         int
	}
	tests := []struct {
		name           string
		args           args
		wantIterations int
	}{
		{"should be done in one iteration", args{[]int{2, 2, 2}, 2}, 1},
		{"should be done in two iteration", args{[]int{2, 2, 2}, 4}, 2},
		{"should be done in three iteration", args{[]int{4, 3, 3}, 9}, 3},
		{"one element", args{[]int{3}, 12}, 4},
		// Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIterations := scenario(tt.args.historicalData, tt.args.target); gotIterations != tt.wantIterations {
				t.Errorf("scenario() = %v, want %v", gotIterations, tt.wantIterations)
			}
		})
	}
}

func Test_scenario_should_panic_on_empty_historicData(t *testing.T) {
	defer func() {
		_ = recover()
	}()
	_ = scenario([]int{}, 4) //this should panic
	// If there was no panic the test will fail
	t.Errorf("scenario() with empty historic data should have paniced but didn't")
}

// strategy is to look if specific strings appear in the output rather than matching the complete output
func Test_formatHistogram(t *testing.T) {
	tests := []struct {
		name      string
		counts    map[int]int
		scenarios int
		want      string
	}{
		{"should contain a header",
			map[int]int{1: 10, 2: 30}, 40,
			"#iterations probably cumulative occurrence\n"},
		{"should contain one row",
			map[int]int{1: 42}, 42,
			"          1   100.00     100.00         42"},
		{"should contain confidence marker",
			map[int]int{1: 42}, 42,
			" <-- 85% confidence\n"},
		// Add more test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatHistogram(tt.counts, tt.scenarios)
			if !strings.Contains(got, tt.want) {
				t.Errorf("formatHistogram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatHistogram_should_have_marker(t *testing.T) {
	const (
		want   = 1 //there can only be one marker
		marker = " <-- "
	)

	tests := []struct {
		name      string
		counts    map[int]int
		scenarios int
	}{
		{"one line one marker",
			map[int]int{1: 42}, 42,
		},
		{"two line one marker",
			map[int]int{1: 10, 2: 30}, 40,
		},
		// Add more test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatHistogram(tt.counts, tt.scenarios)
			if strings.Count(got, marker) != want {
				t.Errorf("formatHistogram() = %v, has not exactly %d marker", got, want)
			}
		})
	}
}

func Test_percent(t *testing.T) {
	tests := []struct {
		name  string
		value int
		total int
		want  float64
	}{
		{"all is 100%", 42, 42, 100.00},
		{"half is 50%", 21, 42, 50.00},

		//Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := percent(tt.value, tt.total); got != tt.want {
				t.Errorf("percent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runSimulation(t *testing.T) {
	type args struct {
		historicalData []int
		target         int
		scenarios      int
	}
	tests := []struct {
		name           string
		args           args
		wantOccurences map[int]int
	}{
		{"one scenario one datapoint",
			args{[]int{2}, 6, 1},
			map[int]int{3: 1},
		},
		{"50 scenarios one datapoint",
			args{[]int{2}, 6, 1},
			map[int]int{3: 50},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOccurences := runSimulation(tt.args.historicalData, tt.args.target, tt.args.scenarios); !reflect.DeepEqual(gotOccurences, tt.wantOccurences) {
				t.Errorf("runSimulation() = %v, want %v", gotOccurences, tt.wantOccurences)
			}
		})
	}
}
