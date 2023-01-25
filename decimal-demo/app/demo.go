package app

import (
	"fmt"
	"log"

	"github.com/Chirag-Nayak/go-basics/decimal-demo/repository"
	"github.com/Chirag-Nayak/go-basics/decimal-demo/service"
	"github.com/shopspring/decimal"
)

var (
	aService *service.Account
	logger   *log.Logger
)

func RunDecimalDemo(l *log.Logger, aRepo repository.Account) {

	// Initizlize the service & logger
	aService = service.NewAccountService(l, aRepo)
	logger = l

	DemoReadAccountsFromDB()
}

func DemoReadAccountsFromDB() {
	acs, err := aService.GetAccounts()
	if err != nil {
		logger.Printf("Error from Service: %s", err)
		return
	}

	logger.Println("-------------- Account information received from DB --------------")
	for _, wallet := range acs {
		logger.Printf("%+v\n", wallet)
	}
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
