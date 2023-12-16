package controller

import (
	"example/snap_roles/cmd/pkg/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobApiControllerInterface interface {
	GetMicrosoftJobs(c *gin.Context) interface{}
	GetAppleJobs(c *gin.Context) interface{}
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
// @Summary      Get microsoft Jobs
// @Tags         Jobs
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

	c.JSON(http.StatusOK, resp)

}

// GetAppleJobs godoc
// @Summary      Get Apple Jobs
// @Tags         Jobs
// @Description  get apple jobs
// @Produce      json
// @Success      200  {object}  model.MicrosoftResponse
// @Router       /api/v1/apple/ [get]
func (j *JobApiControllerStruct) GetAppleJobs(c *gin.Context) {
	resp, err := j.Api.GetAppleJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
