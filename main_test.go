package main

import (
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
		{"should be done in three iteration", args{[]int{3, 2, 2}, 6}, 3},
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
			"#iterations probably cumulative occurrence"},
		{"should contain one row",
			map[int]int{1: 42}, 42,
			"          1   100.00     100.00         42"},
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
