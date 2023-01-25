package model

import "github.com/shopspring/decimal"

// Store account information
type Account struct {
	ID           int64
	AccountName  string
	CurrencyName string
	Balance      decimal.Decimal
}

// Accounts is a collectio of Account
type Accounts []*Account
