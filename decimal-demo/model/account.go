package model

import (
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Store account information
type Account struct {
	ID           int64
	MongoID      primitive.ObjectID `bson:"_id,omitempty"`
	AccountName  string             `bson:"accountName"`
	CurrencyName string             `bson:"currencyName"`
	Balance      decimal.Decimal    `bson:"balance"`
}

// Accounts is a collectio of Account
type Accounts []*Account
