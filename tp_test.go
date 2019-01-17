package tp

import (
	"testing"
)

func TestTriviaParser(t *testing.T) {

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			received := TriviaParser(tc.command)
			if received != tc.expected {
				t.Errorf("Test Failed.. Expected: %v, Received: %v", tc.expected, received)
			}
		})
	}
}
