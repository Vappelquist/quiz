package quiz

type Option struct {
	Id   string `json:"id"`
	Text string `json:"text"`
} //This struct represents a single option for a quiz question, with an ID and text description correct answer or not.

type Question struct {
	Id        string
	Text      string
	Options   []Option
	CorrectId string
} //This struct represents a single quiz question, with an ID, text description, a list of options, and the ID of the correct option.

type PublicQuestion struct {
	Id      string   `json:"id"`
	Text    string   `json:"text"`
	Options []Option `json:"options"`
}

func (q Question) Public() PublicQuestion {
	return PublicQuestion{
		Id:      q.Id,
		Text:    q.Text,
		Options: q.Options,
	}
} //This method converts a Question struct into a PublicQuestion struct, which only includes the ID, text, and options, but not the correct answer. This is useful for sending questions to clients without revealing the correct answer.
