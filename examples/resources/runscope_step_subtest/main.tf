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

variable "team_name" {
  type = string
}

provider "runscope" {
  access_token = var.access_token
}

data "runscope_team" "my_team" {
  name = var.team_name
}

resource "runscope_bucket" "my_bucket" {
  name      = "[Terraform] Example Bucket"
  team_uuid = data.runscope_team.my_team.id
}
