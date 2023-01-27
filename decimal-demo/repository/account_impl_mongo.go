package repository

import (
	"context"
	"errors"
	"log"

	"github.com/Chirag-Nayak/go-basics/decimal-demo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountImplMongo struct {
	logger            *log.Logger
	accountCollection *mongo.Collection
}

func NewAccountImplMongo(l *log.Logger, aC *mongo.Collection) *AccountImplMongo {
	return &AccountImplMongo{
		logger:            l,
		accountCollection: aC,
	}
}

func (a *AccountImplMongo) AddAccount(ctx context.Context, acc model.Account) (*model.Account, error) {
	// Insert the document into MongoDB
	res, err := a.accountCollection.InsertOne(ctx, acc)
	if err != nil {
		a.logger.Printf("Error while inserting document into MongoDB: %#+v\n", err)
		return nil, err
	}

	//TODO CURRENT IMPLEMENTATION OF DECIMAL PACKAGE (`decimal.Decimal`) DOES NOT SUPPORT BSON MARSHALLING & UNMARSHALLING,
	// SO DECIMAL READ - WRITE TO MONGODB DOES NOT WORK HERE
	// THERE IS AN OPEN ISSUE IN GITHUB REGARDING THIS:
	// https://github.com/shopspring/decimal/issues/168
	// CAN TRY ONCE THE ISSUE IS FIXES OR IF SOME SUITABLE WORK AROUND IS FOUND

	// Retrieve the document from MongoDB
	var newAcc *model.Account
	query := bson.M{"_id": res.InsertedID}
	if err = a.accountCollection.FindOne(ctx, query).Decode(&newAcc); err != nil {
		return nil, err
	}
	return newAcc, nil
}

func (a *AccountImplMongo) GetAll(ctx context.Context) (model.Accounts, error) {
	query := bson.M{}

	cursor, err := a.accountCollection.Find(ctx, query)
	if err != nil {
		a.logger.Printf("Error while retrieving the data from MongoDB: %#+v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var acs []*model.Account
	for cursor.Next(ctx) {
		acc := &model.Account{}
		err = cursor.Decode(acc)
		if err != nil {
			a.logger.Printf("Error while reading data from cursor: %#+v\n", err)
			return nil, err
		}
		acs = append(acs, acc)
	}
	return acs, nil
}

func (a *AccountImplMongo) GetById(ctx context.Context, id int64) (*model.Account, error) {
	return nil, errors.New("can not implement this method as MongoDB has its own unique ID")
}

func (a *AccountImplMongo) UpdateAccount(ctx context.Context, id int64, acc model.Account) (*model.Account, error) {
	return nil, errors.New("can not implement this method as MongoDB has its own unique ID")
}

func (a *AccountImplMongo) DeleteAccount(ctx context.Context, id int64) error {
	return errors.New("can not implement this method as MongoDB has its own unique ID")
}
