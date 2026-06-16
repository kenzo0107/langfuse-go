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
	if err = c.Do(ctx, req, r); err != nil {
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
	if err = c.Do(ctx, req, r); err != nil {
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
	if err = c.Do(ctx, req, r); err != nil {
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
	DatasetName         *string `url:"datasetName,omitempty"`
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
	if err = c.Do(ctx, req, r); err != nil {
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
	if err = c.Do(ctx, req, r); err != nil {
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
	if err = c.Do(ctx, req, r); err != nil {
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

// DatasetRun represents a named evaluation run over a dataset.
type DatasetRun struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Metadata    any       `json:"metadata,omitempty"`
	DatasetID   string    `json:"datasetId"`
	DatasetName string    `json:"datasetName"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// GetDatasetRunsOutput is the response for listing dataset runs.
type GetDatasetRunsOutput struct {
	Data []*DatasetRun `json:"data"`
	Meta *Pagination   `json:"meta"`
}

// GetDatasetRunsOptions are the query parameters for listing dataset runs.
type GetDatasetRunsOptions struct {
	Page  *int `url:"page,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

// GetDatasetRuns returns a paginated list of runs for a dataset.
func (c *Client) GetDatasetRuns(ctx context.Context, datasetName string, opts *GetDatasetRunsOptions) (*GetDatasetRunsOutput, error) {
	path := fmt.Sprintf("/api/public/datasets/%s/runs", datasetName)
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

	r := new(GetDatasetRunsOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetDatasetRun returns a single dataset run by name.
func (c *Client) GetDatasetRun(ctx context.Context, datasetName, runName string) (*DatasetRun, error) {
	path := fmt.Sprintf("/api/public/datasets/%s/runs/%s", datasetName, runName)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(DatasetRun)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteDatasetRun permanently deletes a dataset run and all its items.
func (c *Client) DeleteDatasetRun(ctx context.Context, datasetName, runName string) error {
	path := fmt.Sprintf("/api/public/datasets/%s/runs/%s", datasetName, runName)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}

// DatasetRunItem represents a single item within a dataset run.
type DatasetRunItem struct {
	ID             string    `json:"id"`
	DatasetRunID   string    `json:"datasetRunId"`
	DatasetRunName string    `json:"datasetRunName"`
	DatasetItemID  string    `json:"datasetItemId"`
	TraceID        string    `json:"traceId"`
	ObservationID  *string   `json:"observationId,omitempty"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// GetDatasetRunItemsOutput is the response for listing dataset run items.
type GetDatasetRunItemsOutput struct {
	Data []*DatasetRunItem `json:"data"`
	Meta *Pagination       `json:"meta"`
}

// GetDatasetRunItemsOptions are the query parameters for listing dataset run items.
type GetDatasetRunItemsOptions struct {
	Page          *string `url:"page,omitempty"`
	Limit         *string `url:"limit,omitempty"`
	DatasetRunID  *string `url:"datasetRunId,omitempty"`
	DatasetItemID *string `url:"datasetItemId,omitempty"`
}

// GetDatasetRunItems returns a paginated list of dataset run items.
func (c *Client) GetDatasetRunItems(ctx context.Context, opts *GetDatasetRunItemsOptions) (*GetDatasetRunItemsOutput, error) {
	path := "/api/public/dataset-run-items"
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

	r := new(GetDatasetRunItemsOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateDatasetRunItemInput is the request body for linking a trace to a dataset run.
type CreateDatasetRunItemInput struct {
	DatasetItemID  string  `json:"datasetItemId"`
	DatasetRunName string  `json:"datasetRunName"`
	TraceID        string  `json:"traceId"`
	ObservationID  *string `json:"observationId,omitempty"`
	RunDescription *string `json:"runDescription,omitempty"`
	Metadata       any     `json:"metadata,omitempty"`
}

// CreateDatasetRunItem links a trace (or observation) to a dataset run item.
func (c *Client) CreateDatasetRunItem(ctx context.Context, input *CreateDatasetRunItemInput) (*DatasetRunItem, error) {
	req, err := c.NewRequest("POST", "/api/public/dataset-run-items", input)
	if err != nil {
		return nil, err
	}

	r := new(DatasetRunItem)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}
