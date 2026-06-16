package langfuse

import (
	"context"
	"fmt"
	"time"
)

// AnnotationQueue represents a Langfuse annotation queue.
type AnnotationQueue struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    *string   `json:"description,omitempty"`
	ScoreConfigIDs []string  `json:"scoreConfigIds"`
	ProjectID      string    `json:"projectId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// GetAnnotationQueuesOutput is the response for listing annotation queues.
type GetAnnotationQueuesOutput struct {
	Data []*AnnotationQueue `json:"data"`
	Meta *Pagination        `json:"meta"`
}

// GetAnnotationQueuesOptions are the query parameters for listing annotation queues.
type GetAnnotationQueuesOptions struct {
	Page  *int `url:"page,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

// GetAnnotationQueues returns a paginated list of annotation queues.
func (c *Client) GetAnnotationQueues(
	ctx context.Context,
	opts *GetAnnotationQueuesOptions,
) (*GetAnnotationQueuesOutput, error) {
	path := "/api/public/annotation-queues"
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

	r := new(GetAnnotationQueuesOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetAnnotationQueue returns a single annotation queue by ID.
func (c *Client) GetAnnotationQueue(ctx context.Context, queueID string) (*AnnotationQueue, error) {
	path := fmt.Sprintf("/api/public/annotation-queues/%s", queueID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(AnnotationQueue)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateAnnotationQueueInput is the request body for creating an annotation queue.
type CreateAnnotationQueueInput struct {
	Name           string   `json:"name"`
	Description    *string  `json:"description,omitempty"`
	ScoreConfigIDs []string `json:"scoreConfigIds,omitempty"`
}

// CreateAnnotationQueue creates a new annotation queue.
func (c *Client) CreateAnnotationQueue(
	ctx context.Context,
	input *CreateAnnotationQueueInput,
) (*AnnotationQueue, error) {
	req, err := c.NewRequest("POST", "/api/public/annotation-queues", input)
	if err != nil {
		return nil, err
	}

	r := new(AnnotationQueue)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// UpdateAnnotationQueueInput is the request body for updating an annotation queue.
type UpdateAnnotationQueueInput struct {
	Name           *string  `json:"name,omitempty"`
	Description    *string  `json:"description,omitempty"`
	ScoreConfigIDs []string `json:"scoreConfigIds,omitempty"`
}

// UpdateAnnotationQueue updates an existing annotation queue.
func (c *Client) UpdateAnnotationQueue(
	ctx context.Context,
	queueID string,
	input *UpdateAnnotationQueueInput,
) (*AnnotationQueue, error) {
	path := fmt.Sprintf("/api/public/annotation-queues/%s", queueID)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(AnnotationQueue)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteAnnotationQueue deletes an annotation queue by ID.
func (c *Client) DeleteAnnotationQueue(ctx context.Context, queueID string) error {
	path := fmt.Sprintf("/api/public/annotation-queues/%s", queueID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}

// AnnotationQueueItem represents an item in an annotation queue.
type AnnotationQueueItem struct {
	ID            string    `json:"id"`
	QueueID       string    `json:"queueId"`
	TraceID       string    `json:"traceId"`
	ObservationID *string   `json:"observationId,omitempty"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// GetAnnotationQueueItemsOutput is the response for listing queue items.
type GetAnnotationQueueItemsOutput struct {
	Data []*AnnotationQueueItem `json:"data"`
	Meta *Pagination            `json:"meta"`
}

// GetAnnotationQueueItemsOptions are the query parameters for listing queue items.
type GetAnnotationQueueItemsOptions struct {
	Page   *int    `url:"page,omitempty"`
	Limit  *int    `url:"limit,omitempty"`
	Status *string `url:"status,omitempty"`
}

// GetAnnotationQueueItems returns items in an annotation queue.
func (c *Client) GetAnnotationQueueItems(
	ctx context.Context,
	queueID string,
	opts *GetAnnotationQueueItemsOptions,
) (*GetAnnotationQueueItemsOutput, error) {
	path := fmt.Sprintf("/api/public/annotation-queues/%s/items", queueID)
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

	r := new(GetAnnotationQueueItemsOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateAnnotationQueueItemInput is the request body for adding an item to a queue.
type CreateAnnotationQueueItemInput struct {
	TraceID       string  `json:"traceId"`
	ObservationID *string `json:"observationId,omitempty"`
}

// CreateAnnotationQueueItem adds a trace/observation to an annotation queue.
func (c *Client) CreateAnnotationQueueItem(
	ctx context.Context,
	queueID string,
	input *CreateAnnotationQueueItemInput,
) (*AnnotationQueueItem, error) {
	path := fmt.Sprintf("/api/public/annotation-queues/%s/items", queueID)

	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(AnnotationQueueItem)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteAnnotationQueueItem removes an item from an annotation queue.
func (c *Client) DeleteAnnotationQueueItem(ctx context.Context, queueID, itemID string) error {
	path := fmt.Sprintf("/api/public/annotation-queues/%s/items/%s", queueID, itemID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}

// GetAnnotationQueueItem returns a single queue item by ID.
func (c *Client) GetAnnotationQueueItem(ctx context.Context, queueID, itemID string) (*AnnotationQueueItem, error) {
	path := fmt.Sprintf("/api/public/annotation-queues/%s/items/%s", queueID, itemID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(AnnotationQueueItem)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// UpdateAnnotationQueueItemInput is the request body for updating a queue item's status.
type UpdateAnnotationQueueItemInput struct {
	Status string `json:"status"` // QUEUED | ACTIVE | COMPLETED
}

// UpdateAnnotationQueueItem updates the status of a queue item.
func (c *Client) UpdateAnnotationQueueItem(
	ctx context.Context,
	queueID, itemID string,
	input *UpdateAnnotationQueueItemInput,
) (*AnnotationQueueItem, error) {
	path := fmt.Sprintf("/api/public/annotation-queues/%s/items/%s", queueID, itemID)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(AnnotationQueueItem)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// AnnotationQueueAssignment represents a user assignment to an annotation queue.
type AnnotationQueueAssignment struct {
	ID        string    `json:"id"`
	QueueID   string    `json:"queueId"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}

// CreateAnnotationQueueAssignmentInput is the request body for assigning a user to a queue.
type CreateAnnotationQueueAssignmentInput struct {
	UserID string `json:"userId"`
}

// CreateAnnotationQueueAssignment assigns a user to an annotation queue.
func (c *Client) CreateAnnotationQueueAssignment(
	ctx context.Context,
	queueID string,
	input *CreateAnnotationQueueAssignmentInput,
) (*AnnotationQueueAssignment, error) {
	path := fmt.Sprintf("/api/public/annotation-queues/%s/assignments", queueID)

	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(AnnotationQueueAssignment)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteAnnotationQueueAssignmentInput is the request body for removing a user assignment.
type DeleteAnnotationQueueAssignmentInput struct {
	UserID string `json:"userId"`
}

// DeleteAnnotationQueueAssignment removes a user assignment from an annotation queue.
func (c *Client) DeleteAnnotationQueueAssignment(
	ctx context.Context,
	queueID string,
	input *DeleteAnnotationQueueAssignmentInput,
) error {
	path := fmt.Sprintf("/api/public/annotation-queues/%s/assignments", queueID)

	req, err := c.NewRequest("DELETE", path, input)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}
