package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Vappelquist/quiz/api"
	"github.com/Vappelquist/quiz/quiz-folder"
	"github.com/spf13/cobra"
)

func fetchQuestions() ([]quiz.PublicQuestion, error) {
	resp, err := http.Get("http://localhost:8080/questions")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var questions []quiz.PublicQuestion
	if err := json.NewDecoder(resp.Body).Decode(&questions); err != nil {
		return nil, err
	}

	return questions, nil
}

func askQuestions(questions []quiz.PublicQuestion) []api.Answer {
	reader := bufio.NewReader(os.Stdin)
	var answers []api.Answer

	for _, q := range questions {
		fmt.Println()
		fmt.Println(q.Text)
		for _, opt := range q.Options {
			fmt.Printf("  %s) %s\n", opt.Id, opt.Text)
		}
		fmt.Print("Your answer: ")
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)
		answers = append(answers, api.Answer{
			QuestionId: q.Id,
			OptionId:   choice,
		})
	}
	return answers
}

func submitAnswers(answers []api.Answer) (*api.SubmitResult, error) {
	body, err := json.Marshal(api.SubmitRequest{Answers: answers})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post("http://localhost:8080/submit", "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var result api.SubmitResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func printResult(result *api.SubmitResult) {
	fmt.Printf("You scored %d out of %d.\n", result.Score, result.Total)
	if result.HasPercentile {
		fmt.Printf("You were better than %.0f%% of all quiz-takers.\n", result.Percentile)
	} else {
		fmt.Println("You're the first one to take this quiz!")
	}
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play the quiz",
	Run: func(cmd *cobra.Command, args []string) {
		questions, err := fetchQuestions()
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		answers := askQuestions(questions)
		result, err := submitAnswers(answers)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		printResult(result)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
