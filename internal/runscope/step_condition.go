package runscope

import (
	"context"
	"fmt"
	"net/http"

	"github.com/terraform-providers/terraform-provider-runscope/internal/runscope/schema"
)

func StepConditionFromSchema(s *schema.StepCondition) *StepCondition {
	step := StepCondition{
		ID: s.ID,
	}

	step.setFromSchema(s)
	return &step
}

type StepCondition struct {
	ID                   string
	BucketKey            string
	TestUUID             string
	EnvironmentUUID      string
	UseParentEnvironment bool
	LeftValue  			 string
	Comparison 			 string
	RightValue 			 string
	Variables            []StepVariable
	Assertions           []StepAssertion
}

func (ss *StepCondition) setFromSchema(s *schema.StepCondition) {
	ss.BucketKey = s.BucketKey
	ss.EnvironmentUUID = s.EnvironmentUUID
	ss.TestUUID = s.TestUUID
	ss.UseParentEnvironment = s.UseParentEnvironment

	ss.Variables = make([]StepVariable, len(s.Variables))
	ss.Assertions = make([]StepAssertion, len(s.Assertions))
	for i, v := range s.Variables {
		ss.Variables[i] = StepVariable{
			Name:     v.Name,
			Property: v.Property,
			Source:   v.Source,
		}
	}
	for i, a := range s.Assertions {
		ss.Assertions[i] = StepAssertion{
			Source:     a.Source,
			Property:   a.Property,
			Comparison: a.Comparison,
			Value:      a.Value,
		}
	}
}

type StepConditionOpts struct {
	ID                   string
	BucketKey            string
	TestUUID             string
	EnvironmentUUID      string
	UseParentEnvironment bool
	LeftValue            string
	Comparison           string
	RightValue           string
	Assertions           []StepAssertion
	Variables            []StepVariable
}

func (sso *StepConditionOpts) setRequest(ss *schema.StepCondition) {
	ss.BucketKey = sso.BucketKey
	ss.TestUUID = sso.TestUUID
	ss.EnvironmentUUID = sso.EnvironmentUUID
	ss.UseParentEnvironment = sso.UseParentEnvironment

	ss.Variables = make([]schema.StepVariable, len(sso.Variables))
	ss.Assertions = make([]schema.StepAssertion, len(sso.Assertions))
	for i, v := range sso.Variables {
		ss.Variables[i] = schema.StepVariable{
			Name:     v.Name,
			Property: v.Property,
			Source:   v.Source,
		}
	}
	for i, a := range sso.Assertions {
		ss.Assertions[i] = schema.StepAssertion{
			Source:     a.Source,
			Comparison: a.Comparison,
			Value:      a.Value,
			Property:   a.Property,
		}
	}
}

type StepCreateConditionOpts struct {
	StepUriOpts
	StepConditionOpts
}

func (c *StepClient) CreateCondition(ctx context.Context, opts *StepCreateConditionOpts) (*StepCondition, error) {
	body := schema.StepCreateConditionRequest{
		StepType: "condition",
	}
	opts.StepConditionOpts.setRequest(&body.StepCondition)
	req, err := c.client.NewRequest(ctx, http.MethodPost, opts.StepUriOpts.URL(), &body)
	if err != nil {
		return nil, err
	}

	var resp schema.StepCreateConditionResponse
	if err := c.client.Do(req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Step) < 1 {
		return nil, fmt.Errorf("no steps returned after created")
	}

	return StepConditionFromSchema(&resp.Step[len(resp.Step)-1]), nil
}

func (c *StepClient) GetCondition(ctx context.Context, opts *StepGetRequestOpts) (*StepCondition, error) {
	var resp schema.StepGetConditionResponse
	req, err := c.client.NewRequest(ctx, http.MethodGet, opts.URL(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create req: %w", err)
	}

	err = c.client.Do(req, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to call endpoint: %w", err)
	}

	return StepConditionFromSchema(&resp.Step), nil
}

type StepUpdateConditionOpts struct {
	StepGetRequestOpts
	StepConditionOpts
}

func (c *StepClient) UpdateCondition(ctx context.Context, opts *StepUpdateConditionOpts) (*StepCondition, error) {
	body := &schema.StepUpdateConditionRequest{
		StepType: "condition",
	}
	opts.setRequest(&body.StepCondition)
	body.ID = opts.Id

	req, err := c.client.NewRequest(ctx, http.MethodPut, opts.URL(), &body)
	if err != nil {
		return nil, err
	}

	var resp schema.StepUpdateConditionResponse
	err = c.client.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return StepConditionFromSchema(&resp.Step), nil
}
