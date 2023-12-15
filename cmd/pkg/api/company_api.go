package api

import (
	"encoding/json"
	"example/snap_roles/cmd/model"
	"example/snap_roles/cmd/pkg/api/handlers"
	"example/snap_roles/internal/constants"
	"fmt"
	"io/ioutil"
	"log"
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

	marshaledJSON, err := json.Marshal(microsoftApiResponse)
	if err != nil {
		return nil, fmt.Errorf("Error during marshaling: %s", err)
	}

	return marshaledJSON, nil
}

func (c *CompanyApiStruct) GetAppleJobs(interface{}) {
	// appleStruct := c.constant.GetAppleData()
	// resp, err := c.handler.PostApi(appleStruct.URL, "application/json", constants.AppleJobsBodyStruct)

}
