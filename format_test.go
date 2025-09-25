package main

import (
	"strings"
	"testing"
)

// strategy is to look if specific strings appear in the output
// rather than matching the complete output
func Test_formatHistogram(t *testing.T) {
	type args struct {
		counts     map[int]int
		scenarios  int
		confidence float64
	}
	tests := []struct {
		name     string
		args     args
		search   string
		contains bool
	}{
		{"should contain a header",
			args{map[int]int{1: 10, 2: 30}, 40, 85.0},
			"#iterations occurrence probability cumulative\n", true},
		{"should contain one row",
			args{map[int]int{1: 42}, 42, 85.0},
			"          1         42   100.00     100.00", true},
		{"should contain default confidence marker",
			args{map[int]int{1: 42}, 42, 85.0},
			" <-- 85.0% confidence\n", true},
		{"should contain 99.9% confidence marker",
			args{map[int]int{1: 42}, 42, 99.9},
			" <-- 99.9% confidence\n", true},
		{"should contain no confidence marker",
			args{map[int]int{1: 42}, 42, 0.0},
			" <-- ", false},
		// Add more test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatHistogram(tt.args.counts, tt.args.scenarios, tt.args.confidence)
			if strings.Contains(got, tt.search) != tt.contains {
				t.Errorf("formatHistogram() = %v, want %v", got, tt.search)
			}
		})
	}
}

func Test_formatHistogram_should_have_one_marker(t *testing.T) {
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
		{"two lines one marker",
			map[int]int{1: 10, 2: 30}, 40,
		},
		// Add more test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatHistogram(tt.counts, tt.scenarios, 85.0)
			if strings.Count(got, marker) != want {
				t.Errorf("formatHistogram() = %v, has not exactly %d marker", got, want)
			}
		})
	}
}

func Test_formatAverage(t *testing.T) {
	type args struct {
		history []int
		goal    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty array should return empty string",
			args{[]int{}, 50},
			""},
		{"nil array should return empty string",
			args{nil, 50},
			""},
		{"zero goal should return empty string",
			args{[]int{1, 2, 3}, 0},
			""},
		{"goal 10 should return 2",
			args{[]int{5, 5, 5}, 10},
			"Average: 5.00\nIterations based on average: 2.0\n"},
		//Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatPredictionOnAverage(tt.args.history, tt.args.goal); got != tt.want {
				t.Errorf("formatPredictionOnAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
