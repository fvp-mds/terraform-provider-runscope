package runscope

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-providers/terraform-provider-runscope/internal/runscope/schema"
)

type StepBase struct {
	StepType      string
	Method        string
	StepURL       string
	Variables     []StepVariable
	Assertions    []StepAssertion
	Headers       map[string][]string
	Auth          StepAuth
	Body          string
	Form          map[string][]string
	Scripts       []string
	BeforeScripts []string
	Note          string
	Skipped       bool
}

func (sb *StepBase) setFromSchema(s *schema.Step) {
	sb.StepType = s.StepType
	sb.Method = s.Method
	sb.StepURL = s.URL
	sb.Variables = make([]StepVariable, len(s.Variables))
	sb.Assertions = make([]StepAssertion, len(s.Assertions))
	sb.Headers = map[string][]string{}
	sb.Auth = StepAuth{
		Username: s.Auth.Username,
		Password: s.Auth.Password,
		AuthType: s.Auth.AuthType,
	}
	sb.Body = s.Body
	sb.Form = map[string][]string{}
	sb.Scripts = make([]string, len(s.Scripts))
	sb.BeforeScripts = make([]string, len(s.BeforeScripts))
	sb.Note = s.Note
	sb.Skipped = s.Skipped

	for i, v := range s.Variables {
		sb.Variables[i] = StepVariable{
			Name:     v.Name,
			Property: v.Property,
			Source:   v.Source,
		}
	}
	for i, a := range s.Assertions {
		sb.Assertions[i] = StepAssertion{
			Source:     a.Source,
			Property:   a.Property,
			Comparison: a.Comparison,
			Value:      a.Value,
		}
	}
	for header, values := range s.Headers {
		sb.Headers[header] = make([]string, len(values))
		copy(sb.Headers[header], values)
	}
	for name, values := range s.Form {
		sb.Form[name] = make([]string, len(values))
		copy(sb.Form[name], values)
	}
	copy(sb.Scripts, s.Scripts)
	copy(sb.BeforeScripts, s.BeforeScripts)
}

type StepSubtest struct {
	ID              string
	BucketKey       string
	TestUUID        string
	EnvironmentUUID string
	Variables       []StepVariable
	Assertions      []StepAssertion
}

