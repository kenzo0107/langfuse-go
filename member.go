package langfuse

import (
	"context"
	"fmt"
)

// MembershipRole is the role of a member within a project or organization.
type MembershipRole string

const (
	MembershipRoleOwner  MembershipRole = "OWNER"
	MembershipRoleAdmin  MembershipRole = "ADMIN"
	MembershipRoleMember MembershipRole = "MEMBER"
	MembershipRoleViewer MembershipRole = "VIEWER"
)

// Membership represents a single member's role within a project or organization.
type Membership struct {
	UserID string         `json:"userId"`
	Role   MembershipRole `json:"role"`
	Email  string         `json:"email"`
	Name   string         `json:"name"`
}

// GetMembershipsOutput is the response for listing memberships.
type GetMembershipsOutput struct {
	Memberships []*Membership `json:"memberships"`
}

// UpsertMembershipInput is the request body for adding or updating a membership.
type UpsertMembershipInput struct {
	UserID string         `json:"userId"`
	Role   MembershipRole `json:"role"`
}

// DeleteMembershipInput is the request body for removing a membership.
type DeleteMembershipInput struct {
	UserID string `json:"userId"`
}

// DeleteMembershipOutput is the response after removing a membership.
type DeleteMembershipOutput struct {
	Message string `json:"message"`
	UserID  string `json:"userId"`
}

// GetOrganizationMemberships returns all members of the organization.
func (c *Client) GetOrganizationMemberships(ctx context.Context) (*GetMembershipsOutput, error) {
	req, err := c.NewRequest("GET", "/api/public/organizations/memberships", nil)
	if err != nil {
		return nil, err
	}

	r := new(GetMembershipsOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// UpsertOrganizationMembership adds or updates a member in the organization.
func (c *Client) UpsertOrganizationMembership(ctx context.Context, input *UpsertMembershipInput) (*Membership, error) {
	req, err := c.NewRequest("PUT", "/api/public/organizations/memberships", input)
	if err != nil {
		return nil, err
	}

	r := new(Membership)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteOrganizationMembership removes a member from the organization.
func (c *Client) DeleteOrganizationMembership(
	ctx context.Context, input *DeleteMembershipInput,
) (*DeleteMembershipOutput, error) {
	req, err := c.NewRequest("DELETE", "/api/public/organizations/memberships", input)
	if err != nil {
		return nil, err
	}

	r := new(DeleteMembershipOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetProjectMemberships returns all members of the specified project.
func (c *Client) GetProjectMemberships(ctx context.Context, projectID string) (*GetMembershipsOutput, error) {
	path := fmt.Sprintf("/api/public/projects/%s/memberships", projectID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(GetMembershipsOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// UpsertProjectMembership adds or updates a member in the specified project.
// The user must already be a member of the organization.
func (c *Client) UpsertProjectMembership(
	ctx context.Context, projectID string, input *UpsertMembershipInput,
) (*Membership, error) {
	path := fmt.Sprintf("/api/public/projects/%s/memberships", projectID)

	req, err := c.NewRequest("PUT", path, input)
	if err != nil {
		return nil, err
	}

	r := new(Membership)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteProjectMembership removes a member from the specified project.
func (c *Client) DeleteProjectMembership(
	ctx context.Context, projectID string, input *DeleteMembershipInput,
) (*DeleteMembershipOutput, error) {
	path := fmt.Sprintf("/api/public/projects/%s/memberships", projectID)

	req, err := c.NewRequest("DELETE", path, input)
	if err != nil {
		return nil, err
	}

	r := new(DeleteMembershipOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}
