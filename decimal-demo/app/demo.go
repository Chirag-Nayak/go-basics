package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Chirag-Nayak/go-basics/decimal-demo/model"
	"github.com/Chirag-Nayak/go-basics/decimal-demo/repository"
	"github.com/Chirag-Nayak/go-basics/decimal-demo/service"
	"github.com/shopspring/decimal"
)

var (
	aService      *service.Account
	logger        *log.Logger
	ctxDefTimeOut context.Context
)

func RunDecimalDemo(l *log.Logger, aRepo repository.Account) {

	// Define context
	var cancel context.CancelFunc
	ctxDefTimeOut, cancel = context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// Initizlize the service & logger
	aService = service.NewAccountService(l, aRepo)
	logger = l

	logger.Println(">>>>>>>>>>>>>> Reading the current account information from DB.")
	DemoReadAccountsFromDB()

	logger.Println(">>>>>>>>>>>>>> Inserting new account informatino into the DB.")
	DemoInsertAccountToDB()

	logger.Println(">>>>>>>>>>>>>> Updating account informatino in the DB for ID 4.")
	DemoUpdateAccountInDB()

	logger.Println(">>>>>>>>>>>>>> Retrieving account informatino from the DB for ID 4.")
	DemoGetAccountByIDFromDB()

	logger.Println(">>>>>>>>>>>>>> Deleting account informatino from the DB for ID 4.")
	DemoDeleteAccountByIDFromDB()

	logger.Println(">>>>>>>>>>>>>> Reading the current account information from DB.")
	DemoReadAccountsFromDB()
}

func DemoReadAccountsFromDB() {
	acs, err := aService.GetAccounts(ctxDefTimeOut)
	if err != nil {
		logger.Printf("Error in DemoReadAccountsFromDB from Service: %s", err)
		return
	}

	logger.Println("-------------- Account information received from DB --------------")
	for _, wallet := range acs {
		logger.Printf("%+v\n", wallet)
	}
}

func DemoInsertAccountToDB() {
	accToAdd := model.Account{
		AccountName:  "My Account 4",
		CurrencyName: "JPY",
		Balance:      decimal.NewFromInt(200000),
	}

	addedAcc, err := aService.AddAccount(ctxDefTimeOut, accToAdd)
	if err != nil {
		logger.Printf("Error in DemoInsertAccountToDB from service: %s", err)
		return
	}
	logger.Printf("Inserted account details: %+v\n", addedAcc)
}

func DemoUpdateAccountInDB() {

	uBal, err := decimal.NewFromString("12500.12500")
	if err != nil {
		logger.Printf("Can not convert into Decimal, Skipping DemoUpdateAccountInDB")
		return
	}

	accToUpdate := model.Account{
		AccountName:  "My Account 4",
		CurrencyName: "USD",
		Balance:      uBal,
	}

	updatedAcc, err := aService.UpdateAccount(ctxDefTimeOut, 4, accToUpdate)
	if err != nil {
		logger.Printf("Error in DemoUpdateAccountInDB from service: %s", err)
		return
	}
	logger.Printf("Inserted account details: %+v\n", updatedAcc)
}

func DemoGetAccountByIDFromDB() {
	acc, err := aService.GetAccountById(ctxDefTimeOut, 4)
	if err != nil {
		logger.Printf("Error in DemoGetAccountByIDFromDB from Service: %s", err)
		return
	}
	logger.Println("-------------- Account information received from DB for Account ID 4 are --------------")
	logger.Printf("%+v\n", acc)
}

func DemoDeleteAccountByIDFromDB() {
	err := aService.DeleteAccount(ctxDefTimeOut, 4)
	if err != nil {
		logger.Printf("Error in DemoGetAccountByIDFromDB from Service: %s", err)
		return
	}
	log.Println("Account with Account ID 4 is deleted successfully.")
}

func DemoBasicOperations() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("fee:", fee)                                // Subtotal: 0.035
	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875
}
