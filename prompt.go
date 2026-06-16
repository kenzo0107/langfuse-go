package langfuse

import (
	"context"
	"fmt"
	"time"
)

// PromptType is either "text" or "chat".
type PromptType string

const (
	PromptTypeText PromptType = "text"
	PromptTypeChat PromptType = "chat"
)

// ChatMessage is a single message in a chat prompt.
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Prompt represents a resolved Langfuse prompt (text or chat).
type Prompt struct {
	Name            string     `json:"name"`
	Version         int        `json:"version"`
	Type            PromptType `json:"type"`
	Prompt          any        `json:"prompt"` // string for text, []ChatMessage for chat
	Config          any        `json:"config,omitempty"`
	Labels          []string   `json:"labels"`
	Tags            []string   `json:"tags"`
	CommitMessage   *string    `json:"commitMessage,omitempty"`
	ResolutionGraph any        `json:"resolutionGraph,omitempty"`
}

// PromptMeta holds summary metadata for a prompt (returned by list).
type PromptMeta struct {
	Name          string     `json:"name"`
	Type          PromptType `json:"type"`
	Versions      []int      `json:"versions"`
	Labels        []string   `json:"labels"`
	Tags          []string   `json:"tags"`
	LastUpdatedAt time.Time  `json:"lastUpdatedAt"`
	LastConfig    any        `json:"lastConfig"`
}

// GetPromptsOutput is the response for listing prompts.
type GetPromptsOutput struct {
	Data []*PromptMeta `json:"data"`
	Meta *Pagination   `json:"meta"`
}

// GetPromptsOptions are the query parameters for listing prompts.
type GetPromptsOptions struct {
	Name          *string    `url:"name,omitempty"`
	Label         *string    `url:"label,omitempty"`
	Tag           *string    `url:"tag,omitempty"`
	Page          *int       `url:"page,omitempty"`
	Limit         *int       `url:"limit,omitempty"`
	FromUpdatedAt *time.Time `url:"fromUpdatedAt,omitempty"`
	ToUpdatedAt   *time.Time `url:"toUpdatedAt,omitempty"`
}

// GetPrompts returns a paginated list of prompt metadata.
func (c *Client) GetPrompts(ctx context.Context, opts *GetPromptsOptions) (*GetPromptsOutput, error) {
	path := "/api/public/v2/prompts"
	if opts != nil {
		var err error
		path, err = c.AddOptions(path, opts)
		if err != nil {
			return nil, err
		}
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(GetPromptsOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetPromptOptions are the query parameters for fetching a single prompt.
type GetPromptOptions struct {
	Version *int    `url:"version,omitempty"`
	Label   *string `url:"label,omitempty"`
}

// GetPrompt returns a specific prompt by name, optionally at a given version or label.
func (c *Client) GetPrompt(ctx context.Context, name string, opts *GetPromptOptions) (*Prompt, error) {
	path := fmt.Sprintf("/api/public/v2/prompts/%s", name)
	if opts != nil {
		var err error
		path, err = c.AddOptions(path, opts)
		if err != nil {
			return nil, err
		}
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(Prompt)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateTextPromptInput is the request body for creating a text prompt.
type CreateTextPromptInput struct {
	Name          string   `json:"name"`
	Prompt        string   `json:"prompt"`
	Type          string   `json:"type,omitempty"` // defaults to "text"
	Labels        []string `json:"labels,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Config        any      `json:"config,omitempty"`
	CommitMessage *string  `json:"commitMessage,omitempty"`
}

// CreateTextPrompt creates a new text prompt (or a new version if the name already exists).
func (c *Client) CreateTextPrompt(ctx context.Context, input *CreateTextPromptInput) (*Prompt, error) {
	if input.Type == "" {
		input.Type = string(PromptTypeText)
	}

	req, err := c.NewRequest("POST", "/api/public/v2/prompts", input)
	if err != nil {
		return nil, err
	}

	r := new(Prompt)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateChatPromptInput is the request body for creating a chat prompt.
type CreateChatPromptInput struct {
	Name          string        `json:"name"`
	Prompt        []ChatMessage `json:"prompt"`
	Type          string        `json:"type"` // must be "chat"
	Labels        []string      `json:"labels,omitempty"`
	Tags          []string      `json:"tags,omitempty"`
	Config        any           `json:"config,omitempty"`
	CommitMessage *string       `json:"commitMessage,omitempty"`
}

// CreateChatPrompt creates a new chat prompt (or a new version if the name already exists).
func (c *Client) CreateChatPrompt(ctx context.Context, input *CreateChatPromptInput) (*Prompt, error) {
	input.Type = string(PromptTypeChat)

	req, err := c.NewRequest("POST", "/api/public/v2/prompts", input)
	if err != nil {
		return nil, err
	}

	r := new(Prompt)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeletePrompt deletes all versions of a prompt by name.
func (c *Client) DeletePrompt(ctx context.Context, promptName string) error {
	path := fmt.Sprintf("/api/public/v2/prompts/%s", promptName)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}

// UpdatePromptVersionInput is the request body for updating a prompt version's labels.
type UpdatePromptVersionInput struct {
	NewLabels []string `json:"newLabels"`
}

// UpdatePromptVersion updates the labels on a specific prompt version.
func (c *Client) UpdatePromptVersion(
	ctx context.Context, name string, version int, input *UpdatePromptVersionInput,
) (*Prompt, error) {
	path := fmt.Sprintf("/api/public/v2/prompts/%s/versions/%d", name, version)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(Prompt)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}
