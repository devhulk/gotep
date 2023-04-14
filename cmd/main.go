package main

import (
	"context"
	"fmt"
	"os"
	"errors"
	"io"
	"io/ioutil"
	"log"
	openai "github.com/sashabaranov/go-openai"
)


func main() {
	c := openai.NewClient(os.Getenv("OPEN_AI_KEY"))
	ctx := context.Background()
	f, err := os.Create("./outputs/aws-vm/output1.md")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	fileContent, err := ioutil.ReadFile("prompts/aws-vm.md")
	if err != nil {
		log.Fatal(err)
	}

	prompt := string(fileContent)

	req := openai.CompletionRequest{
		Model:     openai.GPT3TextDavinci003,
		Temperature: 0.3,
		MaxTokens: 1620,
		TopP: 1,
		BestOf: 1,
		Prompt: 	prompt,
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

		line := response.Choices[0].Text

		f.WriteString(line)
		fmt.Printf(line)

	}

}
