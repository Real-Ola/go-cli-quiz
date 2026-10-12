package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Welcome to the Go CLI Quiz! ===")

	questions, err := LoadQuestions("questions.json")
	if err != nil {
		fmt.Println("Error loading questions:", err)
		return
	}

	score := 0
	for i, q := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Prompt)
		for _, opt := range q.Options {
			fmt.Println(opt)
		}

		var choice int
		fmt.Print("Your choice (enter option number): ")
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("❌ Invalid input. Skipping question.")
			continue
		}

		if choice == q.Answer {
			fmt.Println("✅ Correct!")
			score++
		} else {
			fmt.Println("❌ Incorrect!")
		}
	}

	fmt.Printf("\n=================================\n")
	fmt.Printf("Quiz Completed! Your Score: %d/%d\n", score, len(questions))
	fmt.Printf("=================================\n")
}
