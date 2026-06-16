package langfuse

import (
	"context"
	"fmt"
	"time"
)

// CommentObjectType is the type of object a comment is attached to.
type CommentObjectType string

const (
	CommentObjectTypeTrace       CommentObjectType = "TRACE"
	CommentObjectTypeObservation CommentObjectType = "OBSERVATION"
	CommentObjectTypeSession     CommentObjectType = "SESSION"
	CommentObjectTypePrompt      CommentObjectType = "PROMPT"
)

// Comment represents a Langfuse comment on a trace, observation, session, or prompt.
type Comment struct {
	ID           string            `json:"id"`
	ProjectID    string            `json:"projectId"`
	ObjectType   CommentObjectType `json:"objectType"`
	ObjectID     string            `json:"objectId"`
	Content      string            `json:"content"`
	AuthorUserID *string           `json:"authorUserId,omitempty"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
}

// GetCommentsOutput is the response for listing comments.
type GetCommentsOutput struct {
	Data []*Comment  `json:"data"`
	Meta *Pagination `json:"meta"`
}

// GetCommentsOptions are the query parameters for listing comments.
type GetCommentsOptions struct {
	Page         *int               `url:"page,omitempty"`
	Limit        *int               `url:"limit,omitempty"`
	ObjectType   *CommentObjectType `url:"objectType,omitempty"`
	ObjectID     *string            `url:"objectId,omitempty"`
	AuthorUserID *string            `url:"authorUserId,omitempty"`
}

// GetComments returns a paginated list of comments.
func (c *Client) GetComments(ctx context.Context, opts *GetCommentsOptions) (*GetCommentsOutput, error) {
	path := "/api/public/comments"
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

	r := new(GetCommentsOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateCommentInput is the request body for creating a comment.
type CreateCommentInput struct {
	ObjectType   CommentObjectType `json:"objectType"`
	ObjectID     string            `json:"objectId"`
	Content      string            `json:"content"`
	AuthorUserID *string           `json:"authorUserId,omitempty"`
}

// CreateCommentOutput is the response after creating a comment.
type CreateCommentOutput struct {
	ID string `json:"id"`
}

// CreateComment creates a new comment.
func (c *Client) CreateComment(ctx context.Context, input *CreateCommentInput) (*CreateCommentOutput, error) {
	req, err := c.NewRequest("POST", "/api/public/comments", input)
	if err != nil {
		return nil, err
	}

	r := new(CreateCommentOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetComment returns a single comment by ID.
func (c *Client) GetComment(ctx context.Context, commentID string) (*Comment, error) {
	path := fmt.Sprintf("/api/public/comments/%s", commentID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(Comment)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}
