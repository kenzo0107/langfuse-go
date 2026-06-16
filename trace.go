package langfuse

import (
	"context"
	"fmt"
	"time"
)

// Trace represents a Langfuse trace.
type Trace struct {
	ID          string    `json:"id"`
	Name        *string   `json:"name,omitempty"`
	Input       any       `json:"input,omitempty"`
	Output      any       `json:"output,omitempty"`
	Metadata    any       `json:"metadata,omitempty"`
	Tags        []string  `json:"tags"`
	Version     *string   `json:"version,omitempty"`
	Release     *string   `json:"release,omitempty"`
	UserID      *string   `json:"userId,omitempty"`
	SessionID   *string   `json:"sessionId,omitempty"`
	Environment string    `json:"environment"`
	ProjectID   string    `json:"projectId"`
	Public      bool      `json:"public"`
	Bookmarked  bool      `json:"bookmarked"`
	Timestamp   time.Time `json:"timestamp"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// GetTracesOutput is the response for listing traces.
type GetTracesOutput struct {
	Data []*Trace    `json:"data"`
	Meta *Pagination `json:"meta"`
}

// GetTracesOptions are the query parameters for listing traces.
type GetTracesOptions struct {
	Page          *int       `url:"page,omitempty"`
	Limit         *int       `url:"limit,omitempty"`
	UserID        *string    `url:"userId,omitempty"`
	Name          *string    `url:"name,omitempty"`
	SessionID     *string    `url:"sessionId,omitempty"`
	FromTimestamp *time.Time `url:"fromTimestamp,omitempty"`
	ToTimestamp   *time.Time `url:"toTimestamp,omitempty"`
	Version       *string    `url:"version,omitempty"`
	Release       *string    `url:"release,omitempty"`
	Environment   *string    `url:"environment,omitempty"`
	Tags          []string   `url:"tags,omitempty"`
}

// GetTraces returns a paginated list of traces.
func (c *Client) GetTraces(ctx context.Context, opts *GetTracesOptions) (*GetTracesOutput, error) {
	path := "/api/public/traces"
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

	r := new(GetTracesOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetTrace returns a single trace by ID.
func (c *Client) GetTrace(ctx context.Context, traceID string) (*Trace, error) {
	path := fmt.Sprintf("/api/public/traces/%s", traceID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(Trace)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteTrace deletes a single trace by ID.
func (c *Client) DeleteTrace(ctx context.Context, traceID string) error {
	path := fmt.Sprintf("/api/public/traces/%s", traceID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}

// DeleteTracesInput is the request body for bulk-deleting traces.
type DeleteTracesInput struct {
	TraceIDs []string `json:"traceIds"`
}

// DeleteTraces deletes multiple traces by ID.
func (c *Client) DeleteTraces(ctx context.Context, input *DeleteTracesInput) error {
	req, err := c.NewRequest("DELETE", "/api/public/traces", input)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}
