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
	prompt := `Process this pharmacy exam question into a SINGLE JSON object:
	{
		"question": "The complete question text including case details and patient data",
		"multipleChoices": "All options formatted exactly as 'A. [text] B. [text] C. [text] D. [text]'",
		"correctAnswer": "The correct answer letter and text",
		"explanation": "FULL explanation including rationales for ALL answers (correct and incorrect)",
		"keywords": "term1,term2,term3,term4"
	}

	IMPORTANT:
	1. Include ALL explanations for both correct AND incorrect answers in the explanation field
	2. Return ONLY this single JSON object, not wrapped in an array
	3. For keywords, provide a comma-separated string without spaces (e.g., "acromegaly,octreotide,growthhormone")
	4. Do not omit any details from the original question text
	Remove any '+' line endings and database artifacts.` + rawQuestionTxt
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
