package handler

import (
	"fmt"
	"ms-http/datastore"
	"ms-http/request"
	"ms-http/response"
	"net/http"
	"strconv"

	"github.com/gitkeng/microservice"
	"github.com/gitkeng/microservice/util/uuid"
	"github.com/jinzhu/copier"
)

const (
	tagCreateEmployeeHandler = "CreateEmployeeHandler"
	tagFilterEmployeeHandler = "FilterEmployeeHandler"
)

func CreateEmployeeHandler(ctx microservice.IContext) error {
	req := &request.CreateEmployeeRequest{}
	if err := ctx.Bind(req); err != nil {
		return ctx.Response(
			microservice.ErrorLevel,
			tagCreateEmployeeHandler,
			http.StatusBadRequest,
			strconv.Itoa(http.StatusBadRequest),
			"invalid req",
			err,
		)
	}

	// copy req
	empEntity := &datastore.Employee{}
	if err := copier.Copy(empEntity, req); err != nil {
		return ctx.Response(
			microservice.ErrorLevel,
			tagCreateEmployeeHandler,
			http.StatusInternalServerError,
			strconv.Itoa(http.StatusInternalServerError),
			err.Error(),
			err,
		)
	}

	// set another require field
	empEntity.EmployeeCode = uuid.NewUUID()

	if emp, err := datastore.InsertEmployee(ctx, empEntity); err != nil {
		return ctx.Response(
			microservice.ErrorLevel,
			tagCreateEmployeeHandler,
			http.StatusInternalServerError,
			strconv.Itoa(http.StatusInternalServerError),
			"insert employee failed",
			err,
		)
	} else {
		//generate response
		resp := &response.EmployeeResponse{}
		if err := copier.Copy(resp, emp); err != nil {
			return ctx.Response(
				microservice.ErrorLevel,
				tagCreateEmployeeHandler,
				http.StatusInternalServerError,
				strconv.Itoa(http.StatusInternalServerError),
				fmt.Sprintf("generate response fail for employee %s", emp.EmployeeCode),
				err,
			)
		}

		return ctx.Response(
			microservice.DebugLevel,
			tagCreateEmployeeHandler,
			http.StatusOK,
			strconv.Itoa(http.StatusOK),
			"create employee success",
			nil,
			microservice.Field{
				Key:   "employee",
				Value: resp,
			},
		)
	}

}

func FilterEmployeeHandler(ctx microservice.IContext) error {
	req := &microservice.FilterRequest{}
	if err := ctx.Bind(req); err != nil {
		return ctx.Response(
			microservice.ErrorLevel,
			tagFilterEmployeeHandler,
			http.StatusBadRequest,
			strconv.Itoa(http.StatusBadRequest),
			"invalid req",
			err,
		)
	}

	if employees, total, err := datastore.GetEmployees(ctx, req.GetFilters(), req.GetOption()); err != nil {

		return ctx.Response(
			microservice.ErrorLevel,
			tagFilterEmployeeHandler,
			http.StatusInternalServerError,
			strconv.Itoa(http.StatusInternalServerError),
			"filter employee failed",
			err,
		)
	} else {

		//generate response response
		resp := make([]*response.EmployeeResponse, 0)
		for idx, _ := range employees {
			emp := &response.EmployeeResponse{}
			if err := copier.Copy(emp, employees[idx]); err != nil {
				return ctx.Response(
					microservice.ErrorLevel,
					tagFilterEmployeeHandler,
					http.StatusInternalServerError,
					strconv.Itoa(http.StatusInternalServerError),
					fmt.Sprintf("generate response fail for employee %s", employees[idx].EmployeeCode),
					err,
				)
			}
			resp = append(resp, emp)
		}

		return ctx.Response(
			microservice.DebugLevel,
			tagFilterEmployeeHandler,
			http.StatusOK,
			strconv.Itoa(http.StatusOK),
			"filter employee success",
			nil,
			microservice.Field{
				Key:   "employees",
				Value: resp,
			},
			microservice.Field{
				Key:   "total",
				Value: total,
			},
		)
	}

}
