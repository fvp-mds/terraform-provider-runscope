terraform {
  required_providers {
    runscope = {
      source  = "Storytel/runscope"
      version = "~> 0.12.0"
    }
  }
}

provider "runscope" {
  access_token = var.access_token
}

variable "access_token" {
  type      = string
  sensitive = true
}

variable "team_uuid" {
  type = string
}
