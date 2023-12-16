package api

import (
	"bytes"
	"encoding/json"
	"example/snap_roles/cmd/model"
	"example/snap_roles/cmd/pkg/api/handlers"
	"example/snap_roles/internal/constants"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type CompanyApiInterface interface {
	GetMicrosoftJobs() []byte
	GetAppleJobs() interface{}
}

type CompanyApiStruct struct {
	Constant constants.CompanyStruct
	Handler  handlers.HandlerStruct
}

// UnmarshalJSON is a generic function to unmarshal JSON data into a provided data structure
func unmarshalMicrosoftJSON(jsonData []byte, target *model.MicrosoftApiResponse) error {
	if err := json.Unmarshal(jsonData, target); err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err)
		return err
	}
	return nil
}

func checkMicrosoftPostedDate(timestampString string) bool {
	timestamp, err := time.Parse(time.RFC3339, timestampString)
	if err != nil {
		log.Fatalf("Unable to parse time, Microsoft API")
		return false
	}

	// Calculate the duration between the current time and the timestamp
	duration := time.Since(timestamp)

	// Check if the duration is within the last 15 minutes
	// if duration <= 3*24*time.Hour {

	// Assuming duration is a time.Duration variable
	if duration <= time.Duration(constants.MicrosoftPingSecs)*time.Second {
		return true
	} else {
		return false
	}

}

func getLatestMicrosoftJobs(resp model.MicrosoftApiResponse) []model.MicrosoftResponse {
	jobs := resp.OperationResult.Result.Jobs
	var response []model.MicrosoftResponse
	for _, job := range jobs {
		var one_job model.MicrosoftResponse
		one_job.JobId = job.JobId
		one_job.PostingDate = job.PostingDate
		one_job.Properties = job.Properties
		one_job.Properties.Description = ""
		one_job.Title = job.Title
		one_job.PostedUrl = fmt.Sprintf(constants.MicrosoftJobShareUrl, job.JobId)
		if checkMicrosoftPostedDate(one_job.PostingDate) {
			response = append(response, one_job)
		}

	}
	return response

}

func (c *CompanyApiStruct) GetMicrosoftJobs() ([]byte, error) {
	microsoftStruct := c.Constant.GetMicrosoftData()
	url := microsoftStruct.GenerateGetURL()
	resp, err := c.Handler.GetApi(url)
	if err != nil {
		return nil, fmt.Errorf("Microsoft API error: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Microsoft API error. Error reading response body: %s", err)
	}

	var microsoftApiResponse model.MicrosoftApiResponse
	err = unmarshalMicrosoftJSON(body, &microsoftApiResponse)
	if err != nil {
		return nil, fmt.Errorf("Error during unmarshaling: %s", err)
	}
	constructedResponse := getLatestMicrosoftJobs(microsoftApiResponse)

	marshaledJSON, err := json.Marshal(constructedResponse)
	if err != nil {
		return nil, fmt.Errorf("Error during marshaling: %s", err)
	}

	return marshaledJSON, nil
}

func (c *CompanyApiStruct) GetAppleJobs() ([]byte, error) {
	appleStruct := c.Constant.GetAppleData()
	// Assuming appleStruct.Body is of type interface{}
	bodyBytes, ok := appleStruct.Body.([]byte)
	if !ok {
		// Handle the case where the assertion fails
		// For example, you might return an error or use a default value
		bodyBytes = []byte("default body")
	}
	resp, err := c.Handler.PostApi(appleStruct.URL, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("Microsoft API error: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Microsoft API error. Error reading response body: %s", err)
	}

	var microsoftApiResponse model.MicrosoftApiResponse
	err = unmarshalMicrosoftJSON(body, &microsoftApiResponse)
	if err != nil {
		return nil, fmt.Errorf("Error during unmarshaling: %s", err)
	}
	constructedResponse := getLatestMicrosoftJobs(microsoftApiResponse)

	marshaledJSON, err := json.Marshal(constructedResponse)
	if err != nil {
		return nil, fmt.Errorf("Error during marshaling: %s", err)
	}

	return marshaledJSON, nil

}
