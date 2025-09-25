package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

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
