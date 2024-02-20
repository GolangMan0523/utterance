package utterance

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsFragment(t *testing.T) {
	testCases := []struct {
		word     string
		expected bool
	}{
		{"sentence", true},
		{"Fragment.", false},
		{"CAPITALIZED", true},
		{"123", true},
	}

	for _, tc := range testCases {
		t.Run(tc.word, func(t *testing.T) {
			result := IsFragment(tc.word)
			if result != tc.expected {
				t.Errorf("Expected IsFragment(%s) to be %v, but got %v", tc.word, tc.expected, result)
			}
		})
	}
}

func TestFixAttribution(t *testing.T) {
	testCases := []struct {
		utterance Utterance
		prev      Utterance
		expected  Utterance
	}{
		{
			Utterance{Speaker: "A", Text: "going well. Thanks for asking."},
			Utterance{Speaker: "B", Text: "How is the product launch going? It is not     "},
			Utterance{Speaker: "A", Text: "It is not going well. Thanks for asking."},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v - %v", tc.prev, tc.utterance), func(t *testing.T) {
			FixAttribution(&tc.utterance, &tc.prev)
			if !reflect.DeepEqual(tc.utterance, tc.expected) {
				t.Errorf("Expected FixAttribution(%v, %v) to be %v, but got %v", tc.utterance, tc.prev, tc.expected, tc.utterance)
			}
		})
	}
}

func TestProcessUtterances(t *testing.T) {
	testCases := []struct {
		utterances []Utterance
		expected   Transcript
	}{
		// Test case 1: Basic scenario with two utterances
		{
			utterances: []Utterance{
				{Speaker: "A", Text: "How is the product launch going? It is not     ", TimestampMs: 1},
				{Speaker: "B", Text: "going well. Thanks for asking.", TimestampMs: 2},
			},
			expected: Transcript{
				Utterances: []Utterance{
					{Speaker: "A", Text: "How is the product launch going?", TimestampMs: 1},
					{Speaker: "B", Text: "It is not going well. Thanks for asking.", TimestampMs: 2},
				},
			},
		},

		// Test case 2: Handling fragments and empty strings
		{
			utterances: []Utterance{
				{Speaker: "A", Text: "How is the product launch going? It is not     ", TimestampMs: 1},
				{Speaker: "B", Text: "going well. Thanks for asking. What are you", TimestampMs: 2},
				{Speaker: "C", Text: "doing? I am running.", TimestampMs: 3},
			},
			expected: Transcript{
				Utterances: []Utterance{
					{Speaker: "A", Text: "How is the product launch going?", TimestampMs: 1},
					{Speaker: "B", Text: "It is not going well. Thanks for asking.", TimestampMs: 2},
					{Speaker: "C", Text: "What are you doing? I am running.", TimestampMs: 3},
				},
			},
		},

		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.utterances), func(t *testing.T) {
			result := ProcessUtterances(tc.utterances)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected ProcessUtterances(%v) to be %v, but got %v", tc.utterances, tc.expected, result)
			}
		})
	}
}
