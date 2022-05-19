package schema

type Account struct {
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
	Email string `json:"email"`
	Teams []Team `json:"teams"`
}

type AccountResponse struct {
	Data Account `json:"data"`
}
