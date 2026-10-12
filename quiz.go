package main

import (
	"encoding/json"
	"os"
)

type Question struct {
	Prompt  string   `json:"question"`
	Options []string `json:"options"`
	Answer  int      `json:"answer"`
}

func LoadQuestions(filename string) ([]Question, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var questions []Question
	err = json.Unmarshal(file, &questions)
	return questions, err
}
