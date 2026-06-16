package langfuse

import (
	"context"
	"fmt"
	"time"
)

// LLMConnection represents a configured external LLM provider connection.
type LLMConnection struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Provider          string    `json:"provider"`
	BaseURL           *string   `json:"baseUrl,omitempty"`
	WithDefaultModels bool      `json:"withDefaultModels"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

// GetLLMConnectionsOutput is the response for listing LLM connections.
type GetLLMConnectionsOutput struct {
	Data []*LLMConnection `json:"data"`
	Meta *Pagination      `json:"meta"`
}

// GetLLMConnectionsOptions are the query parameters for listing LLM connections.
type GetLLMConnectionsOptions struct {
	Page  *int `url:"page,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

// GetLLMConnections returns a paginated list of LLM connections.
func (c *Client) GetLLMConnections(
	ctx context.Context, opts *GetLLMConnectionsOptions,
) (*GetLLMConnectionsOutput, error) {
	path := "/api/public/llm-connections"
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

	r := new(GetLLMConnectionsOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// UpsertLLMConnectionInput is the request body for creating or updating an LLM connection.
type UpsertLLMConnectionInput struct {
	Name              string  `json:"name"`
	Provider          string  `json:"provider"`
	BaseURL           *string `json:"baseUrl,omitempty"`
	APIKey            *string `json:"apiKey,omitempty"`
	WithDefaultModels *bool   `json:"withDefaultModels,omitempty"`
}

// UpsertLLMConnection creates or updates an LLM connection.
func (c *Client) UpsertLLMConnection(ctx context.Context, input *UpsertLLMConnectionInput) (*LLMConnection, error) {
	req, err := c.NewRequest("PUT", "/api/public/llm-connections", input)
	if err != nil {
		return nil, err
	}

	r := new(LLMConnection)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteLLMConnection deletes an LLM connection by ID.
func (c *Client) DeleteLLMConnection(ctx context.Context, connectionID string) error {
	path := fmt.Sprintf("/api/public/llm-connections/%s", connectionID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}
