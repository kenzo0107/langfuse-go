package langfuse

import (
	"context"
	"fmt"
	"time"
)

// ObservationType is the type of observation (generation, span, event).
type ObservationType string

const (
	ObservationTypeGeneration ObservationType = "GENERATION"
	ObservationTypeSpan       ObservationType = "SPAN"
	ObservationTypeEvent      ObservationType = "EVENT"
)

// ObservationLevel is the log level for an observation.
type ObservationLevel string

const (
	ObservationLevelDebug   ObservationLevel = "DEBUG"
	ObservationLevelDefault ObservationLevel = "DEFAULT"
	ObservationLevelWarning ObservationLevel = "WARNING"
	ObservationLevelError   ObservationLevel = "ERROR"
)

// Observation represents a generation, span, or event in a trace.
type Observation struct {
	ID                  string           `json:"id"`
	TraceID             string           `json:"traceId"`
	Type                ObservationType  `json:"type"`
	Name                *string          `json:"name,omitempty"`
	StartTime           time.Time        `json:"startTime"`
	EndTime             *time.Time       `json:"endTime,omitempty"`
	Model               *string          `json:"model,omitempty"`
	ModelParameters     any              `json:"modelParameters,omitempty"`
	Input               any              `json:"input,omitempty"`
	Output              any              `json:"output,omitempty"`
	Level               ObservationLevel `json:"level"`
	StatusMessage       *string          `json:"statusMessage,omitempty"`
	ParentObservationID *string          `json:"parentObservationId,omitempty"`
	PromptID            *string          `json:"promptId,omitempty"`
	PromptName          *string          `json:"promptName,omitempty"`
	PromptVersion       *int             `json:"promptVersion,omitempty"`
	Version             *string          `json:"version,omitempty"`
	Metadata            any              `json:"metadata,omitempty"`
	ProjectID           string           `json:"projectId"`
	CreatedAt           time.Time        `json:"createdAt"`
	UpdatedAt           time.Time        `json:"updatedAt"`
}

// GetObservationsOutput is the response for listing observations.
type GetObservationsOutput struct {
	Data []*Observation `json:"data"`
	Meta *Pagination    `json:"meta"`
}

// GetObservationsOptions are the query parameters for listing observations.
type GetObservationsOptions struct {
	Page          *int             `url:"page,omitempty"`
	Limit         *int             `url:"limit,omitempty"`
	Type          *ObservationType `url:"type,omitempty"`
	Name          *string          `url:"name,omitempty"`
	TraceID       *string          `url:"traceId,omitempty"`
	ParentObservationID *string    `url:"parentObservationId,omitempty"`
	FromStartTime *time.Time       `url:"fromStartTime,omitempty"`
	ToStartTime   *time.Time       `url:"toStartTime,omitempty"`
	Model         *string          `url:"model,omitempty"`
	Environment   *string          `url:"environment,omitempty"`
}

// GetObservations returns a paginated list of observations.
func (c *Client) GetObservations(ctx context.Context, opts *GetObservationsOptions) (*GetObservationsOutput, error) {
	path := "/api/public/v2/observations"
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

	r := new(GetObservationsOutput)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetObservation returns a single observation by ID.
func (c *Client) GetObservation(ctx context.Context, observationID string) (*Observation, error) {
	path := fmt.Sprintf("/api/public/v2/observations/%s", observationID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(Observation)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}
