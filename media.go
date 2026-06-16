package langfuse

import (
	"context"
	"fmt"
	"time"
)

// Media represents a media record in Langfuse.
type Media struct {
	MediaID       string     `json:"mediaId"`
	ContentType   string     `json:"contentType"`
	ContentLength int64      `json:"contentLength"`
	Sha256Hash    string     `json:"sha256Hash"`
	UploadedAt    *time.Time `json:"uploadedAt,omitempty"`
	URL           *string    `json:"url,omitempty"`
	URLExpiry     *time.Time `json:"urlExpiry,omitempty"`
}

// GetMediaUploadURLInput is the request body for obtaining a presigned upload URL.
type GetMediaUploadURLInput struct {
	MimeType      string  `json:"mimeType"`
	ContentLength int64   `json:"contentLength"`
	Sha256Hash    string  `json:"sha256Hash"`
	Field         string  `json:"field"`
	TraceID       *string `json:"traceId,omitempty"`
	ObservationID *string `json:"observationId,omitempty"`
}

// GetMediaUploadURLOutput is the response containing the presigned URL for uploading media.
type GetMediaUploadURLOutput struct {
	MediaID      string            `json:"mediaId"`
	UploadURL    *string           `json:"uploadUrl,omitempty"`
	UploadFields map[string]string `json:"uploadFields,omitempty"`
	ContentType  string            `json:"contentType"`
}

// GetMediaUploadURL requests a presigned URL for uploading a media file.
func (c *Client) GetMediaUploadURL(ctx context.Context, input *GetMediaUploadURLInput) (*GetMediaUploadURLOutput, error) {
	req, err := c.NewRequest("POST", "/api/public/media", input)
	if err != nil {
		return nil, err
	}

	r := new(GetMediaUploadURLOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetMedia returns a media record by ID, including a short-lived download URL.
func (c *Client) GetMedia(ctx context.Context, mediaID string) (*Media, error) {
	path := fmt.Sprintf("/api/public/media/%s", mediaID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(Media)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// UpdateMediaInput is the request body for updating media metadata after upload.
type UpdateMediaInput struct {
	UploadedAt       string `json:"uploadedAt"`
	UploadHTTPStatus *int   `json:"uploadHttpStatus,omitempty"`
}

// UpdateMedia updates media metadata (e.g., marks the upload as completed).
func (c *Client) UpdateMedia(ctx context.Context, mediaID string, input *UpdateMediaInput) error {
	path := fmt.Sprintf("/api/public/media/%s", mediaID)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}
