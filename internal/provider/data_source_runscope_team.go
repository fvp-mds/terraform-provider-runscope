package provider

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-runscope/internal/runscope"
)

var ErrNotFound = errors.New("not found")

func dataSourceRunscopeTeam() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRunscopeTeamRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceRunscopeTeamRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*providerConfig).client

	opts := &runscope.AccountGetOpts{}
	account, err := client.Account.Get(ctx, opts)
	if err != nil {
		return diag.FromErr(err)
	}

	team, err := findTeam(account.Teams, d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(team.UUID)
	d.Set("name", team.Name)

	return nil
}

func findTeam(teams []runscope.Team, name string) (*runscope.Team, error) {
	for _, t := range teams {
		if t.Name == name {
			return &t, nil
		}
	}

	return nil, ErrNotFound
}
