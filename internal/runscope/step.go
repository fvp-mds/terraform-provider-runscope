package runscope

import (
	"context"
	"fmt"
	"net/http"
)

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

type StepUriOpts struct {
	BucketId string
	TestId   string
}

func (s StepUriOpts) URL() string {
	return fmt.Sprintf("/buckets/%s/tests/%s/steps", s.BucketId, s.TestId)
}

type StepDeleteOpts struct {
	StepGetRequestOpts
}

func (c *StepClient) Delete(ctx context.Context, opts *StepDeleteOpts) error {
	req, err := c.client.NewRequest(ctx, http.MethodDelete, opts.URL(), nil)
	if err != nil {
		return err
	}

	err = c.client.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