func (ss *StepSubtest) setFromSchema(s *schema.StepSubtest) {
	ss.BucketKey = s.BucketKey
	ss.EnvironmentUUID = s.EnvironmentUUID
	ss.TestUUID = s.TestUUID

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

type Step struct {
	StepBase
	Id string
}

type StepVariable struct {
	Name     string
	Property string
	Source   string
}

type StepAssertion struct {
	Source     string
	Property   string
	Comparison string
	Value      string
}

type StepAuth struct {
	Username string
	Password string
	AuthType string
}

func (s StepAuth) Empty() bool {
	return s.Username == "" && s.Password == "" && s.AuthType == ""
}

type StepClient struct {
	client *Client
}

func StepFromSchema(s *schema.Step) *Step {
	step := &Step{
		Id: s.Id,
	}
	step.StepBase.setFromSchema(s)

	return step
}

func StepSubtestFromSchema(s *schema.StepSubtest) *StepSubtest {
	step := StepSubtest{
		ID: s.ID,
	}

	step.setFromSchema(s)
	return &step
}

type StepUriOpts struct {
	BucketId string
	TestId   string
}

func (s StepUriOpts) URL() string {
	return fmt.Sprintf("/buckets/%s/tests/%s/steps", s.BucketId, s.TestId)
}

type StepBaseOpts struct {
	StepType      string
	Method        string
	StepURL       string
	Variables     []StepVariable
	Assertions    []StepAssertion
	Headers       map[string][]string
	Auth          StepAuth
	Body          string
	Form          map[string][]string
	Scripts       []string
	BeforeScripts []string
	Note          string
	Skipped       bool
}

func (sbo *StepBaseOpts) setRequest(sb *schema.StepBase) {
	sb.StepType = sbo.StepType
	sb.Method = sbo.Method
	sb.URL = sbo.StepURL
	sb.Variables = make([]schema.StepVariable, len(sbo.Variables))
	sb.Assertions = make([]schema.StepAssertion, len(sbo.Assertions))
	sb.Headers = map[string][]string{}
	sb.Auth = schema.StepAuth{
		Username: sbo.Auth.Username,
		Password: sbo.Auth.Password,
		AuthType: sbo.Auth.AuthType,
	}
	sb.Body = sbo.Body
	sb.Form = map[string][]string{}
	sb.Scripts = make([]string, len(sbo.Scripts))
	sb.BeforeScripts = make([]string, len(sbo.BeforeScripts))
	sb.Note = sbo.Note
	sb.Skipped = sbo.Skipped

	for i, v := range sbo.Variables {
		sb.Variables[i] = schema.StepVariable{
			Name:     v.Name,
			Property: v.Property,
			Source:   v.Source,
		}
	}
	for i, a := range sbo.Assertions {
		sb.Assertions[i] = schema.StepAssertion{
			Source:     a.Source,
			Comparison: a.Comparison,
			Value:      a.Value,
			Property:   a.Property,
		}
	}
	for header, values := range sbo.Headers {
		sb.Headers[header] = make([]string, len(values))
		copy(sb.Headers[header], values)
	}
	for name, values := range sbo.Form {
		sb.Form[name] = make([]string, len(values))
		copy(sb.Form[name], values)
	}
	copy(sb.Scripts, sbo.Scripts)
	copy(sb.BeforeScripts, sbo.BeforeScripts)
}

type StepSubtestOpts struct {
	TestUUID        string
	EnvironmentUUID string
	BucketKey       string
	Assertions      []StepAssertion
	Variables       []StepVariable
}

func (sso *StepSubtestOpts) setRequest(ss *schema.StepSubtest) {
	ss.BucketKey = sso.BucketKey
	ss.TestUUID = sso.TestUUID
	ss.EnvironmentUUID = sso.EnvironmentUUID

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

type StepCreateSubtestOpts struct {
	StepUriOpts
	StepSubtestOpts
}

func (c *StepClient) CreateSubtest(ctx context.Context, opts *StepCreateSubtestOpts) (*Step, error) {
	body := schema.StepCreateSubtestRequest{
		StepType: "subtest",
	}
	opts.setRequest(&body.StepSubtest)
	req, err := c.client.NewRequest(ctx, http.MethodPost, opts.StepUriOpts.URL(), &body)
	if err != nil {
		return nil, err
	}

	var resp schema.StepCreateResponse
	if err := c.client.Do(req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Step) < 1 {
		return nil, fmt.Errorf("no steps returned after created")
	}

	return StepFromSchema(&resp.Step[len(resp.Step)-1]), nil
}

type StepCreateOpts struct {
	StepUriOpts
	StepBaseOpts
}

func (c *StepClient) Create(ctx context.Context, opts *StepCreateOpts) (*Step, error) {
	body := &schema.StepCreateRequest{}
	opts.StepBaseOpts.setRequest(&body.StepBase)

	req, err := c.client.NewRequest(ctx, http.MethodPost, opts.URL(), &body)
	if err != nil {
		return nil, err
	}

	var resp schema.StepCreateResponse
	err = c.client.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return StepFromSchema(&resp.Step[len(resp.Step)-1]), nil
}

type StepGetOpts struct {
	StepUriOpts
	Id string
}

func (opts *StepGetOpts) URL() string {
	return fmt.Sprintf("%s/%s", opts.StepUriOpts.URL(), opts.Id)
}

func (c *StepClient) Get(ctx context.Context, opts *StepGetOpts) (*Step, error) {
	var resp schema.StepGetResponse
	req, err := c.client.NewRequest(ctx, http.MethodGet, opts.URL(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create req: %w", err)
	}

	err = c.client.Do(req, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to call endpoint: %w", err)
	}

	return StepFromSchema(&resp.Step), nil
}

func (c *StepClient) GetSubtest(ctx context.Context, opts *StepGetOpts) (*StepSubtest, error) {
	var resp schema.StepGetSubstepResponse
	req, err := c.client.NewRequest(ctx, http.MethodGet, opts.URL(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create req: %w", err)
	}

	err = c.client.Do(req, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to call endpoint: %w", err)
	}

	return StepSubtestFromSchema(&resp.Step), nil
}

type StepUpdateOpts struct {
	StepGetOpts
	StepBaseOpts
}

func (c *StepClient) Update(ctx context.Context, opts *StepUpdateOpts) (*Step, error) {
	body := &schema.StepUpdateRequest{}
	opts.StepBaseOpts.setRequest(&body.StepBase)
	body.Id = opts.Id

	req, err := c.client.NewRequest(ctx, http.MethodPut, opts.URL(), &body)
	if err != nil {
		return nil, err
	}

	var resp schema.StepUpdateResponse
	err = c.client.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return StepFromSchema(&resp.Step), nil
}

type StepUpdateSubtestOpts struct {
	StepGetOpts
	StepSubtestOpts
}

func (c *StepClient) UpdateSubtest(ctx context.Context, opts *StepUpdateSubtestOpts) (*StepSubtest, error) {
	body := &schema.StepUpdateSubtestRequest{
		StepType: "subtest",
	}
	opts.setRequest(&body.StepSubtest)
	body.ID = opts.Id

	tflog.Info(ctx, "updating subtest", map[string]interface{}{"url": opts.URL(), "body": body})
	req, err := c.client.NewRequest(ctx, http.MethodPut, opts.URL(), &body)
	if err != nil {
		return nil, err
	}

	var resp schema.StepUpdateSubtestResponse
	err = c.client.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return StepSubtestFromSchema(&resp.Step), nil
}

type StepDeleteOpts struct {
	StepGetOpts
}

func (c *StepClient) Delete(ctx context.Context, opts *StepDeleteOpts) error {
	req, err := c.client.NewRequest(ctx, "DELETE", opts.URL(), nil)
	if err != nil {
		return err
	}

	err = c.client.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
