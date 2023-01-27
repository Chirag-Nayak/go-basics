package service

import (
	"context"
	"errors"
	"log"

	"github.com/Chirag-Nayak/go-basics/decimal-demo/model"
	"github.com/Chirag-Nayak/go-basics/decimal-demo/repository"
)

var (
	ErrInvalidAccountInfo = errors.New("invalid account details given in argument")
	ErrInvalidContext     = errors.New("invalid context given in the argument")
	ErrInvalidAccountId   = errors.New("invalid account ID given in the argument")
)

type Account struct {
	logger *log.Logger
	repo   repository.Account
}

func NewAccountService(l *log.Logger, r repository.Account) *Account {
	return &Account{
		logger: l,
		repo:   r,
	}
}

func (a *Account) GetAccounts(ctx context.Context) (model.Accounts, error) {

	// Validate & Process input parameters to the service
	if ctx == nil {
		return nil, ErrInvalidContext
	}

	// Perform operation using the repository
	acs, err := a.repo.GetAll(ctx)
	if err != nil {
		// Log the error for service module also in case seperate logging is needed
		a.logger.Printf("Error in Account.GetAccounts from repository: %s", err)
		return nil, err
	}
	return acs, nil
}

func (a *Account) GetAccountById(ctx context.Context, id int64) (*model.Account, error) {
	// Validate & Process input parameters to the service
	if ctx == nil {
		return nil, ErrInvalidContext
	}
	if id <= 0 {
		return nil, ErrInvalidAccountId
	}

	// Perform operation using the repository
	acc, err := a.repo.GetById(ctx, id)
	if err != nil {
		a.logger.Printf("Error in Account.GetAccountById from repository: %s", err)
		return nil, err
	}
	return acc, nil
}

func (a *Account) AddAccount(ctx context.Context, acc model.Account) (*model.Account, error) {
	// Validate & Process input parameters to the service
	if ctx == nil {
		return nil, ErrInvalidContext
	}
	if acc == (model.Account{}) {
		return nil, ErrInvalidAccountId
	}

	addedAcc, err := a.repo.AddAccount(ctx, acc)
	if err != nil {
		a.logger.Printf("Error in Account.AddAccount from repository: %s", err)
		return nil, err
	}
	return addedAcc, nil
}

func (a *Account) UpdateAccount(ctx context.Context, id int64, acc model.Account) (*model.Account, error) {
	// Validate & Process input parameters to the service
	if ctx == nil {
		return nil, ErrInvalidContext
	}
	if id <= 0 {
		return nil, ErrInvalidAccountId
	}
	if acc == (model.Account{}) {
		return nil, ErrInvalidAccountInfo
	}

	// Perform operation using repository
	updatedAcc, err := a.repo.UpdateAccount(ctx, id, acc)
	if err != nil {
		a.logger.Printf("Error in Account.UpdateAccount from repository: %s", err)
		return nil, err
	}
	return updatedAcc, err
}

func (a *Account) DeleteAccount(ctx context.Context, id int64) error {
	// Validate & Process input parameters to the service
	if ctx == nil {
		return ErrInvalidContext
	}
	if id <= 0 {
		return ErrInvalidAccountId
	}

	err := a.repo.DeleteAccount(ctx, id)
	if err != nil {
		a.logger.Printf("Error in Account.DeleteAccount from repository: %s", err)
	}
	return err
}
