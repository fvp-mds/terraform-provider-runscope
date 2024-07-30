package schema

type StepRequest struct {
	ID            string              `json:"id"`
	StepType      string              `json:"step_type"`
	Method        string              `json:"method"`
	URL           string              `json:"url"`
	Variables     []StepVariable      `json:"variables"`
	Assertions    []StepAssertion     `json:"assertions"`
	Headers       map[string][]string `json:"headers"`
	Auth          StepAuth            `json:"auth"`
	Body          string              `json:"body"`
	Form          map[string][]string `json:"form"`
	Scripts       []string            `json:"scripts"`
	BeforeScripts []string            `json:"before_scripts"`
	Note          string              `json:"note"`
	Skipped       bool                `json:"skipped"`
}

type StepSubtest struct {
	ID                   string          `json:"id"`
	TestUUID             string          `json:"test_uuid"`
	EnvironmentUUID      string          `json:"environment_uuid"`
	BucketKey            string          `json:"bucket_key"`
	UseParentEnvironment bool            `json:"use_parent_environment"`
	Variables            []StepVariable  `json:"variables"`
	Assertions           []StepAssertion `json:"assertions"`
}

type StepCondition struct {
	ID                   string          `json:"id"`
	TestUUID             string          `json:"test_uuid"`
	EnvironmentUUID      string          `json:"environment_uuid"`
	BucketKey            string          `json:"bucket_key"`
	UseParentEnvironment bool            `json:"use_parent_environment"`
	LeftValue            string          `json:"left_value"`
	Comparison           string          `json:"comparison"`
	RightValue           string          `json:"right_value"`
	Variables            []StepVariable  `json:"variables"`
	Assertions           []StepAssertion `json:"assertions"`
}

type StepVariable struct {
	Name     string `json:"name"`
	Property string `json:"property"`
	Source   string `json:"source"`
}

type StepAssertion struct {
	Source     string `json:"source"`
	Property   string `json:"property"`
	Comparison string `json:"comparison"`
	Value      string `json:"value"`
}

type StepAuth struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	AuthType string `json:"auth_type,omitempty"`
}

type StepCreateRequestRequest struct {
	StepRequest
	StepType string `json:"step_type"`
}

type StepCreateRequestResponse struct {
	Step []StepRequest `json:"data"`
}

type StepGetRequestResponse struct {
	Step StepRequest `json:"data"`
}

type StepGetSubstepResponse struct {
	Step StepSubtest `json:"data"`
}
type StepGetConditionResponse struct {
	Step StepCondition `json:"data"`
}
type StepUpdateRequestRequest struct {
	StepRequest
	StepType string `json:"step_type"`
}
type StepUpdateRequstResponse struct {
	Step StepRequest `json:"data"`
}

type StepCreateSubtestRequest struct {
	StepSubtest
	StepType string `json:"step_type"`
}
type StepCreateSubtestResponse struct {
	Step []StepSubtest `json:"data"`
}
type StepUpdateSubtestRequest struct {
	StepSubtest
	StepType string `json:"step_type"`
}
type StepUpdateSubtestResponse struct {
	Step StepSubtest `json:"data"`
}
type StepCreateConditionRequest struct {
	StepCondition
	StepType string `json:"step_type"`
}
type StepCreateConditionResponse struct {
	Step []StepCondition `json:"data"`
}
type StepUpdateConditionRequest struct {
	StepCondition
	StepType string `json:"step_type"`
}
type StepUpdateConditionResponse struct {
	Step StepCondition `json:"data"`
}
