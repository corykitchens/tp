package tp

var testCases = []struct {
	name     string
	command  string
	expected triviaQuestion
}{
	{
		name:     "Basic Format",
		command:  ":q Question :a Answer :c Category",
		expected: triviaQuestion{Question: "Question", Answer: "Answer", Category: "Category"},
	},
}
