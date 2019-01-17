package tp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type triviaQuestion struct {
	question string
	answer   string
	category string
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
		question: strings.Title(strings.TrimSpace(questionText[0])),
		answer:   strings.Title(strings.TrimSpace(answerText)),
		category: strings.Title(strings.TrimSpace(categoryText)),
	}
	return s
}

func SubmitQuestion(tq triviaQuestion) error {
	// m, err := json.Marshal(tq)
	// if err != nil {
	// 	return err
	// }
	u := triviaQuestion{question: "test", answer: "answer", category: "test"}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)

	client := &http.Client{}
	req, err := http.NewRequest("POST",
		api,
		b)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	spew.Dump(resp)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	spew.Dump(resp.Body)
	// defer func() error {
	// 	if resp != nil && resp.Body != nil {
	// 		resp.Body.Close()
	// 		return nil
	// 	}
	// 	return errors.New("Error body of response is nil")
	// }()
	return nil
}
