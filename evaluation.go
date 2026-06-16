package langfuse

// NOTE: These APIs are marked as UNSTABLE by Langfuse and may change without notice.

import (
	"context"
	"fmt"
)

// EvaluationRuleState is the state of an evaluation rule.
type EvaluationRuleState string

const (
	EvaluationRuleStateActive   EvaluationRuleState = "ACTIVE"
	EvaluationRuleStateInactive EvaluationRuleState = "INACTIVE"
)

// EvaluationRuleTarget is the target type for an evaluation rule.
type EvaluationRuleTarget string

const (
	EvaluationRuleTargetTrace      EvaluationRuleTarget = "TRACE"
	EvaluationRuleTargetDatasetRun EvaluationRuleTarget = "DATASET_RUN"
)

// EvaluationRule represents an evaluation rule configuration.
type EvaluationRule struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	State       EvaluationRuleState  `json:"state"`
	Target      EvaluationRuleTarget `json:"target"`
	Filter      any                  `json:"filter,omitempty"`
	EvaluatorID string               `json:"evaluatorId"`
	Mapping     any                  `json:"mapping,omitempty"`
	Sampling    *float64             `json:"sampling,omitempty"`
	Priority    *int                 `json:"priority,omitempty"`
}

// GetEvaluationRulesOutput is the response for listing evaluation rules.
type GetEvaluationRulesOutput struct {
	Data []*EvaluationRule `json:"data"`
	Meta *Pagination       `json:"meta"`
}

// GetEvaluationRulesOptions are the query parameters for listing evaluation rules.
type GetEvaluationRulesOptions struct {
	Page  *int                 `url:"page,omitempty"`
	Limit *int                 `url:"limit,omitempty"`
	State *EvaluationRuleState `url:"state,omitempty"`
}

