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

type CompanyStruct struct {
	CompanyName string
	RequestType RequestType
	URL         string
	Parameters  map[string]string
	Body        interface{}
}

func (c *CompanyStruct) GenerateGetURL() string {
	u, err := url.Parse(c.URL)
	if err != nil {
		panic(err)
	}

	q := u.Query()
	for key, value := range c.Parameters {
		q.Add(key, value)
	}

	u.RawQuery = q.Encode()

	return u.String()
}

var MicrosoftJobs = CompanyStruct{
	CompanyName: "Microsoft Jobs",
	RequestType: Get,
	URL:         MicrosoftJobUrl,
	Parameters: map[string]string{
		"lc":   "United States",
		"p":    "Engineering",
		"d":    "Applied Sciences",
		"exp":  "Students and graduates",
		"ws":   "Microsoft on-site only",
		"el":   "Bachelors",
		"l":    "en_us",
		"pg":   "1",
		"pgSz": "20",
		"o":    "Relevance",
		"flt":  "true",
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

func main() {
	// Example usage
	url := MicrosoftJobs.GenerateGetURL()
	fmt.Println("Microsoft Jobs URL:", url)
}
