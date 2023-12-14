package api

import (
	"go/constant"
	"example/snap_roles/internal/constants"
)


type CompanyApiInterface interface {
	GetMicrosoftJobs()
	GetAppleJobs()
}

type CompanyApiStruct struct {
	constant 
}


func 