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

type AppleApiResponse struct {
	SearchResults []SearchResults `json:"searchResults"`
	totalRecords  int             `json:totalRecords`
}

type Location struct {
	City          string `json:"city"`
	StateProvince string `json:"stateProvince"`
	CountryName   string `json:"countryName"`
	Metro         string `json:"metro"`
	Region        string `json:"region"`
	Name          string `json:"name"`
	CountryID     string `json:"countryID"`
	Level         int    `json:"level"`
}

type LocaleInfo struct {
	LanguageCode      string `json:"languageCode"`
	DateFormat        string `json:"dateFormat"`
	DefaultLocaleCode string `json:"defaultLocaleCode"`
}

type Team struct {
	TeamName string `json:"teamName"`
	TeamID   string `json:"teamID"`
	TeamCode string `json:"teamCode"`
}

type SearchResults struct {
	ID                      string     `json:"id"`
	JobSummary              string     `json:"jobSummary"`
	Locations               []Location `json:"locations"`
	PositionID              string     `json:"positionId"`
	PostingDate             string     `json:"postingDate"`
	PostingTitle            string     `json:"postingTitle"`
	LocaleInfo              LocaleInfo `json:"localeInfo"`
	PostDateInGMT           string     `json:"postDateInGMT"`
	TransformedPostingTitle string     `json:"transformedPostingTitle"`
	ReqID                   string     `json:"reqId"`
	ManagedPipelineRole     bool       `json:"managedPipelineRole"`
	Team                    Team       `json:"team"`
	HomeOffice              bool       `json:"homeOffice"`
}
