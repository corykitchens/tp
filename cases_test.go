package tp

var testCases = []struct {
	name     string
	command  string
	expected triviaQuestion
}{
	{
		name:     "Basic Format",
		command:  ":q Question :a Answer :c Category",
		expected: triviaQuestion{question: "Question", answer: "Answer", category: "Category"},
	},
	{
		name:     "White space",
		command:  ":q  Question :a  Answer :c    Category",
		expected: triviaQuestion{question: "Question", answer: "Answer", category: "Category"},
	},
	{
		name:     "Different Casing",
		command:  ":Q  Question :a  Answer :C    Category",
		expected: triviaQuestion{question: "Question", answer: "Answer", category: "Category"},
	},
}
