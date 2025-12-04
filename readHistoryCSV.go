package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
)

// returns an array of data read from the CVS. The reader is an open file.
func readHistoryCSV(r io.Reader) (history []int, err error) {

	history = make([]int, 0, 100) //allocate a bigger chunk to minimize reallocation
	rows, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, err
	}

	donePos := doneColumnPosition(rows[0])
	for i := 1; i < len(rows); i++ {
		doneStr := rows[i][donePos]
		done, err := strconv.Atoi(doneStr)
		if err != nil {
			return nil, err
		}
		history = append(history, done)
	}
	return history, nil
}

// returns the position of the first occurrence of the done column of the CSV
func doneColumnPosition(header []string) (pos int) {
	pos = -1 //default to invalid
	for idx, col := range header {
		if strings.TrimSpace(strings.ToLower(col)) == "done" {
			return idx
		}
	}
	return
}

func closeIgnoreError(f *os.File) {
	_ = f.Close()
}
