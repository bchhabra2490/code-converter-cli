package converter

import (
	"context"
	"fmt"
	"os"
	"strings"
	openai "github.com/sashabaranov/go-openai"
)

// Define GenerateText as a variable that holds a function
var GenerateText = generateTextImpl

// The actual implementation is now in this function
func generateTextImpl(prompt string) (string, error) {
	client := openai.NewClient(
		os.Getenv("OPENAI_API_KEY"),
	)
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate text: %w", err)
	}

	// clean the code.
	code := resp.Choices[0].Message.Content
	// Remove first and last lines of code.
	cleanedCode := removeFirstAndLastLines(code)
	fmt.Println(cleanedCode)
	return cleanedCode, nil
}

func removeFirstAndLastLines(code string) string {
	lines := strings.Split(code, "\n")
	
	// If there are fewer than 3 lines, return the original text
	// (need at least 3 lines to remove first and last and have content left)
	if len(lines) < 3 {
		return code
	}
	
	// Check if first line contains code block markers (```), if so remove it
	if strings.Contains(lines[0], "```") {
		lines = lines[1:]
	}
	
	// Check if last line contains code block markers (```), if so remove it
	lastIdx := len(lines) - 1
	if strings.Contains(lines[lastIdx], "```") {
		lines = lines[:lastIdx]
	}
	
	// Join the remaining lines back together
	return strings.Join(lines, "\n")
}