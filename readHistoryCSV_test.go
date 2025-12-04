package main

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

// scaffold to force a read error
type errReader struct{}

func (e errReader) Read(_ []byte) (n int, err error) {
	return 0, errors.New("forced read error")
}

func Test_readHistoryCSV(t *testing.T) {

	tests := []struct {
		name        string
		r           io.Reader
		wantHistory []int
		wantErr     bool
	}{
		{"header only should return empty array",
			strings.NewReader("iteration, done"),
			[]int{},
			false},
		{"done value not an integer should throw an error",
			strings.NewReader("a,done\n1,zwei"),
			nil,
			true},
		{"one data row should return one value",
			strings.NewReader("#iteration,done\na,1"),
			[]int{1},
			false},
		{"zwo data row should return two values",
			strings.NewReader("#iteration,done\na,1\nb,2"),
			[]int{1, 2},
			false},
		{"reader error should throw an error",
			errReader{},
			nil,
			true},
		// Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHistory, err := readHistoryCSV(tt.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("readHistoryCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotHistory, tt.wantHistory) {
				t.Errorf("readHistoryCSV() gotHistory = %v, want %v", gotHistory, tt.wantHistory)
			}
		})
	}
}

func Test_doneColumnPosition(t *testing.T) {
	tests := []struct {
		name    string
		header  []string
		wantPos int
	}{
		{`"done" on position 1`, []string{"a", "done"}, 1},
		{`"done" on position 0`, []string{"done", "b"}, 0},
		{`"DoNe" on position 0`, []string{"DoNe", "b"}, 0},
		{`"  Done   " on position 0`, []string{"  Done   ", "b"}, 0},
		// Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPos := doneColumnPosition(tt.header); gotPos != tt.wantPos {
				t.Errorf("doneColumnPosition() = %v, want %v", gotPos, tt.wantPos)
			}
		})
	}
}
