package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-runscope/internal/runscope"
)

var stepSources = []string{
	"response_status",
	"response_headers",
	"response_json",
	"response_xml",
	"response_text",
	"response_time",
	"response_size",
}

var stepComparisons = []string{
	"equal",
	"empty",
	"not_empty",
	"not_equal",
	"contains",
	"does_not_contain",
	"is_a_number",
	"equal_number",
	"is_less_than",
	"is_less_than_or_equal",
	"is_greater_than",
	"is_greater_than_or_equal",
	"has_key",
	"has_value",
	"is_null",
}

func resourceStepDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*providerConfig).client

	opts := &runscope.StepDeleteOpts{}
	expandStepGetOpts(d, &opts.StepGetRequestOpts)

	if err := client.Step.Delete(ctx, opts); err != nil {
		return diag.Errorf("Couldn't read step: %s", err)
	}

	return nil
}

func expandStepUriOpts(d *schema.ResourceData, opts *runscope.StepUriOpts) {
	opts.BucketId = d.Get("bucket_id").(string)
	opts.TestId = d.Get("test_id").(string)
}
