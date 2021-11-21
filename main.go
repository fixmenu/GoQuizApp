package main

import (
	"flag"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	starter := NewQuizStarter(*timeLimit, *NewQuizFromFile(*csvFilename))
	starter.start()
}
