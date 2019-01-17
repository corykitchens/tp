package tp

import (
	"strings"
)

type triviaQuestion struct {
	question string
	answer   string
	category string
}

var (
	keys = [3]string{":q", ":a", ":c"}
)

//TriviaParser parses a string with the given format
//":q Question :a Answer :c Category"
//and returns a struct of the following structure
//{question: Question, answer: Answer, category: Category}
func TriviaParser(command string) triviaQuestion {
	//TODO
	//Develop more elegant/non brute force solution
	qMin := strings.Index(command, ":q")
	aMin := strings.Index(command, ":a")
	cMin := strings.Index(command, ":c")
	qMax := (qMin + aMin) - 1
	aMax := cMin - 1
	cMax := len(command)
	s := triviaQuestion{
		question: command[qMin+3 : qMax],
		answer:   command[aMin+3 : aMax],
		category: command[cMin+3 : cMax],
	}
	return s
}
