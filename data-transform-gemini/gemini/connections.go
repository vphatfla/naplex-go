package gemini

import (
	"context"

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
