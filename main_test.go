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
			strings.NewReader("iteration, completed"),
			[]int{},
			false},
		{"done value not an integer should throw an error",
			strings.NewReader("a,b\n1,zwei"),
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
