package api

import (
	"example/snap_roles/cmd/model"
	"example/snap_roles/cmd/pkg/api/handlers"
	"example/snap_roles/cmd/pkg/api/helper"
	"example/snap_roles/internal/constants"
	"io/ioutil"
	"log"
)

type CompanyApiInterface interface {
	GetMicrosoftJobs() []byte
	GetAppleJobs() interface{}
}

type CompanyApiStruct struct {
	constant constants.CompanyStruct
	handler  handlers.HandlerStruct
}

func (c *CompanyApiStruct) GetMicrosoftJobs() []byte {
	microsoftStruct := c.constant.GetMicrosoftData()
	url := microsoftStruct.GenerateGetURL()
	resp, err := c.handler.GetApi(url)
	if err != nil {
		log.Fatalf("Microsoft API error:%s", err)
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatalf("Microsoft API error. Error reading response body: %s", err)
		return nil
	}
	var microsoftApiResponse model.MicrosoftApiResponse
	helper.UnmarshalJSON(body, &microsoftApiResponse)
	marshel, _ := helper.MarshalJSON(microsoftApiResponse)
	return marshel
}

func (c *CompanyApiStruct) GetAppleJobs(interface{}) {
	// appleStruct := c.constant.GetAppleData()
	// resp, err := c.handler.PostApi(appleStruct.URL, "application/json", constants.AppleJobsBodyStruct)

}
