package langfuse

import (
	"context"
	"fmt"
	"time"
)

// DatasetItemStatus indicates whether a dataset item is active or archived.
type DatasetItemStatus string

const (
	DatasetItemStatusActive   DatasetItemStatus = "ACTIVE"
	DatasetItemStatusArchived DatasetItemStatus = "ARCHIVED"
)

// Dataset represents a Langfuse dataset.
type Dataset struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Metadata    any       `json:"metadata,omitempty"`
	ProjectID   string    `json:"projectId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// GetDatasetsOutput is the response for listing datasets.
type GetDatasetsOutput struct {
	Data []*Dataset  `json:"data"`
	Meta *Pagination `json:"meta"`
}

// GetDatasetsOptions are the query parameters for listing datasets.
type GetDatasetsOptions struct {
	Page  *int `url:"page,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

// GetDatasets returns a paginated list of datasets.
func (c *Client) GetDatasets(ctx context.Context, opts *GetDatasetsOptions) (*GetDatasetsOutput, error) {
	path := "/api/public/v2/datasets"
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

	r := new(GetDatasetsOutput)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetDataset returns a single dataset by name.
func (c *Client) GetDataset(ctx context.Context, datasetName string) (*Dataset, error) {
	path := fmt.Sprintf("/api/public/v2/datasets/%s", datasetName)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(Dataset)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateDatasetInput is the request body for creating or updating a dataset.
type CreateDatasetInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Metadata    any     `json:"metadata,omitempty"`
}

// CreateDataset creates a new dataset (upserts by name).
func (c *Client) CreateDataset(ctx context.Context, input *CreateDatasetInput) (*Dataset, error) {
	req, err := c.NewRequest("POST", "/api/public/v2/datasets", input)
	if err != nil {
		return nil, err
	}

	r := new(Dataset)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DatasetItem represents a single item within a dataset.
type DatasetItem struct {
	ID                  string            `json:"id"`
	DatasetID           string            `json:"datasetId"`
	DatasetName         string            `json:"datasetName"`
	Input               any               `json:"input,omitempty"`
	ExpectedOutput      any               `json:"expectedOutput,omitempty"`
	Metadata            any               `json:"metadata,omitempty"`
	SourceTraceID       *string           `json:"sourceTraceId,omitempty"`
	SourceObservationID *string           `json:"sourceObservationId,omitempty"`
	Status              DatasetItemStatus `json:"status"`
	CreatedAt           time.Time         `json:"createdAt"`
	UpdatedAt           time.Time         `json:"updatedAt"`
}

// GetDatasetItemsOutput is the response for listing dataset items.
type GetDatasetItemsOutput struct {
	Data []*DatasetItem `json:"data"`
	Meta *Pagination    `json:"meta"`
}

// GetDatasetItemsOptions are the query parameters for listing dataset items.
type GetDatasetItemsOptions struct {
	DatasetName *string `url:"datasetName,omitempty"`
	SourceTraceID       *string `url:"sourceTraceId,omitempty"`
	SourceObservationID *string `url:"sourceObservationId,omitempty"`
	Page                *int    `url:"page,omitempty"`
	Limit               *int    `url:"limit,omitempty"`
}

// GetDatasetItems returns a paginated list of dataset items.
func (c *Client) GetDatasetItems(ctx context.Context, opts *GetDatasetItemsOptions) (*GetDatasetItemsOutput, error) {
	path := "/api/public/dataset-items"
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

	r := new(GetDatasetItemsOutput)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetDatasetItem returns a single dataset item by ID.
func (c *Client) GetDatasetItem(ctx context.Context, itemID string) (*DatasetItem, error) {
	path := fmt.Sprintf("/api/public/dataset-items/%s", itemID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(DatasetItem)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateDatasetItemInput is the request body for creating or upserting a dataset item.
type CreateDatasetItemInput struct {
	DatasetName         string            `json:"datasetName"`
	Input               any               `json:"input,omitempty"`
	ExpectedOutput      any               `json:"expectedOutput,omitempty"`
	Metadata            any               `json:"metadata,omitempty"`
	SourceTraceID       *string           `json:"sourceTraceId,omitempty"`
	SourceObservationID *string           `json:"sourceObservationId,omitempty"`
	Status              DatasetItemStatus `json:"status,omitempty"`
	ID                  *string           `json:"id,omitempty"`
}

// CreateDatasetItem creates or upserts a dataset item.
func (c *Client) CreateDatasetItem(ctx context.Context, input *CreateDatasetItemInput) (*DatasetItem, error) {
	req, err := c.NewRequest("POST", "/api/public/dataset-items", input)
	if err != nil {
		return nil, err
	}

	r := new(DatasetItem)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteDatasetItem deletes a dataset item by ID.
func (c *Client) DeleteDatasetItem(ctx context.Context, itemID string) error {
	path := fmt.Sprintf("/api/public/dataset-items/%s", itemID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}
