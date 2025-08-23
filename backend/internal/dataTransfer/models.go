package dataTransfer

type Question struct {
	Title           string `json:"title"`
	Question        string `json:"question"`
	MultipleChoices string `json:"multiple_choices" validate:"required"`
	CorrectAnswer   string `json:"correct_answer" validate:"required"`
	Explanation     string `json:"explanation,omitempty"`
	Keywords        string `json:"keywords,omitempty"`
	Link            string `json:"link,omitempty" validate:"omitempty,url"`
}
