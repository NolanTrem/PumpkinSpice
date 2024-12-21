package oai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
)

func (c *Client) CreateCompletion(ctx context.Context, prompt string) (string, error) {
	resp, err := c.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		Model: openai.F(openai.ChatModelGPT4o2024_11_20),
	})
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (c *Client) StreamCompletion(ctx context.Context, prompt string) (<-chan string, error) {
	stream := c.client.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		Model: openai.F(openai.ChatModelGPT4o2024_11_20),
	})

	out := make(chan string)

	go func() {
		defer close(out)
		for stream.Next() {
			chunk := stream.Current()
			if len(chunk.Choices) > 0 {
				out <- chunk.Choices[0].Delta.Content
			}
		}

		if err := stream.Err(); err != nil {
			out <- fmt.Sprintf("Error: %v", err)
		}
	}()

	return out, nil
}
