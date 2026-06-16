package langfuse

import (
	"context"
	"fmt"
	"time"
)

// BlobStorageType is the type of blob storage provider.
type BlobStorageType string

const (
	BlobStorageTypeS3        BlobStorageType = "S3"
	BlobStorageTypeAzureBlob BlobStorageType = "AZURE_BLOB"
	BlobStorageTypeGCS       BlobStorageType = "GCS"
)

// BlobStorageIntegration represents a configured external blob storage integration.
type BlobStorageIntegration struct {
	ID           string          `json:"id"`
	Type         BlobStorageType `json:"type"`
	BucketName   string          `json:"bucketName"`
	Prefix       *string         `json:"prefix,omitempty"`
	Region       *string         `json:"region,omitempty"`
	Endpoint     *string         `json:"endpoint,omitempty"`
	ExportPrefix *string         `json:"exportPrefix,omitempty"`
	Enabled      bool            `json:"enabled"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
}

// GetBlobStorageIntegrationsOutput is the response for listing blob storage integrations.
type GetBlobStorageIntegrationsOutput struct {
	Data []*BlobStorageIntegration `json:"data"`
}

// GetBlobStorageIntegrations returns all configured blob storage integrations.
func (c *Client) GetBlobStorageIntegrations(ctx context.Context) (*GetBlobStorageIntegrationsOutput, error) {
	req, err := c.NewRequest("GET", "/api/public/integrations/blob-storage", nil)
	if err != nil {
		return nil, err
	}

	r := new(GetBlobStorageIntegrationsOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// UpsertBlobStorageIntegrationInput is the request body for creating or updating a blob storage integration.
type UpsertBlobStorageIntegrationInput struct {
	Type            BlobStorageType `json:"type"`
	BucketName      string          `json:"bucketName"`
	Prefix          *string         `json:"prefix,omitempty"`
	Region          *string         `json:"region,omitempty"`
	Endpoint        *string         `json:"endpoint,omitempty"`
	ExportPrefix    *string         `json:"exportPrefix,omitempty"`
	AccessKeyID     *string         `json:"accessKeyId,omitempty"`
	SecretAccessKey *string         `json:"secretAccessKey,omitempty"`
	Enabled         *bool           `json:"enabled,omitempty"`
}

// UpsertBlobStorageIntegration creates or updates a blob storage integration.
func (c *Client) UpsertBlobStorageIntegration(
	ctx context.Context, input *UpsertBlobStorageIntegrationInput,
) (*BlobStorageIntegration, error) {
	req, err := c.NewRequest("PUT", "/api/public/integrations/blob-storage", input)
	if err != nil {
		return nil, err
	}

	r := new(BlobStorageIntegration)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetBlobStorageIntegration returns a single blob storage integration by ID.
func (c *Client) GetBlobStorageIntegration(ctx context.Context, integrationID string) (*BlobStorageIntegration, error) {
	path := fmt.Sprintf("/api/public/integrations/blob-storage/%s", integrationID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(BlobStorageIntegration)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteBlobStorageIntegration deletes a blob storage integration by ID.
func (c *Client) DeleteBlobStorageIntegration(ctx context.Context, integrationID string) error {
	path := fmt.Sprintf("/api/public/integrations/blob-storage/%s", integrationID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}
