package langfuse

import (
	"context"
	"fmt"
)

// SCIMUserName holds the name components of a SCIM user.
type SCIMUserName struct {
	Formatted  string `json:"formatted,omitempty"`
	GivenName  string `json:"givenName,omitempty"`
	FamilyName string `json:"familyName,omitempty"`
}

// SCIMEmail holds a single email entry for a SCIM user.
type SCIMEmail struct {
	Value   string `json:"value"`
	Primary bool   `json:"primary,omitempty"`
	Type    string `json:"type,omitempty"`
}

// SCIMUser represents a SCIM 2.0 User resource.
type SCIMUser struct {
	ID         string       `json:"id"`
	ExternalID *string      `json:"externalId,omitempty"`
	UserName   string       `json:"userName"`
	Name       SCIMUserName `json:"name"`
	Emails     []SCIMEmail  `json:"emails"`
	Active     bool         `json:"active"`
}

// SCIMListResponse is the SCIM 2.0 list response envelope.
type SCIMListResponse struct {
	TotalResults int         `json:"totalResults"`
	StartIndex   int         `json:"startIndex"`
	ItemsPerPage int         `json:"itemsPerPage"`
	Resources    []*SCIMUser `json:"Resources"`
}

// SCIMListUsersOptions are the query parameters for listing SCIM users.
type SCIMListUsersOptions struct {
	Filter     *string `url:"filter,omitempty"`
	StartIndex *int    `url:"startIndex,omitempty"`
	Count      *int    `url:"count,omitempty"`
}

// GetSCIMUsers returns a list of SCIM users (requires org-scoped key).
func (c *Client) GetSCIMUsers(ctx context.Context, opts *SCIMListUsersOptions) (*SCIMListResponse, error) {
	path := "/api/public/scim/Users"
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

	r := new(SCIMListResponse)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateSCIMUserInput is the request body for creating a SCIM user.
type CreateSCIMUserInput struct {
	UserName   string       `json:"userName"`
	Name       SCIMUserName `json:"name,omitempty"`
	Emails     []SCIMEmail  `json:"emails"`
	Active     *bool        `json:"active,omitempty"`
	ExternalID *string      `json:"externalId,omitempty"`
	Password   *string      `json:"password,omitempty"`
}

// CreateSCIMUser creates a new user via SCIM (requires org-scoped key).
func (c *Client) CreateSCIMUser(ctx context.Context, input *CreateSCIMUserInput) (*SCIMUser, error) {
	req, err := c.NewRequest("POST", "/api/public/scim/Users", input)
	if err != nil {
		return nil, err
	}

	r := new(SCIMUser)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetSCIMUser returns a single SCIM user by ID.
func (c *Client) GetSCIMUser(ctx context.Context, userID string) (*SCIMUser, error) {
	path := fmt.Sprintf("/api/public/scim/Users/%s", userID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(SCIMUser)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteSCIMUser deletes a SCIM user by ID.
func (c *Client) DeleteSCIMUser(ctx context.Context, userID string) error {
	path := fmt.Sprintf("/api/public/scim/Users/%s", userID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}

// SCIMServiceProviderConfig holds the SCIM service provider capabilities.
type SCIMServiceProviderConfig struct {
	DocumentationURI string `json:"documentationUri,omitempty"`
	AuthenticationSchemes []any `json:"authenticationSchemes,omitempty"`
	Supported             any   `json:"supported,omitempty"`
}

// GetSCIMServiceProviderConfig returns the SCIM service provider configuration.
func (c *Client) GetSCIMServiceProviderConfig(ctx context.Context) (*SCIMServiceProviderConfig, error) {
	req, err := c.NewRequest("GET", "/api/public/scim/ServiceProviderConfig", nil)
	if err != nil {
		return nil, err
	}

	r := new(SCIMServiceProviderConfig)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetSCIMResourceTypes returns the SCIM resource type definitions.
func (c *Client) GetSCIMResourceTypes(ctx context.Context) (any, error) {
	req, err := c.NewRequest("GET", "/api/public/scim/ResourceTypes", nil)
	if err != nil {
		return nil, err
	}

	var r any
	if err = c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetSCIMSchemas returns the SCIM schema definitions.
func (c *Client) GetSCIMSchemas(ctx context.Context) (any, error) {
	req, err := c.NewRequest("GET", "/api/public/scim/Schemas", nil)
	if err != nil {
		return nil, err
	}

	var r any
	if err = c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}
