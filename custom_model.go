package langfuse

import (
	"context"
	"fmt"
	"time"
)

// ModelUsageUnit is the unit for model usage pricing.
type ModelUsageUnit string

const (
	ModelUsageUnitTokens ModelUsageUnit = "TOKENS"
)

// CustomModel represents a custom model definition in Langfuse.
type CustomModel struct {
	ID                string         `json:"id"`
	ModelName         string         `json:"modelName"`
	MatchPattern      string         `json:"matchPattern"`
	StartDate         *time.Time     `json:"startDate,omitempty"`
	Unit              *ModelUsageUnit `json:"unit,omitempty"`
	InputPrice        *float64       `json:"inputPrice,omitempty"`
	OutputPrice       *float64       `json:"outputPrice,omitempty"`
	TotalPrice        *float64       `json:"totalPrice,omitempty"`
	TokenizerID       *string        `json:"tokenizerId,omitempty"`
	TokenizerConfig   any            `json:"tokenizerConfig,omitempty"`
	IsLangfuseManaged bool           `json:"isLangfuseManaged"`
	CreatedAt         time.Time      `json:"createdAt"`
}

// GetCustomModelsOutput is the response for listing custom models.
type GetCustomModelsOutput struct {
	Data []*CustomModel `json:"data"`
	Meta *Pagination    `json:"meta"`
}

// GetCustomModelsOptions are the query parameters for listing models.
type GetCustomModelsOptions struct {
	Page  *int `url:"page,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

// GetCustomModels returns a paginated list of custom models.
func (c *Client) GetCustomModels(ctx context.Context, opts *GetCustomModelsOptions) (*GetCustomModelsOutput, error) {
	path := "/api/public/models"
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

	r := new(GetCustomModelsOutput)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetCustomModel returns a single custom model by ID.
func (c *Client) GetCustomModel(ctx context.Context, modelID string) (*CustomModel, error) {
	path := fmt.Sprintf("/api/public/models/%s", modelID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(CustomModel)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateCustomModelInput is the request body for creating a custom model.
type CreateCustomModelInput struct {
	ModelName       string          `json:"modelName"`
	MatchPattern    string          `json:"matchPattern"`
	StartDate       *time.Time      `json:"startDate,omitempty"`
	Unit            *ModelUsageUnit `json:"unit,omitempty"`
	InputPrice      *float64        `json:"inputPrice,omitempty"`
	OutputPrice     *float64        `json:"outputPrice,omitempty"`
	TotalPrice      *float64        `json:"totalPrice,omitempty"`
	TokenizerID     *string         `json:"tokenizerId,omitempty"`
	TokenizerConfig any             `json:"tokenizerConfig,omitempty"`
}

// CreateCustomModel creates a new custom model definition.
func (c *Client) CreateCustomModel(ctx context.Context, input *CreateCustomModelInput) (*CustomModel, error) {
	req, err := c.NewRequest("POST", "/api/public/models", input)
	if err != nil {
		return nil, err
	}

	r := new(CustomModel)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteCustomModel deletes a custom model by ID.
func (c *Client) DeleteCustomModel(ctx context.Context, modelID string) error {
	path := fmt.Sprintf("/api/public/models/%s", modelID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}
