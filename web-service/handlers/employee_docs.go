// Package classification Employee API
//
// Documentation for Employee API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "github.com/Chirag-Nayak/go-basics/web-service/model"

// swagger:route GET /employee Employees listEmployee
// Return list of employees from the database
// responses:
//	200: employeesResponse

// swagger:route GET /employee/{id} Employees getEmployee
// Return a single employee information from the database
// responses:
//	200: employeeResponse
//	400: errorResponse
//	500: errorResponse

// A list of employees
// swagger:response employeesResponse
type employeesResponseWrapper struct {
	// All current employees
	// in: body
	Body []model.Employee
}

// Data structure representing a single employee
// swagger:response employeeResponse
type employeeResponseWrapper struct {
	// Newly created employee
	// in: body
	Body model.Employee
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// swagger:parameters updateEmployee getEmployee
type employeeIDParamsWrapper struct {
	// The id of the employee for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters addEmployee updateEmployee
type employeeParamsWrapper struct {
	// Employee data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body model.Employee
}
