package model

type MicrosoftApiResponse struct {
	OperationResult OperationResult `json:"operationResult"`
	ErrorInfo       string          `json:"errorInfo"`
}
type OperationResult struct {
	Result    Result `json:"result"`
	Status    string `json:"status"`
	Quality   string `json:"quality"`
	Errorcode string `json:"errorcode"`
}
type Result struct {
	SearchId  string `json:"searchId"`
	TotalJobs int64  `json:"totalJobs"`
	Jobs      []Jobs `json:"jobs"`
	Id        int64  `json:"id"`
}
type Jobs struct {
	JobId       string     `json:"jobId"`
	Title       string     `json:"title"`
	PostingDate string     `json:"postingDate"`
	Properties  Properties `json:"properties"`
}
type Properties struct {
	Description     string   `json:"description"`
	Locations       []string `json:"locations"`
	PrimaryLocation string   `json:"primaryLocation"`
	Profession      string   `json:"profession"`
	JobType         string   `json:"jobType"`
	RoleType        string   `json:"roleType"`
	EmploymentType  string   `json:"employmentType"`
	EducationLevel  string   `json:"educationLevel"`
}

type MicrosoftResponse struct {
	JobId       string     `json:"jobId"`
	Title       string     `json:"title"`
	PostingDate string     `json:"postingDate"`
	PostedUrl   string     `json:"postedUrl"`
	Properties  Properties `json:"properties"`
}
