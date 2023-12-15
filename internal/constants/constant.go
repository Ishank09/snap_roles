package constants

import (
	"fmt"
	"net/url"
)

type RequestType string

const (
	MicrosoftJobUrl = "https://gcsservices.careers.microsoft.com/search/api/v1/search"
	AppleJobUrl     = "https://jobs.apple.com/api/role/search"
)

const (
	Get  RequestType = "Get"
	Post RequestType = "Post"
)

type ConstantsInterface interface {
	GenerateGetURL()
	GetMicrosoftData()
	GetAppleData()
}

type CompanyStruct struct {
	CompanyName string
	RequestType RequestType
	URL         string
	Parameters  map[string][]string
	Body        interface{}
}

func (c *CompanyStruct) GenerateGetURL() string {
	baseURL := c.URL
	params := c.Parameters
	queryParams := url.Values{}

	// Add parameters to the URL
	for key, values := range params {
		for _, value := range values {
			queryParams.Add(key, value)
		}
	}

	// Construct the final URL
	finalURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())
	return finalURL
}

func (c *CompanyStruct) GetMicrosoftData() CompanyStruct {
	return MicrosoftJobs
}
func (*CompanyStruct) GetAppleData() CompanyStruct {
	return AppleJobs
}

var MicrosoftJobs = CompanyStruct{
	CompanyName: "Microsoft Jobs",
	RequestType: Get,
	URL:         MicrosoftJobUrl,
	Parameters: map[string][]string{
		"lc":   {"United States"},
		"p":    {"Engineering", "Research, Applied, & Data Sciences", "Software Engineering"},
		"d":    {"Applied Sciences", "Data Science", "Research Sciences", "Software Engineering"},
		"exp":  {"Students and graduates"},
		"rt":   {"Individual Contributor"},
		"et":   {"Full-Time", "Internship", "Post Doc Research"},
		"ws":   {"Microsoft on-site only", "Up to 50% work from home"},
		"el":   {"Bachelors", "Doctorate"},
		"l":    {"en_us"},
		"pg":   {"1"},
		"pgSz": {"20"},
		"o":    {"Relevance"},
		"flt":  {"true"},
	},
	Body: nil,
}

var AppleJobs = CompanyStruct{
	CompanyName: "Apple Inc.",
	RequestType: Post,
	URL:         AppleJobUrl,
	Parameters:  nil,
	Body: AppleJobsBodyStruct{
		Query: "internship",
		Filters: AppleJobsFilterStruct{
			Range: AppleJobsFilterRangeStruct{},
		},
		Page:   1,
		Locale: "en-us",
		Sort:   "relevance",
	},
}

type AppleJobsFilterRangeStruct struct {
	StandardWeeklyHours struct {
		Start interface{} `json:"start"`
		End   interface{} `json:"end"`
	} `json:"standardWeeklyHours"`
}

type AppleJobsFilterStruct struct {
	Range AppleJobsFilterRangeStruct `json:"range"`
}

type AppleJobsBodyStruct struct {
	Query   string                `json:"query"`
	Filters AppleJobsFilterStruct `json:"filters"`
	Page    int                   `json:"page"`
	Locale  string                `json:"locale"`
	Sort    string                `json:"sort"`
}
