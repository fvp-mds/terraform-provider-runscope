terraform {
  required_providers {
    runscope = {
      source  = "Storytel/runscope"
      version = ">= 0.14.0"
    }
  }
}

variable "access_token" {
  type      = string
  sensitive = true
}

variable "team_uuid" {
  type = string
}

provider "runscope" {
  access_token = var.access_token
}

resource "runscope_bucket" "my_bucket" {
  name      = "My Bucket [Managed by Terraform]"
  team_uuid = var.team_uuid
}
