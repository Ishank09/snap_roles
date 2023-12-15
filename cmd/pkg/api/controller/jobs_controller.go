package controller

import (
	"encoding/json"
	"example/snap_roles/cmd/pkg/api"
	"net/http"
)

type JobApiControllerInterface interface {
	GetMicrosoftJobs() interface{}
}

type JobApiControllerStruct struct {
	api api.CompanyApiStruct
}

func (j *JobApiControllerStruct) GetMicrosoftJobs(w http.ResponseWriter, r *http.Request) {
	resp := j.api.GetMicrosoftJobs()
	json.NewEncoder(w).Encode(resp)
}
