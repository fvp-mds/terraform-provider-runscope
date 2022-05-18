package schema

type Bucket struct {
	Key        string `json:"key"`
	Name       string `json:"name"`
	Team       Team   `json:"team"`
	AuthToken  string `json:"auth_token"`
	Default    bool   `json:"default"`
	VerifySSL  bool   `json:"verify_ssl"`
	TriggerURL string `json:"trigger_url"`
}

type BucketCreateResponse struct {
	Bucket `json:"data"`
}

type BucketGetResponse struct {
	Bucket `json:"data"`
}

type BucketListResponse struct {
	Buckets []Bucket `json:"data"`
}
