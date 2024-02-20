package utterance

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

// Utterance represents a transcript event with speaker, text, and timestamp
type Utterance struct {
	Speaker     string `json:"speaker"`
	Text        string `json:"text"`
	TimestampMs int64  `json:"timestampMs"`
}

// Transcript represents a collection of utterances
type Transcript struct {
	Utterances []Utterance `json:"utterances"`
}

// IsFragment checks if a word is a sentence fragment or not
func IsFragment(word string) bool {
	// A simple heuristic is to check if the word is lowercase and does not end with punctuation
	return !strings.ContainsAny(word, ".?!")
}

// FixAttribution fixes the speaker attribution for a given utterance and the previous one
func FixAttribution(utterance *Utterance, prev *Utterance) {
	// Split the text into words
	words := strings.Split(prev.Text, " ")

	// If the first word is a fragment, append it to the previous utterance and remove it from the current one
	for IsFragment(words[len(words)-1]) {
		if words[len(words)-1] != "" {
			utterance.Text = words[len(words)-1] + " " + utterance.Text
		}
		words = words[:len(words)-1]
	}
	prev.Text = strings.Join(words, " ")
}

// ProcessUtterances processes a slice of utterances and returns a transcript
func ProcessUtterances(utterances []Utterance) Transcript {
	// Create an empty transcript
	transcript := Transcript{}
	length := len(utterances)

	// Loop through the utterances
	for i, utterance := range utterances {
		// If this is not the first utterance, fix the attribution with the previous one
		if i < length-1 {
			FixAttribution(&utterances[i+1], &utterance)
		}

		// Append the utterance to the transcript
		transcript.Utterances = append(transcript.Utterances, utterance)
	}

	return transcript
}

// ReadUtterances reads a JSON file and returns a slice of utterances
func ReadUtterances(filename string) ([]Utterance, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var utterances []Utterance
	err = json.Unmarshal(data, &utterances)
	if err != nil {
		return nil, err
	}

	return utterances, nil
}

// WriteTranscript writes a transcript to a JSON file
func WriteTranscript(transcript Transcript, filename string) error {
	data, err := json.MarshalIndent(transcript, "", "")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
