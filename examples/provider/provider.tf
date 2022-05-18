resource "runscope_bucket" "my_bucket" {
  name      = "My Bucket"
  team_uuid = var.team_uuid
}

resource "runscope_test" "my_test" {
  name      = "My Test"
  bucket_id = runscope_bucket.my_bucket.id
}

resource "runscope_step_request" "my_req" {
  bucket_id = runscope_bucket.my_bucket.id
  test_id   = runscope_test.my_test.id
  url       = "https://www.google.com"
  method    = "GET"
  assertion {
    source     = "response_status"
    comparison = "equal_number"
    value      = "200"
  }
}
