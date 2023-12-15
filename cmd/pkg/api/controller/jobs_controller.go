package controller

import (
	"example/snap_roles/cmd/pkg/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobApiControllerInterface interface {
	GetMicrosoftJobs(c *gin.Context) interface{}
}

// NewController creates a new instance of the Controller.
func NewJobApiController(companyApiImplementation api.CompanyApiStruct) *JobApiControllerStruct {
	return &JobApiControllerStruct{
		Api: companyApiImplementation,
	}
}

type JobApiControllerStruct struct {
	Api api.CompanyApiStruct
}

// GetMicrosoftJobs godoc
// @Summary      Show an account
// @Description  get microsoft jobs
// @Produce      json
// @Success      200  {object}  model.MicrosoftResponse
// @Router       /api/v1/microsoft/ [get]
func (j *JobApiControllerStruct) GetMicrosoftJobs(c *gin.Context) {

	resp, err := j.Api.GetMicrosoftJobs()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(resp))

}
