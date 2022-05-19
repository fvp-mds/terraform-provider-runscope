package runscope

import (
	"context"
	"net/http"

	"github.com/terraform-providers/terraform-provider-runscope/internal/runscope/schema"
)

type AccountClient struct {
	client *Client
}

type Account struct {
	Name  string
	UUID  string
	Email string
	Teams []Team
}

func AccountFromSchema(s *schema.Account) *Account {
	a := Account{
		Name:  s.Name,
		UUID:  s.UUID,
		Email: s.Email,
	}

	for _, t := range s.Teams {
		a.Teams = append(a.Teams, Team{
			Name: t.Name,
			UUID: t.UUID,
		})
	}

	return &a
}

type AccountGetOpts struct {
}

func (AccountGetOpts) URL() string {
	return "/account"
}

func (a *AccountClient) Get(ctx context.Context, opts *AccountGetOpts) (*Account, error) {
	req, err := a.client.NewRequest(ctx, http.MethodGet, opts.URL(), nil)
	if err != nil {
		return nil, err
	}

	var resp schema.AccountResponse
	if err := a.client.Do(req, &resp); err != nil {
		return nil, err
	}

	return AccountFromSchema(&resp.Data), nil

}
