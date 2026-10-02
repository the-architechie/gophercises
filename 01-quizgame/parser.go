package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Parser struct {
	fileName  string
	timeLimit time.Duration
}

type Score struct {
	right int
	wrong int
	total int
}

type Quiz struct {
	question string
	ans      int
}

func (q Parser) Parse() ([][]string, error) {
	f, err := os.Open(q.fileName)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the files: %v", err)
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to parse csv: %v", err)
	}
	return records, nil
}

func (q Parser) TokeniseTheQuiz(records [][]string) []Quiz {
	var quiz []Quiz
	for _, r := range records {
		val, err := strconv.Atoi(r[1])
		if err != nil {
			fmt.Printf("Failed to parse : %v, skipping", r[1])
			continue
		}
		q := Quiz{
			question: r[0],
			ans:      val,
		}
		quiz = append(quiz, q)
	}
	return quiz

}
