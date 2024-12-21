// internal/ai/chat.go
package ai

import (
	"context"

	"github.com/openai/openai-go"
)

func (c *Client) CreateCompletion(ctx context.Context, prompt string) (string, error) {
	resp, err := c.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
