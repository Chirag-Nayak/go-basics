package repository

import (
	"context"

	"github.com/Chirag-Nayak/go-basics/decimal-demo/model"
)

type Account interface {
	GetAll(context.Context) (model.Accounts, error)
}
