package request

import (
	"fmt"

	"github.com/gitkeng/microservice"
	"github.com/gitkeng/microservice/util/convutil"
	"github.com/gitkeng/microservice/util/stringutil"
)

var (
	err0001 = microservice.NewError("err0001",
		"first_name is require")
	err0002 = microservice.NewError("err0002",
		"last_name is require")
	err0003 = microservice.NewError("err0003",
		"email is require")
	err0004 = microservice.NewError("err0004",
		"department is require")
	err0005 = func(salary float64) microservice.Error {
		return microservice.NewError("err0005", fmt.Sprintf("invalid salary: %.2f", salary))
	}
)

type CreateEmployeeRequest struct {
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Email      string  `json:"email"`
	Age        int     `json:"age"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

func (emp *CreateEmployeeRequest) String() string {
	return stringutil.Json(*emp)
}

func (emp *CreateEmployeeRequest) ToMap() map[string]any {
	return convutil.Obj2Map(*emp)
}

func (emp *CreateEmployeeRequest) Validate() error {
	var errs microservice.Errors
	if stringutil.IsEmptyString(emp.FirstName) {
		errs = append(errs, err0001)
	}
	if stringutil.IsEmptyString(emp.LastName) {
		errs = append(errs, err0002)
	}
	if stringutil.IsEmptyString(emp.Email) {
		errs = append(errs, err0003)
	}
	if stringutil.IsEmptyString(emp.Department) {
		errs = append(errs, err0004)
	}
	if emp.Salary <= 0 {
		errs = append(errs, err0005(emp.Salary))
	}
	if len(errs) > 0 {
		return &errs
	}
	return nil
}
