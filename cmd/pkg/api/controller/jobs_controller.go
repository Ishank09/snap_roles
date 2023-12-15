package controller

import (
	"example/snap_roles/cmd/pkg/api"
	"net/http"
)

type JobApiControllerInterface interface {
	GetMicrosoftJobs() interface{}
}

type JobApiControllerStruct struct {
	Api api.CompanyApiStruct
}

func (j *JobApiControllerStruct) GetMicrosoftJobs(w http.ResponseWriter, r *http.Request) {
	resp, err := j.Api.GetMicrosoftJobs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
