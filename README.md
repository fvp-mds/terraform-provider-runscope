# Terraform Runscope Provider

> **_For internal use only, no support provided_**

The Runscope provider is used to interact with the resources
supported by [Runscope](https://runscope.com/).

### Usage Terraform 0.13+

Add into your Terraform configuration this code:

```hcl-terraform
terraform {
  required_providers {
    runscope = {
      source = "Storytel/runscope"
      version = "~> 0.12.0"
    }
  }
}
```

and run `terraform init`

## Usage

Read the [documentation on Terraform Registry site](https://registry.terraform.io/providers/sport24ru/runscope/latest/docs).
