terraform {
  required_providers {
	  runscope = {
		  source = "terraform.storytel.com/storytel/runscope"
		  version = "0.1.0"
	  }
  }
}

provider "runscope" {
	access_token = "6cc13a12-c975-4f8b-b574-46b15b076bd0"
}

resource "runscope_bucket" "my_bucket" {
  name      = "Storytel/terraform-provider-runscope test"
  team_uuid = "ff1d7a6c-8b46-4a9c-9096-647aa7033990"
}

resource "runscope_environment" "my_env" {
  bucket_id = runscope_bucket.my_bucket.id
  name = "My Env"
}

resource "runscope_test" "my_test" {
  name        = "My Test!"
  description = "desc"
  bucket_id   = runscope_bucket.my_bucket.id
}

resource "runscope_step_request" "my_req" {
  bucket_id = runscope_bucket.my_bucket.id
  test_id = runscope_test.my_test.id
  url = "https://www.google.com"
  method = "POST"
  assertion {
    source     = "response_status"
    comparison = "equal_number"
    value      = "200"
  }
}

resource "runscope_step_subtest" "my_sub" {
  depends_on = [
    runscope_step_request.my_req
  ]
  bucket_id = runscope_bucket.my_bucket.id
  test_id = runscope_test.my_test.id
  source_bucket_id = "22rxkx2dqinc"
  source_test_id = "adb2f6f2-ae03-4842-8ff0-a386961aee8f"
  source_environment_id = "3e21cc35-df5a-4637-8e82-0a01a35a074"
}

# resource "runscope_schedule" "my_sched" {
#   bucket_id = runscope_bucket.my_bucket.id
#   test_id = runscope_test.my_test.id
#   environment_id = runscope_environment.my_env.id
#   interval = "1d"
# }
