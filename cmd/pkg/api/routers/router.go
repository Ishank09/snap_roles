package routers

import (
	"example/snap_roles/cmd/pkg/api"
	"example/snap_roles/cmd/pkg/api/controller"

	"example/snap_roles/cmd/pkg/api/handlers"
	"example/snap_roles/internal/constants"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	// gin-swagger middleware
	swaggerFiles "github.com/swaggo/files"
)

// swagger embed files
type RouterInterface interface {
	InitializeRouter()
}
type RouterStruct struct {
	CompanyApiStruct api.CompanyApiStruct
	Constant         constants.CompanyStruct
	Handler          handlers.HandlerStruct
	Controller       controller.JobApiControllerStruct
}

func (rr *RouterStruct) InitializeRouter() {
	r := gin.Default()
	constant := &constants.CompanyStruct{}
	handler := &handlers.HandlerStruct{}

	companyApiStruct := &api.CompanyApiStruct{
		Constant: *constant,
		Handler:  *handler,
	}
	c := controller.NewJobApiController(*companyApiStruct)

	v1 := r.Group("/api/v1")
	{
		jobs := v1.Group("/microsoft")
		{
			jobs.GET("", c.GetMicrosoftJobs)
		}
		appleJobs := v1.Group("/apple")
		{
			appleJobs.GET("", c.GetAppleJobs)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")

}
