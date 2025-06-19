package question

// DTO stands for data transfer object
// DTO use to define a schema JSON structure for client's HTTP Response
type QuestionDTO struct {
	ID               int32    `json:"question_id"`
	Title            string   `json:"title"`
	Question         string   `json:"question"`
	Multiple_choices []string `json:"multiple_choices"`
	Correct_answer   string   `json:"correct_answer"`
	Explanation      string   `json:"explanation"`
	Keywords         []string `json:"keywords"`
	Link             string   `json:"link"`
	Status           string   `json:"status"`
	Attempts         int32    `json:"attempts"`
	Saved            bool     `json:"saved"`
	Hidden           bool     `json:"hidden"`
}