// GetEvaluationRules returns a paginated list of evaluation rules (unstable API).
func (c *Client) GetEvaluationRules(
	ctx context.Context, opts *GetEvaluationRulesOptions,
) (*GetEvaluationRulesOutput, error) {
	path := "/api/public/unstable/evaluation-rules"
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

	r := new(GetEvaluationRulesOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateEvaluationRuleInput is the request body for creating an evaluation rule.
type CreateEvaluationRuleInput struct {
	Name        string               `json:"name"`
	Target      EvaluationRuleTarget `json:"target"`
	Filter      any                  `json:"filter,omitempty"`
	EvaluatorID string               `json:"evaluatorId"`
	Mapping     any                  `json:"mapping,omitempty"`
	Sampling    *float64             `json:"sampling,omitempty"`
	Priority    *int                 `json:"priority,omitempty"`
}

// CreateEvaluationRule creates a new evaluation rule (unstable API).
func (c *Client) CreateEvaluationRule(ctx context.Context, input *CreateEvaluationRuleInput) (*EvaluationRule, error) {
	req, err := c.NewRequest("POST", "/api/public/unstable/evaluation-rules", input)
	if err != nil {
		return nil, err
	}

	r := new(EvaluationRule)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetEvaluationRule returns a single evaluation rule by ID (unstable API).
func (c *Client) GetEvaluationRule(ctx context.Context, ruleID string) (*EvaluationRule, error) {
	path := fmt.Sprintf("/api/public/unstable/evaluation-rules/%s", ruleID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(EvaluationRule)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// UpdateEvaluationRuleInput is the request body for updating an evaluation rule.
type UpdateEvaluationRuleInput struct {
	Name     *string              `json:"name,omitempty"`
	State    *EvaluationRuleState `json:"state,omitempty"`
	Filter   any                  `json:"filter,omitempty"`
	Mapping  any                  `json:"mapping,omitempty"`
	Sampling *float64             `json:"sampling,omitempty"`
	Priority *int                 `json:"priority,omitempty"`
}

// UpdateEvaluationRule updates an existing evaluation rule (unstable API).
func (c *Client) UpdateEvaluationRule(
	ctx context.Context, ruleID string, input *UpdateEvaluationRuleInput,
) (*EvaluationRule, error) {
	path := fmt.Sprintf("/api/public/unstable/evaluation-rules/%s", ruleID)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(EvaluationRule)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteEvaluationRule deletes an evaluation rule by ID (unstable API).
func (c *Client) DeleteEvaluationRule(ctx context.Context, ruleID string) error {
	path := fmt.Sprintf("/api/public/unstable/evaluation-rules/%s", ruleID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}

// EvaluatorType is the type of evaluator.
type EvaluatorType string

const (
	EvaluatorTypeLLMAsJudge EvaluatorType = "llm_as_judge"
	EvaluatorTypeCode       EvaluatorType = "code"
)

// EvaluatorScope is the scope of an evaluator.
type EvaluatorScope string

const (
	EvaluatorScopeProject EvaluatorScope = "project"
	EvaluatorScopeManaged EvaluatorScope = "managed"
)

// Evaluator represents an evaluator configuration.
type Evaluator struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Type       EvaluatorType  `json:"type"`
	Scope      EvaluatorScope `json:"scope"`
	Prompt     any            `json:"prompt,omitempty"`
	SourceCode *string        `json:"sourceCode,omitempty"`
}

// GetEvaluatorsOutput is the response for listing evaluators.
type GetEvaluatorsOutput struct {
	Data []*Evaluator `json:"data"`
	Meta *Pagination  `json:"meta"`
}

// GetEvaluatorsOptions are the query parameters for listing evaluators.
type GetEvaluatorsOptions struct {
	Page  *int            `url:"page,omitempty"`
	Limit *int            `url:"limit,omitempty"`
	Scope *EvaluatorScope `url:"scope,omitempty"`
}

// GetEvaluators returns a paginated list of evaluators (unstable API).
func (c *Client) GetEvaluators(ctx context.Context, opts *GetEvaluatorsOptions) (*GetEvaluatorsOutput, error) {
	path := "/api/public/unstable/evaluators"
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

	r := new(GetEvaluatorsOutput)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// CreateEvaluatorInput is the request body for creating an evaluator.
type CreateEvaluatorInput struct {
	Name       string        `json:"name"`
	Type       EvaluatorType `json:"type"`
	Prompt     any           `json:"prompt,omitempty"`
	SourceCode *string       `json:"sourceCode,omitempty"`
}

// CreateEvaluator creates a new evaluator (unstable API).
func (c *Client) CreateEvaluator(ctx context.Context, input *CreateEvaluatorInput) (*Evaluator, error) {
	req, err := c.NewRequest("POST", "/api/public/unstable/evaluators", input)
	if err != nil {
		return nil, err
	}

	r := new(Evaluator)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetEvaluator returns a single evaluator by ID (unstable API).
func (c *Client) GetEvaluator(ctx context.Context, evaluatorID string) (*Evaluator, error) {
	path := fmt.Sprintf("/api/public/unstable/evaluators/%s", evaluatorID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(Evaluator)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// UpdateEvaluatorInput is the request body for updating an evaluator.
type UpdateEvaluatorInput struct {
	Name       *string `json:"name,omitempty"`
	Prompt     any     `json:"prompt,omitempty"`
	SourceCode *string `json:"sourceCode,omitempty"`
}

// UpdateEvaluator updates an existing evaluator (unstable API).
func (c *Client) UpdateEvaluator(
	ctx context.Context, evaluatorID string, input *UpdateEvaluatorInput,
) (*Evaluator, error) {
	path := fmt.Sprintf("/api/public/unstable/evaluators/%s", evaluatorID)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(Evaluator)
	if err = c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// DeleteEvaluator deletes an evaluator by ID (unstable API).
func (c *Client) DeleteEvaluator(ctx context.Context, evaluatorID string) error {
	path := fmt.Sprintf("/api/public/unstable/evaluators/%s", evaluatorID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	return c.Do(ctx, req, nil)
}
