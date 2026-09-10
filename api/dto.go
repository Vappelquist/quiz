package api

type SubmitRequest struct {
	Answers []Answer `json:"answers"`
}

type Answer struct {
	QuestionId string `json:questionId`
	OptionId   string `json:optionId`
}

type SubmitResult struct {
	Score         int     `json:"score"`
	Total         int     `json:"total"`
	Percentile    float64 `json:"percentile"`
	HasPercentile bool    `json:hasPercentile"`
}
