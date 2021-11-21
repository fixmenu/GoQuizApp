package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Quiz struct {
	problems []problem
}

type problem struct {
	q string
	a string
}

func NewQuizFromFile(fileName string) *Quiz {
	q := Quiz{}

	file, err := os.Open(fileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", fileName))
		os.Exit(1)
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file.")
	}

	q.parseLines(lines)

	fmt.Println(q)

	return &q
}

func (q *Quiz) parseLines(lines [][]string) {
	q.problems = make([]problem, len(lines))
	for i, line := range lines {
		q.problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
