package service

import (
	"context"
	"log"

	"github.com/Chirag-Nayak/go-basics/web-service/model"
	"github.com/Chirag-Nayak/go-basics/web-service/repository"
)

type Employee struct {
	repo   repository.Employee
	logger *log.Logger
}

func NewEmployee(l *log.Logger, r repository.Employee) *Employee {
	return &Employee{
		logger: l,
		repo:   r,
	}
}

func (e *Employee) GetAll(ctx context.Context) (model.Employees, error) {
	return e.repo.GetAll(ctx)
}

func (e *Employee) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	return e.repo.GetByID(ctx, id)
}

func (e *Employee) AddEmployee(ctx context.Context, emp *model.Employee) (*model.Employee, error) {
	return e.repo.AddEmployee(ctx, emp)
}

func (e *Employee) UpdateEmployee(ctx context.Context, id int64, emp *model.Employee) (*model.Employee, error) {
	return e.repo.UpdateEmployee(ctx, id, emp)
}
