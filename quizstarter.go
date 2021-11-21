package main

import (
	"fmt"
	"time"
)

type QuizStarter struct {
	q     Quiz
	timer *time.Timer
}

func NewQuizStarter(timeInSeconds int, q Quiz) QuizStarter {
	qs := QuizStarter{}
	qs.timer = time.NewTimer(time.Duration(timeInSeconds) * time.Second)
	qs.q = q

	return qs
}

func (qs *QuizStarter) start() {
	quiz := qs.q
	correct := 0

	for i, p := range quiz.problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-qs.timer.C:
			fmt.Printf("\nYou scored %d out of %d. \n", correct, len(quiz.problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d. \n", correct, len(quiz.problems))
}
