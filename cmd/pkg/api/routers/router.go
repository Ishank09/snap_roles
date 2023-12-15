package routers

import (
	"example/snap_roles/cmd/pkg/api"
	"example/snap_roles/cmd/pkg/api/controller"
	"example/snap_roles/cmd/pkg/api/handlers"
	"example/snap_roles/internal/constants"

	"github.com/gorilla/mux"
)

type RouterInterface interface {
	initializeRouter()
}
type RouterStruct struct {
	companyApiStruct api.CompanyApiStruct
	constant         constants.CompanyStruct
	handler          handlers.HandlerStruct
	controller       controller.JobApiControllerStruct
}

func (rr *RouterStruct) initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/microsoft", rr.controller.GetMicrosoftJobs).Methods("GET")
}
