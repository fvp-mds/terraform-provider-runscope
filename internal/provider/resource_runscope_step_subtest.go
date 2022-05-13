package provider

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-providers/terraform-provider-runscope/internal/runscope"
)

func resourceRunscopeStepSubtest() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceStepSubtestCreate,
		ReadContext:   resourceStepSubtestRead,
		UpdateContext: resourceStepSubtestUpdate,
		DeleteContext: resourceStepDelete,
		Schema: map[string]*schema.Schema{
			"bucket_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"test_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_bucket_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_test_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_environment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"variable": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"property": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"source": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice(stepSources, false),
						},
					},
				},
				Optional: true,
			},
			"assertion": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice(stepSources, false),
						},
						"property": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"comparison": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice(stepComparisons, false),
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				Optional: true,
			},
		},
	}
}

func resourceStepSubtestCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*providerConfig).client

	var opts runscope.StepCreateSubtestOpts
	expandStepUriOpts(d, &opts.StepUriOpts)
	expandStepSubtestOpts(d, &opts.StepSubtestOpts)

	step, err := client.Step.CreateSubtest(ctx, &opts)
	if err != nil {
		if isNotFound(err) {
			d.SetId("")
			return nil
		}

		return diag.Errorf("couldn't read step: %s", err)
	}

	d.SetId(step.Id)

	return resourceStepSubtestRead(ctx, d, meta)
}

func resourceStepSubtestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*providerConfig).client

	opts := &runscope.StepGetOpts{
		StepUriOpts: runscope.StepUriOpts{
			BucketId: d.Get("bucket_id").(string),
			TestId:   d.Get("test_id").(string),
		},
		Id: d.Id(),
	}

	step, err := client.Step.GetSubtest(ctx, opts)
	if err != nil {
		if isNotFound(err) {
			d.SetId("")
			return nil
		}
		var runscopeErr runscope.Error
		is := errors.As(err, &runscopeErr)

		return diag.Errorf("couldn't (is=%t) read step: %s", is, err)
	}

	d.Set("source_bucket_id", step.BucketKey)
	d.Set("source_test_id", step.TestUUID)
	d.Set("source_environment_id", step.EnvironmentUUID)
	d.Set("variable", flattenStepVariables(step.Variables))
	d.Set("assertion", flattenStepAssertions(step.Assertions))

	return nil
}

func resourceStepSubtestUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*providerConfig).client

	opts := &runscope.StepUpdateSubtestOpts{}
	expandStepGetOpts(d, &opts.StepGetOpts)
	expandStepSubtestOpts(d, &opts.StepSubtestOpts)

	_, err := client.Step.UpdateSubtest(ctx, opts)
	if err != nil {
		return diag.Errorf("Couldn't create step: %s", err)
	}

	return resourceStepSubtestRead(ctx, d, meta)
}

func expandStepSubtestOpts(d *schema.ResourceData, opts *runscope.StepSubtestOpts) {
	if v, ok := d.GetOk("source_bucket_id"); ok {
		opts.BucketKey = v.(string)
	}
	if v, ok := d.GetOk("source_test_id"); ok {
		opts.TestUUID = v.(string)
	}
	if v, ok := d.GetOk("source_environment_id"); ok {
		opts.EnvironmentUUID = v.(string)
	}
	if v, ok := d.GetOk("variable"); ok {
		opts.Variables = expandStepVariables(v.(*schema.Set).List())
	}
	if v, ok := d.GetOk("assertion"); ok {
		opts.Assertions = expandStepAssertions(v.([]interface{}))
	}
}
