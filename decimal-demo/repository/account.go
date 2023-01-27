package repository

import (
	"context"
	"errors"

	"github.com/Chirag-Nayak/go-basics/decimal-demo/model"
)

var (
	ErrDuplicateRecord      = errors.New("record already exists")
	ErrDuplicateAccountName = errors.New("record with same account name already exist")
	ErrRecordNotExist       = errors.New("record does not exist")
	ErrInvalidAccountID     = errors.New("invalid account ID")
	ErrUpdateFailed         = errors.New("failed to update the record")
	ErrDeleteFailed         = errors.New("failed to delete the record")
)

const (
	DB_CONST_CONSTRAINT_UNQ_ACC_NAME = "account_name_must_be_unique"
)

type Account interface {
	GetAll(context.Context) (model.Accounts, error)
	GetById(context.Context, int64) (*model.Account, error)
	AddAccount(context.Context, model.Account) (*model.Account, error)
	UpdateAccount(context.Context, int64, model.Account) (*model.Account, error)
	DeleteAccount(context.Context, int64) error
}
