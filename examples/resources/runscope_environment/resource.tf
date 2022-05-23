resource "runscope_environment" "my_environment" {
  bucket_id = runscope_bucket.my_bucket.id
  name      = "My Environment"
  initial_variables = {
    VAR1 = "value1"
    VAR2 = "value2"
  }
  header {
    header = "Authorization"
    value  = "bearer very-secret-token"
  }
  regions = [
    "in1",
    "br1",
    "gcpeu3"
  ]
}
