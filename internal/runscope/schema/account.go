package schema

type AccountTeam struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type Account struct {
	Name  string        `json:"name"`
	UUID  string        `json:"uuid"`
	Email string        `json:"email"`
	Teams []AccountTeam `json:"teams"`
}

type AccountResponse struct {
	Data Account `json:"data"`
}
