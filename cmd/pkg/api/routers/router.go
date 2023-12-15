package routers

import (
	"example/snap_roles/cmd/pkg/api"
	"example/snap_roles/cmd/pkg/api/controller"
	"example/snap_roles/cmd/pkg/api/handlers"
	"example/snap_roles/internal/constants"
	"net/http"

	"github.com/gorilla/mux"
)

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
	r := mux.NewRouter()
	r.HandleFunc("/microsoft", rr.Controller.GetMicrosoftJobs).Methods("GET")
	http.ListenAndServe(":8080", r)

}
