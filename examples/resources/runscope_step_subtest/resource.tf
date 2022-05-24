resource "runscope_test" "my_substep_test" {
  name      = "My Test With Substep"
  bucket_id = runscope_bucket.my_bucket.id
}

resource "runscope_step_request" "my_request" {
  bucket_id = runscope_bucket.my_bucket.id
  test_id   = runscope_test.my_substep_test.id
  url       = "https://www.google.com"
  note      = "Use as substep"
  method    = "GET"

  header {
    header = "X-My-Custom-Header"
    value  = "charge"
  }

  assertion {
    source     = "response_status"
    comparison = "equal_number"
    value      = "200"
  }
}

resource "runscope_test" "my_test" {
  name      = "My Test"
  bucket_id = runscope_bucket.my_bucket.id
}

resource "runscope_step_subtest" "my_subtest" {
  bucket_id              = runscope_bucket.my_bucket.id
  test_id                = runscope_test.my_test.id
  source_bucket_id       = runscope_bucket.my_bucket.id
  source_test_id         = runscope_test.my_substep_test.id
  use_parent_environment = true
}
