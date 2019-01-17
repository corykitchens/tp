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

func splitTextFromSymbol(str, sym string) []string {
	return strings.Split(str, sym)
}

//TriviaParser parses a string with the given format
//":q Question :a Answer :c Category"
//and returns a struct of the following structure
//{question: Question, answer: Answer, category: Category}
func TriviaParser(command string) triviaQuestion {
	sanitizedCommand := strings.ToLower(command)
	splitCommand := strings.Split(sanitizedCommand, ":q")
	questionText := strings.Split(splitCommand[1], ":a")
	splitCommand = strings.Split(questionText[1], ":c")
	answerText := splitCommand[0]
	categoryText := splitCommand[1]
	s := triviaQuestion{
		question: strings.Title(strings.TrimSpace(questionText[0])),
		answer:   strings.Title(strings.TrimSpace(answerText)),
		category: strings.Title(strings.TrimSpace(categoryText)),
	}
	return s
}
