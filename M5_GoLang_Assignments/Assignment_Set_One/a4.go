package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Question represents a single quiz question
type Question struct {
	Question string
	Options  [4]string
	Answer   int
}

// Function to take the quiz
func takeQuiz(questions []Question, timeLimit int) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	score := 0

	fmt.Println("Welcome to the Tech Quiz!")
	fmt.Println("Type 'exit' at any time to leave the quiz.\n")

	for i, q := range questions {
		fmt.Printf("Question %d: %s\n", i+1, q.Question)
		for j, opt := range q.Options {
			fmt.Printf("%d. %s\n", j+1, opt)
		}

		answerChan := make(chan int)
		go func() {
			for {
				fmt.Print("Enter your answer (1-4): ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if strings.ToLower(input) == "exit" {
					close(answerChan)
					return
				}

				answer, err := strconv.Atoi(input)
				if err != nil || answer < 1 || answer > 4 {
					fmt.Println("Invalid input. Please enter a number between 1 and 4.")
					continue
				}
				answerChan <- answer
				return
			}
		}()

		select {
		case answer := <-answerChan:
			if answer == q.Answer {
				score++
			}
		case <-time.After(time.Duration(timeLimit) * time.Second):
			fmt.Println("\nTime's up for this question!")
		}

		fmt.Println()
	}

	return score, nil
}

// Function to classify performance
func classifyPerformance(score, total int) string {
	percentage := (score * 100) / total
	switch {
	case percentage >= 90:
		return "Outstanding! Keep up the great work!"
	case percentage >= 75:
		return "Excellent Performance! Well done!"
	case percentage >= 50:
		return "Good Effort! You are improving!"
	default:
		return "Needs Improvement. Keep practicing!"
	}
}

// Main function
func main() {
	// Define the question bank
	questions := []Question{
		{
			Question: "What does HTML stand for?",
			Options:  [4]string{"HyperText Markup Language", "HighText Machine Language", "HyperText Management Language", "None of the above"},
			Answer:   1,
		},
		{
			Question: "Which programming language is known as the 'mother of all programming languages'?",
			Options:  [4]string{"Python", "C", "Java", "Assembly"},
			Answer:   2,
		},
		{
			Question: "What is the purpose of the 'git' command?",
			Options:  [4]string{"File compression", "Version control", "Network management", "Data encryption"},
			Answer:   2,
		},
		{
			Question: "What does 'AI' stand for in the tech world?",
			Options:  [4]string{"Automated Interface", "Artificial Intelligence", "Application Integration", "Advanced Internet"},
			Answer:   2,
		},
		{
			Question: "Which company developed the 'Go' programming language?",
			Options:  [4]string{"Microsoft", "Apple", "Google", "Amazon"},
			Answer:   3,
		},
		{
			Question: "What is the full form of 'IDE' in software development?",
			Options:  [4]string{"Internet Development Environment", "Integrated Development Environment", "Interactive Data Explorer", "Intelligent Design Editor"},
			Answer:   2,
		},
	}

	timeLimitPerQuestion := 15 // Time limit in seconds for each question

	// Start the quiz
	score, err := takeQuiz(questions, timeLimitPerQuestion)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Display the result
	fmt.Printf("\nYou completed the quiz!\nYour score: %d/%d\n", score, len(questions))
	fmt.Println("Performance:", classifyPerformance(score, len(questions)))
	fmt.Println("\nThank you for participating in the Tech Quiz!")
}
