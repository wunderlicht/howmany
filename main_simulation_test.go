package main

import (
	"reflect"
	"testing"
)

func Test_scenario(t *testing.T) {
	type args struct {
		historicalData []int
		goal           int
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
		{"goal 0 should be done in one", args{[]int{2, 2, 2}, 0}, 1},
		// Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIterations := scenario(tt.args.historicalData, tt.args.goal); gotIterations != tt.wantIterations {
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

func Test_runSimulation(t *testing.T) {
	type args struct {
		historicalData []int
		goal           int
		scenarios      int
	}
	tests := []struct {
		name        string
		args        args
		occurrences map[int]int
	}{
		{"one scenario one datapoint",
			args{[]int{2}, 6, 1},
			map[int]int{3: 1},
		},
		{"50 scenarios one datapoint",
			args{[]int{2}, 6, 50},
			map[int]int{3: 50},
		},
		// Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOccurrences := runSimulation(tt.args.historicalData, tt.args.goal, tt.args.scenarios); !reflect.DeepEqual(gotOccurrences, tt.occurrences) {
				t.Errorf("runSimulation() = %v, want %v", gotOccurrences, tt.occurrences)
			}
		})
	}
}
