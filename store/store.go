package store

import (
	"sync"

	"github.com/Vappelquist/quiz/quiz-folder"
)

type Store struct {
	mu        sync.RWMutex
	questions []quiz.Question
	results   []int
}

func New() *Store {
	return &Store{
		questions: quiz.Seed(),
	}
}

func (s *Store) Questions() []quiz.Question {
	return s.questions
}

func (s *Store) Percentile(score int) (float64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.results) == 0 {
		return 0, false
	}
	beaten := 0
	for _, r := range s.results {
		if score > r {
			beaten++
		}
	}
	percentile := float64(beaten) / float64(len(s.results)) * 100
	return percentile, true
}

func (s *Store) SaveResult(score int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.results = append(s.results, score)
}

func (s *Store) IsCorrect(questionId, optionId string) bool {
	for _, q := range s.questions {
		if q.Id == questionId {
			return q.CorrectId == optionId
		}
	}
	return false
}
