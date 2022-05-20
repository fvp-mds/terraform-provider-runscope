resource "runscope_test" "my_test" {
  name        = "My Test"
  description = "Synthethic API test for some core functionality"
  bucket_id   = runscope_bucket.my_bucket.id
}

resource "runscope_step_request" "my_request" {
  bucket_id = runscope_bucket.my_bucket.id
  test_id   = runscope_test.my_test.id
  url       = "https://www.google.com"
  note      = "Check google.com availability"
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
