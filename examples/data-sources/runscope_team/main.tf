terraform {
  required_providers {
    runscope = {
      source  = "Storytel/runscope"
      version = ">= 0.14.0"
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
