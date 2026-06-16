package langfuse

import "context"

// IngestionEventType is the type of a batch ingestion event.
type IngestionEventType string

const (
	IngestionEventTypeTraceCreate       IngestionEventType = "trace-create"
	IngestionEventTypeObservationCreate IngestionEventType = "observation-create"
	IngestionEventTypeObservationUpdate IngestionEventType = "observation-update"
	IngestionEventTypeScoreCreate       IngestionEventType = "score-create"
	IngestionEventTypeSDKLog            IngestionEventType = "sdk-log"
)

// IngestionEvent is a single event in a batch ingestion request.
type IngestionEvent struct {
	ID        string             `json:"id"`
	Type      IngestionEventType `json:"type"`
	Body      any                `json:"body"`
	Timestamp *string            `json:"timestamp,omitempty"` // ISO 8601
}

// BatchIngestionInput is the request body for batch ingestion.
type BatchIngestionInput struct {
	Batch    []*IngestionEvent `json:"batch"`
	Metadata any               `json:"metadata,omitempty"`
}

// IngestionSuccess represents a successfully processed event in the batch.
type IngestionSuccess struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
}

// IngestionError represents a failed event in the batch.
type IngestionError struct {
	ID      string `json:"id"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

// BatchIngestionOutput is the response from a batch ingestion request.
type BatchIngestionOutput struct {
	Successes []*IngestionSuccess `json:"successes"`
	Errors    []*IngestionError   `json:"errors"`
}

// BatchIngestion sends a batch of events to Langfuse for ingestion.
func (c *Client) BatchIngestion(ctx context.Context, input *BatchIngestionInput) (*BatchIngestionOutput, error) {
	req, err := c.NewRequest("POST", "/api/public/ingestion", input)
	if err != nil {
		return nil, err
	}

	r := new(BatchIngestionOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}
