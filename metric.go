package langfuse

import (
	"context"
	"time"
)

// GetMetricsOptions are the query parameters for the v1 metrics endpoint.
type GetMetricsOptions struct {
	GroupBy     *string    `url:"groupBy,omitempty"` // auto | daily | weekly
	From        *time.Time `url:"from,omitempty"`
	To          *time.Time `url:"to,omitempty"`
	TraceName   *string    `url:"traceName,omitempty"`
	UserID      *string    `url:"userId,omitempty"`
	Tags        []string   `url:"tags,omitempty"`
	Environment *string    `url:"environment,omitempty"`
}

// GetMetrics returns aggregated usage and cost metrics.
func (c *Client) GetMetrics(ctx context.Context, opts *GetMetricsOptions) (any, error) {
	path := "/api/public/metrics"
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

	var r any
	if err = c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetMetricsV2Options are the query parameters for the v2 metrics endpoint.
type GetMetricsV2Options struct {
	Page        *int       `url:"page,omitempty"`
	Limit       *int       `url:"limit,omitempty"`
	GroupBy     *string    `url:"groupBy,omitempty"`
	From        *time.Time `url:"from,omitempty"`
	To          *time.Time `url:"to,omitempty"`
	TraceName   *string    `url:"traceName,omitempty"`
	UserID      *string    `url:"userId,omitempty"`
	Tags        []string   `url:"tags,omitempty"`
	Environment *string    `url:"environment,omitempty"`
}

// GetMetricsV2 returns aggregated usage and cost metrics (v2).
func (c *Client) GetMetricsV2(ctx context.Context, opts *GetMetricsV2Options) (any, error) {
	path := "/api/public/v2/metrics"
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

	var r any
	if err = c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}
