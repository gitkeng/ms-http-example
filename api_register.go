package main

import (
	"github.com/gitkeng/microservice"
	"ms-http/handler"
)

const (
	CreateEmployeeEndPoint = "/v1/employee"
	FilterEmployeeEndPoint = "/v1/employee/filter"
)

func registerEmployeeHandler(ms microservice.IMicroservice) {
	ms.POST(CreateEmployeeEndPoint, handler.CreateEmployeeHandler)
	ms.POST(FilterEmployeeEndPoint, handler.FilterEmployeeHandler)
}
