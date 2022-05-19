data "runscope_team" "my_team" {
  name = "Storytel Sweden AB"
}

resource "runscope_bucket" "my_bucket" {
  name      = "My Bucket"
  team_uuid = data.runscope_team.my_team.id
}
