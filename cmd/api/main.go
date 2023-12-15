package main

import (
	"example/snap_roles/cmd/pkg/api"
	"example/snap_roles/cmd/pkg/api/controller"
	"example/snap_roles/cmd/pkg/api/handlers"
	"example/snap_roles/cmd/pkg/api/routers"
	_ "example/snap_roles/docs"
	"example/snap_roles/internal/constants"
)

func main() {
	// Creating an instance of *routers.RouterStruct
	constant := &constants.CompanyStruct{}
	handler := &handlers.HandlerStruct{}

	companyApiStruct := &api.CompanyApiStruct{
		Constant: *constant,
		Handler:  *handler,
	}

	routerInstance := &routers.RouterStruct{
		CompanyApiStruct: *companyApiStruct,
		Constant:         *constant,
		Handler:          *handler,
		Controller: controller.JobApiControllerStruct{
			Api: *companyApiStruct,
		},
	}

	routerInstance.InitializeRouter()
}
