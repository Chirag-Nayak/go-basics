package service

import (
	"context"
	"log"
	"time"

	"github.com/Chirag-Nayak/go-basics/decimal-demo/model"
	"github.com/Chirag-Nayak/go-basics/decimal-demo/repository"
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

func (w *Account) GetAccounts() (model.Accounts, error) {

	// Creating the context with time out of 10 seconds to get data form repository
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	acs, err := w.repo.GetAll(ctx)
	if err != nil {
		w.logger.Printf("Error from repository: %s", err)
		return nil, err
	}
	return acs, nil
}
