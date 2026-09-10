package api

import "net/http"

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /questions", h.GetQuestions)
	mux.HandleFunc("POST /submit", h.PostSubmit)
	return mux
}
