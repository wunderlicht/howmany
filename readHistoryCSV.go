package main

import (
	"encoding/csv"
	"errors"
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

	donePos, err := doneColumnPosition(rows[0])
	if err != nil {
		return nil, err
	}
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

// errors returnable by doneColumnPosition()
var (
	noDone       = errors.New("no done found")
	multipleDone = errors.New("multiple dones found")
)

// returns the position of the first occurrence of the done column of the CSV
func doneColumnPosition(header []string) (pos int, err error) {
	const notFound = -1
	pos = notFound //default to invalid
	for idx, col := range header {
		found := strings.TrimSpace(strings.ToLower(col)) == "done"
		if found {
			if pos != notFound { //we found something earlier
				return -1, multipleDone
			}
			pos = idx
		}
	}
	if pos == notFound { //♫♫ but we still haven't found what we were looking for
		return notFound, noDone
	}
	return pos, nil
}

func closeIgnoreError(f *os.File) {
	_ = f.Close()
}
