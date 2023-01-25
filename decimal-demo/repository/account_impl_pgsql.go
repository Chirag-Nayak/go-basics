package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/Chirag-Nayak/go-basics/decimal-demo/model"
)

type AccountImplPgsql struct {
	logger *log.Logger
	db     *sql.DB
}

func NewAccountImplPgsql(l *log.Logger, dbClient *sql.DB) *AccountImplPgsql {
	return &AccountImplPgsql{
		logger: l,
		db:     dbClient,
	}
}

func (w *AccountImplPgsql) GetAll(ctx context.Context) (model.Accounts, error) {
	query := `
	SELECT id, account_name, currency_name, balance
		FROM account_info
	`
	rows, err := w.db.QueryContext(ctx, query)
	if err != nil {
		w.logger.Printf("Error while retrieving data from DB: %s", err)
		return nil, err
	}
	defer rows.Close()

	var acs []*model.Account
	for rows.Next() {
		var acc model.Account
		if err := rows.Scan(&acc.ID, &acc.AccountName, &acc.CurrencyName, &acc.Balance); err != nil {
			w.logger.Printf("Error while scaning the result returned by query: %s", err)
			return nil, err
		}
		acs = append(acs, &acc)
	}
	return acs, nil
}
