package gemini

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"github.com/vphatfla/naplex-go/data-transform-gemini/config"
	"google.golang.org/api/option"
)

func NewClient(ctx context.Context, cfg *config.Config) (*genai.Client, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.Gemini.APIKEY))
	if err != nil {
		return nil, err
	}
	return client, err
}

func NewModelJson(client *genai.Client, cfg *config.Config) *genai.GenerativeModel {
	model := client.GenerativeModel(cfg.Gemini.Model)
	model.ResponseMIMEType = "application/json"

	return model
}

func GetContent(ctx context.Context, model *genai.GenerativeModel, rawQuestionTxt string) (string, error) {
	prompt := `Process this pharmacy exam question into a JSON object:
	- "title": Brief descriptive title (e.g., "Vancomycin Dosing")
	- "question": Full question text with patient case
	- "multipleChoices": Format as "A. [text] B. [text] C. [text] D. [text]"
	- "correctAnswer": Letter + text (e.g., "B. 1000 mg IV q12h")
	- "explanation": Rationale for correct answer
	- "keywords": Single string with 2-4 terms separated by commas only without spaces (e.g., "vancomycin,dosing,antimicrobial")
	Remove any '+' line endings and database artifacts.`+ rawQuestionTxt
	res, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	for _, c := range res.Candidates {
		if c != nil {
			for _, part := range c.Content.Parts {
				if txt, ok := part.(genai.Text); ok {
					return string(txt), nil
				} else {
					return "", fmt.Errorf("Error parsing string text from part candidate of model response")
				}
			}
		}
	}

	return "", fmt.Errorf("No available response in get content")
}
