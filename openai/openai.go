package openai

import (
	"context"
	"fmt"
	"os"
	"errors"
	"io"
	"io/ioutil"
	"log"
	opai "github.com/sashabaranov/go-openai"
)


func SubmitPrompt(name string) {
	c := opai.NewClient(os.Getenv("OPEN_AI_KEY"))
	ctx := context.Background()
	f, err := os.Create(fmt.Sprintf("./outputs/hashi-assistant/%v.md", name))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	fileContent, err := ioutil.ReadFile("./prompts/hashicorp-assistant.md")
	if err != nil {
		log.Fatal(err)
	}

	prompt := string(fileContent)

	req := opai.ChatCompletionRequest{
		Model:     opai.GPT3Dot5Turbo,
		MaxTokens: 3500,
		Messages: []opai.ChatCompletionMessage{
			{
				Role:    opai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Stream: true,
	}

	stream, err := c.CreateChatCompletionStream(ctx, req)

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

		line := response.Choices[0].Delta.Content

		f.WriteString(line)
		fmt.Printf(line)

	}


}
