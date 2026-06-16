package langfuse

import (
	"context"
	"fmt"
	"time"
)

// ScoreDataType is the data type of a score value.
type ScoreDataType string

const (
	ScoreDataTypeNumeric     ScoreDataType = "NUMERIC"
	ScoreDataTypeBoolean     ScoreDataType = "BOOLEAN"
	ScoreDataTypeCategorical ScoreDataType = "CATEGORICAL"
	ScoreDataTypeText        ScoreDataType = "TEXT"
	ScoreDataTypeCorrection  ScoreDataType = "CORRECTION"
)

// ScoreSource indicates the origin of a score.
type ScoreSource string

const (
	ScoreSourceAPI        ScoreSource = "API"
	ScoreSourceAnnotation ScoreSource = "ANNOTATION"
	ScoreSourceEval       ScoreSource = "EVAL"
)

// Score represents a Langfuse score.
type Score struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Source        ScoreSource    `json:"source"`
	DataType      *ScoreDataType `json:"dataType,omitempty"`
	Value         *float64       `json:"value,omitempty"`
	StringValue   *string        `json:"stringValue,omitempty"`
	TraceID       *string        `json:"traceId,omitempty"`
	SessionID     *string        `json:"sessionId,omitempty"`
	ObservationID *string        `json:"observationId,omitempty"`
	DatasetRunID  *string        `json:"datasetRunId,omitempty"`
	Comment       *string        `json:"comment,omitempty"`
	Metadata      any            `json:"metadata,omitempty"`
	ConfigID      *string        `json:"configId,omitempty"`
	QueueID       *string        `json:"queueId,omitempty"`
	AuthorUserID  *string        `json:"authorUserId,omitempty"`
	Environment   string         `json:"environment"`
	Timestamp     time.Time      `json:"timestamp"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
}

// CreateScoreOutput is the response after creating a score.
type CreateScoreOutput struct {
	ID string `json:"id"`
}

// CreateScoreInput is the request body for creating a score.
type CreateScoreInput struct {
	Name          string         `json:"name"`
	Value         any            `json:"value"` // float64 for numeric/boolean, string for categorical/text/correction
	TraceID       *string        `json:"traceId,omitempty"`
	SessionID     *string        `json:"sessionId,omitempty"`
	ObservationID *string        `json:"observationId,omitempty"`
	DatasetRunID  *string        `json:"datasetRunId,omitempty"`
	Comment       *string        `json:"comment,omitempty"`
	Metadata      any            `json:"metadata,omitempty"`
	Environment   *string        `json:"environment,omitempty"`
	DataType      *ScoreDataType `json:"dataType,omitempty"`
	ConfigID      *string        `json:"configId,omitempty"`
	QueueID       *string        `json:"queueId,omitempty"`
	ID            *string        `json:"id,omitempty"`
}

// CreateScore creates a new score.
func (c *Client) CreateScore(ctx context.Context, input *CreateScoreInput) (*CreateScoreOutput, error) {
	req, err := c.NewRequest("POST", "/api/public/scores", input)
	if err != nil {
		return nil, err
	}

	r := new(CreateScoreOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetScoresOutput is the response for listing scores (v3).
type GetScoresOutput struct {
	Data []*Score    `json:"data"`
	Meta *Pagination `json:"meta"`
}

// GetScoresOptions are the query parameters for listing scores.
type GetScoresOptions struct {
	Page          *int       `url:"page,omitempty"`
	Limit         *int       `url:"limit,omitempty"`
	UserID        *string    `url:"userId,omitempty"`
	Name          *string    `url:"name,omitempty"`
	FromTimestamp *time.Time `url:"fromTimestamp,omitempty"`
	ToTimestamp   *time.Time `url:"toTimestamp,omitempty"`
	Source        *string    `url:"source,omitempty"`
	ConfigID      *string    `url:"configId,omitempty"`
	TraceID       *string    `url:"traceId,omitempty"`
	SessionID     *string    `url:"sessionId,omitempty"`
	ObservationID *string    `url:"observationId,omitempty"`
	DataType      *string    `url:"dataType,omitempty"`
}

// GetScores returns a paginated list of scores.
func (c *Client) GetScores(ctx context.Context, opts *GetScoresOptions) (*GetScoresOutput, error) {
	path := "/api/public/scores"
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

	r := new(GetScoresOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetScore returns a single score by ID.
func (c *Client) GetScore(ctx context.Context, scoreID string) (*Score, error) {
	path := fmt.Sprintf("/api/public/scores/%s", scoreID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(Score)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteScore deletes a score by ID.
func (c *Client) DeleteScore(ctx context.Context, scoreID string) error {
	path := fmt.Sprintf("/api/public/scores/%s", scoreID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}
