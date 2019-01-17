package tp

type triviaQuestion struct {
	question string
	answer   string
	category string
}

//TriviaParser parses a string with the given format
//":q Question :a Answer :c Category"
//and returns a struct of the following structure
//{question: Question, answer: Answer, category: Category}
func TriviaParser(command string) triviaQuestion {
	s := triviaQuestion{}
	return s
}
