package langfuse

import "context"

// HealthResponse is the response from the health check endpoint.
type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version,omitempty"`
}

// GetHealth checks the health of the Langfuse API and database.
func (c *Client) GetHealth(ctx context.Context) (*HealthResponse, error) {
	req, err := c.NewRequest("GET", "/api/public/health", nil)
	if err != nil {
		return nil, err
	}

	r := new(HealthResponse)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}
