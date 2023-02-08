package repository

import (
	"context"
	"errors"

	"github.com/Chirag-Nayak/go-basics/web-service/model"
)

var (
	ErrRecordNotFound = errors.New("Employee record not found")
)

type Employee interface {
	GetAll(context.Context) (model.Employees, error)
	GetByID(context.Context, int64) (*model.Employee, error)
	AddEmployee(context.Context, *model.Employee) (*model.Employee, error)
	UpdateEmployee(context.Context, int64, *model.Employee) (*model.Employee, error)
}
