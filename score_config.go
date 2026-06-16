package langfuse

import (
	"context"
	"fmt"
	"time"
)

// ScoreConfigDataType is the data type for a score configuration.
type ScoreConfigDataType string

const (
	ScoreConfigDataTypeNumeric     ScoreConfigDataType = "NUMERIC"
	ScoreConfigDataTypeBoolean     ScoreConfigDataType = "BOOLEAN"
	ScoreConfigDataTypeCategorical ScoreConfigDataType = "CATEGORICAL"
	ScoreConfigDataTypeText        ScoreConfigDataType = "TEXT"
	ScoreConfigDataTypeCorrection  ScoreConfigDataType = "CORRECTION"
)

// ConfigCategory is a label-value pair for categorical score configs.
type ConfigCategory struct {
	Value float64 `json:"value"`
	Label string  `json:"label"`
}

// ScoreConfig represents a Langfuse score configuration.
type ScoreConfig struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	CreatedAt   time.Time           `json:"createdAt"`
	UpdatedAt   time.Time           `json:"updatedAt"`
	ProjectID   string              `json:"projectId"`
	DataType    ScoreConfigDataType `json:"dataType"`
	IsArchived  bool                `json:"isArchived"`
	MinValue    *float64            `json:"minValue,omitempty"`
	MaxValue    *float64            `json:"maxValue,omitempty"`
	Categories  []*ConfigCategory   `json:"categories,omitempty"`
	Description *string             `json:"description,omitempty"`
}

// GetScoreConfigsOutput is the response for listing score configs.
type GetScoreConfigsOutput struct {
	Data []*ScoreConfig `json:"data"`
	Meta *Pagination    `json:"meta"`
}

// GetScoreConfigsOptions are the query parameters for listing score configs.
type GetScoreConfigsOptions struct {
	Page  *int `url:"page,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

// GetScoreConfigs returns a paginated list of score configurations.
func (c *Client) GetScoreConfigs(ctx context.Context, opts *GetScoreConfigsOptions) (*GetScoreConfigsOutput, error) {
	path := "/api/public/score-configs"
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

	r := new(GetScoreConfigsOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetScoreConfig returns a single score configuration by ID.
func (c *Client) GetScoreConfig(ctx context.Context, configID string) (*ScoreConfig, error) {
	path := fmt.Sprintf("/api/public/score-configs/%s", configID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(ScoreConfig)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateScoreConfigInput is the request body for creating a score configuration.
type CreateScoreConfigInput struct {
	Name        string              `json:"name"`
	DataType    ScoreConfigDataType `json:"dataType"`
	Categories  []*ConfigCategory   `json:"categories,omitempty"`
	MinValue    *float64            `json:"minValue,omitempty"`
	MaxValue    *float64            `json:"maxValue,omitempty"`
	Description *string             `json:"description,omitempty"`
}

// CreateScoreConfig creates a new score configuration.
func (c *Client) CreateScoreConfig(ctx context.Context, input *CreateScoreConfigInput) (*ScoreConfig, error) {
	req, err := c.NewRequest("POST", "/api/public/score-configs", input)
	if err != nil {
		return nil, err
	}

	r := new(ScoreConfig)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// UpdateScoreConfigInput is the request body for updating a score configuration.
type UpdateScoreConfigInput struct {
	IsArchived  *bool             `json:"isArchived,omitempty"`
	Name        *string           `json:"name,omitempty"`
	Categories  []*ConfigCategory `json:"categories,omitempty"`
	MinValue    *float64          `json:"minValue,omitempty"`
	MaxValue    *float64          `json:"maxValue,omitempty"`
	Description *string           `json:"description,omitempty"`
}

// UpdateScoreConfig updates an existing score configuration.
func (c *Client) UpdateScoreConfig(
	ctx context.Context,
	configID string,
	input *UpdateScoreConfigInput,
) (*ScoreConfig, error) {
	path := fmt.Sprintf("/api/public/score-configs/%s", configID)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(ScoreConfig)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}
