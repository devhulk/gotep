package main

import (
	"context"
	"fmt"
	"os"
	"errors"
	"io"
	openai "github.com/sashabaranov/go-openai"
)


func main() {
	c := openai.NewClient(os.Getenv("OPEN_AI_KEY"))
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3TextDavinci003,
		Temperature: 0.3,
		MaxTokens: 1620,
		TopP: 1,
		BestOf: 1,
		Prompt: 	"write golang leftpad function",
		Stream: true,
	}
	stream, err := c.CreateCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("CreateCompletion error: %v\n", err)
		return
	}
	defer stream.Close()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		fmt.Printf(response.Choices[0].Text)
	}

	}
