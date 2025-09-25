package main

import (
	"testing"
)

func Test_percent(t *testing.T) {
	tests := []struct {
		name  string
		value int
		total int
		want  float64
	}{
		{"all is 100%", 42, 42, 100.00},
		{"half is 50%", 21, 42, 50.00},
		{"nothing is 0%", 0, 42, 0.00},
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

func Test_average(t *testing.T) {
	tests := []struct {
		name string
		d    []int
		want float64
	}{
		{"empty array should be 0.0", []int{}, 0},
		{"nil array should be 0.0", nil, 0},
		{"array with one element should be the element", []int{5}, 5.0},
		{"array 1-10 should be 5.5", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5.5},
		{"array -3,0,+3 should be 0", []int{-3, -2, -1, 0, 1, 2, 3}, 0},
		//Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.d); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
