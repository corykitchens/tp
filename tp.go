package tp

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

type triviaQuestion struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Category string `json:"category"`
}

var (
	api = os.Getenv("API_ENDPOINT")
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
		Question: strings.Title(strings.TrimSpace(questionText[0])),
		Answer:   strings.Title(strings.TrimSpace(answerText)),
		Category: strings.Title(strings.TrimSpace(categoryText)),
	}
	return s
}

func SubmitQuestion(tq triviaQuestion) {
	m, err := json.Marshal(tq)
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST",
		api,
		bytes.NewBuffer(m),
	)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
}
