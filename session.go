package langfuse

import (
	"context"
	"fmt"
	"time"
)

// Session represents a Langfuse session summary.
type Session struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	ProjectID string    `json:"projectId"`
}

// SessionDetail includes the traces associated with a session.
type SessionDetail struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	ProjectID string    `json:"projectId"`
	Traces    []*Trace  `json:"traces"`
}

// GetSessionsOutput is the response for listing sessions.
type GetSessionsOutput struct {
	Data []*Session  `json:"data"`
	Meta *Pagination `json:"meta"`
}

// GetSessionsOptions are the query parameters for listing sessions.
type GetSessionsOptions struct {
	Page          *int       `url:"page,omitempty"`
	Limit         *int       `url:"limit,omitempty"`
	FromTimestamp *time.Time `url:"fromTimestamp,omitempty"`
	ToTimestamp   *time.Time `url:"toTimestamp,omitempty"`
	Environment   *string    `url:"environment,omitempty"`
}

// GetSessions returns a paginated list of sessions.
func (c *Client) GetSessions(ctx context.Context, opts *GetSessionsOptions) (*GetSessionsOutput, error) {
	path := "/api/public/sessions"
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

	r := new(GetSessionsOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetSession returns a single session with its associated traces.
func (c *Client) GetSession(ctx context.Context, sessionID string) (*SessionDetail, error) {
	path := fmt.Sprintf("/api/public/sessions/%s", sessionID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(SessionDetail)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}
