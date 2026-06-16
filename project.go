package langfuse

import (
	"context"
	"fmt"
	"time"
)

// Project represents a Langfuse project.
type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// GetProjectsOutput is the response for listing projects.
type GetProjectsOutput struct {
	Data []*Project  `json:"data"`
	Meta *Pagination `json:"meta"`
}

// GetProjects returns all projects accessible with the current API key.
func (c *Client) GetProjects(ctx context.Context) (*GetProjectsOutput, error) {
	req, err := c.NewRequest("GET", "/api/public/projects", nil)
	if err != nil {
		return nil, err
	}

	r := new(GetProjectsOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateProjectInput is the request body for creating a project (requires org-scoped key).
type CreateProjectInput struct {
	Name string `json:"name"`
}

// CreateProject creates a new project within the organization.
func (c *Client) CreateProject(ctx context.Context, input *CreateProjectInput) (*Project, error) {
	req, err := c.NewRequest("POST", "/api/public/projects", input)
	if err != nil {
		return nil, err
	}

	r := new(Project)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// UpdateProjectInput is the request body for updating a project.
type UpdateProjectInput struct {
	Name *string `json:"name,omitempty"`
}

// UpdateProject updates an existing project (requires org-scoped key).
func (c *Client) UpdateProject(ctx context.Context, projectID string, input *UpdateProjectInput) (*Project, error) {
	path := fmt.Sprintf("/api/public/projects/%s", projectID)

	req, err := c.NewRequest("PUT", path, input)
	if err != nil {
		return nil, err
	}

	r := new(Project)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteProject deletes a project by ID (requires org-scoped key).
func (c *Client) DeleteProject(ctx context.Context, projectID string) error {
	path := fmt.Sprintf("/api/public/projects/%s", projectID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}

// GetOrganizationProjects returns all projects within the organization.
func (c *Client) GetOrganizationProjects(ctx context.Context) (*GetProjectsOutput, error) {
	req, err := c.NewRequest("GET", "/api/public/organizations/projects", nil)
	if err != nil {
		return nil, err
	}

	r := new(GetProjectsOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// APIKey represents a Langfuse API key.
type APIKey struct {
	ID               string     `json:"id"`
	CreatedAt        time.Time  `json:"createdAt"`
	ExpiresAt        *time.Time `json:"expiresAt,omitempty"`
	LastUsedAt       *time.Time `json:"lastUsedAt,omitempty"`
	Note             *string    `json:"note,omitempty"`
	PublicKey        string     `json:"publicKey"`
	DisplaySecretKey string     `json:"displaySecretKey"`
}

// GetAPIKeysOutput is the response for listing API keys.
type GetAPIKeysOutput struct {
	Data []*APIKey   `json:"data"`
	Meta *Pagination `json:"meta"`
}

// GetOrganizationAPIKeys returns all API keys for the organization.
func (c *Client) GetOrganizationAPIKeys(ctx context.Context) (*GetAPIKeysOutput, error) {
	req, err := c.NewRequest("GET", "/api/public/organizations/apiKeys", nil)
	if err != nil {
		return nil, err
	}

	r := new(GetAPIKeysOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetProjectAPIKeys returns all API keys for a project.
func (c *Client) GetProjectAPIKeys(ctx context.Context, projectID string) (*GetAPIKeysOutput, error) {
	path := fmt.Sprintf("/api/public/projects/%s/apiKeys", projectID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(GetAPIKeysOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateAPIKeyInput is the request body for creating a project API key.
type CreateAPIKeyInput struct {
	Note *string `json:"note,omitempty"`
}

// CreateAPIKeyOutput is the response after creating an API key.
// SecretKey is only returned once at creation time.
type CreateAPIKeyOutput struct {
	ID               string     `json:"id"`
	CreatedAt        time.Time  `json:"createdAt"`
	ExpiresAt        *time.Time `json:"expiresAt,omitempty"`
	Note             *string    `json:"note,omitempty"`
	PublicKey        string     `json:"publicKey"`
	DisplaySecretKey string     `json:"displaySecretKey"`
	SecretKey        string     `json:"secretKey"`
}

// CreateProjectAPIKey creates a new API key for a project.
func (c *Client) CreateProjectAPIKey(ctx context.Context, projectID string, input *CreateAPIKeyInput) (*CreateAPIKeyOutput, error) {
	path := fmt.Sprintf("/api/public/projects/%s/apiKeys", projectID)

	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(CreateAPIKeyOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteProjectAPIKey deletes an API key from a project.
func (c *Client) DeleteProjectAPIKey(ctx context.Context, projectID, apiKeyID string) error {
	path := fmt.Sprintf("/api/public/projects/%s/apiKeys/%s", projectID, apiKeyID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}
