package model

type MicrosoftApiResponse struct {
	operationResult OperationResult `json:"operationResult"`
	errorInfo       string          `json:"errorInfo"`
}
type OperationResult struct {
	result    Result `json:"result"`
	status    string `json:"status"`
	quality   string `json:"quality"`
	errorcode string `json:"errorcode"`
}
type Result struct {
	searchId  string `json:"searchId"`
	totalJobs int64  `json:"totalJobs"`
	jobs      []Jobs `json:"jobs"`
	id        int64  `json:"id"`
}
type Jobs struct {
	jobId       string       `json:"jobId"`
	title       string       `json:"title"`
	postingDate string       `json:"postingDate"`
	properties  []Properties `json:"properties"`
}
type Properties struct {
	description     string   `json:"description"`
	locations       []string `json:"locations"`
	primaryLocation string   `json:"primaryLocation"`
	profession      string   `json:"profession"`
	jobType         string   `json:"jobType"`
	roleType        string   `json:"roleType"`
	employmentType  string   `json:"employmentType"`
	educationLevel  string   `json:"educationLevel"`
}
