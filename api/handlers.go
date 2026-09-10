package api

import (
	"encoding/json"
	"net/http"

	"github.com/Vappelquist/quiz/quiz-folder"
	"github.com/Vappelquist/quiz/store"
)

type Handler struct {
	store *store.Store
}

func NewHandler(s *store.Store) *Handler {
	return &Handler{store: s}
}

func (h *Handler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions := h.store.Questions()
	public := make([]quiz.PublicQuestion, len(questions))
	for i, q := range questions {
		public[i] = q.Public()
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(public)
}

func (h *Handler) PostSubmit(w http.ResponseWriter, r *http.Request) {
	var req SubmitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	score := 0
	total := len(req.Answers)
	for _, a := range req.Answers {
		if h.store.IsCorrect(a.QuestionId, a.OptionId) {
			score++
		}
	}
	percentile, hasPercentile := h.store.Percentile(score)
	h.store.SaveResult(score)

	result := SubmitResult{
		Score:         score,
		Total:         total,
		Percentile:    percentile,
		HasPercentile: hasPercentile,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
