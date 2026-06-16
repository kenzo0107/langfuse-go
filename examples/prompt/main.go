package main

import (
	"context"
	"fmt"
	"log"

	langfuse "github.com/kenzo0107/langfuse-go"
)

func main() {
	// Reads LANGFUSE_PUBLIC_KEY, LANGFUSE_SECRET_KEY, LANGFUSE_HOST from env.
	c := langfuse.NewFromEnv(langfuse.OptionDebug(true))

	ctx := context.Background()

	// Create a text prompt.
	created, err := c.CreateTextPrompt(ctx, &langfuse.CreateTextPromptInput{
		Name:   "my-summary-prompt",
		Prompt: "Summarize the following text: {{text}}",
		Labels: []string{"production"},
		Tags:   []string{"summarization"},
		Config: map[string]any{
			"temperature": 0.7,
			"maxTokens":   500,
		},
	})
	if err != nil {
		log.Fatalf("CreateTextPrompt: %v", err)
	}
	fmt.Printf("Created prompt: name=%s version=%d\n", created.Name, created.Version)

	// Get the prompt.
	got, err := c.GetPrompt(ctx, "my-summary-prompt", nil)
	if err != nil {
		log.Fatalf("GetPrompt: %v", err)
	}
	fmt.Printf("Got prompt: name=%s version=%d type=%s\n", got.Name, got.Version, got.Type)

	// List prompts.
	list, err := c.GetPrompts(ctx, &langfuse.GetPromptsOptions{
		Limit: langfuse.Int(10),
	})
	if err != nil {
		log.Fatalf("GetPrompts: %v", err)
	}
	fmt.Printf("Total prompts: %d (page %d/%d)\n", list.Meta.TotalItems, list.Meta.Page, list.Meta.TotalPages)
	for _, p := range list.Data {
		fmt.Printf("  - %s (type=%s, versions=%v)\n", p.Name, p.Type, p.Versions)
	}

	// Create a chat prompt.
	chatPrompt, err := c.CreateChatPrompt(ctx, &langfuse.CreateChatPromptInput{
		Name: "my-chat-prompt",
		Prompt: []langfuse.ChatMessage{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: "{{user_message}}"},
		},
		Labels: []string{"latest"},
	})
	if err != nil {
		log.Fatalf("CreateChatPrompt: %v", err)
	}
	fmt.Printf("Created chat prompt: name=%s version=%d\n", chatPrompt.Name, chatPrompt.Version)
}
