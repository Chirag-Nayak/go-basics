package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/Chirag-Nayak/go-basics/decimal-demo/model"
	"github.com/lib/pq"
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

func (a *AccountImplPgsql) GetAll(ctx context.Context) (model.Accounts, error) {
	query := `
	SELECT id, account_name, currency_name, balance
		FROM account_info
	`

	// Execute the select query
	rows, err := a.db.QueryContext(ctx, query)
	if err != nil {
		a.logger.Printf("Error while retrieving data from DB: %#+v\n", err)
		return nil, err
	}
	defer rows.Close()

	// Read the result returned from DB
	var acs []*model.Account
	for rows.Next() {
		var acc model.Account
		if err := rows.Scan(&acc.ID, &acc.AccountName, &acc.CurrencyName, &acc.Balance); err != nil {
			a.logger.Printf("Error while scaning the result returned by Select query: %#+v\n", err)
			repoErr := GetAccountRepoErr(err)
			return nil, repoErr
		}
		acs = append(acs, &acc)
	}
	return acs, nil
}

func (a *AccountImplPgsql) GetById(ctx context.Context, id int64) (*model.Account, error) {
	query := `SELECT  id, account_name, currency_name, balance FROM account_info
		WHERE id = $1`

	var acc model.Account
	res := a.db.QueryRowContext(ctx, query, id)
	err := res.Scan(&acc.ID, &acc.AccountName, &acc.CurrencyName, &acc.Balance)
	if err != nil {
		log.Printf("Error while scanning the result returned by Select query: %#+v\n", err)
		repoErr := GetAccountRepoErr(err)
		return nil, repoErr
	}
	return &acc, nil
}

func (a *AccountImplPgsql) AddAccount(ctx context.Context, acc model.Account) (*model.Account, error) {
	query := `
	INSERT INTO account_info (account_name, currency_name, balance) VALUES
		($1, $2, $3) RETURNING id;
	`
	var id int64

	// Perform the insert operation & scan the returned ID from db
	err := a.db.QueryRowContext(ctx, query, acc.AccountName, acc.CurrencyName, acc.Balance).Scan(&id)
	if err != nil {
		a.logger.Printf("Error while executing the insert query: %#+v\n", err)
		repoErr := GetAccountRepoErr(err)
		return nil, repoErr
	}

	// Set the ID returned from db & reutrn
	acc.ID = id
	return &acc, nil
}

func (a *AccountImplPgsql) UpdateAccount(ctx context.Context, id int64, acc model.Account) (*model.Account, error) {
	// Validate input parameters
	if id <= 0 {
		return nil, ErrInvalidAccountID
	}

	query := `
	UPDATE account_info SET
		account_name = $1, currency_name = $2, balance = $3
		WHERE id = $4
	`

	// Execute the update query on DB
	res, err := a.db.ExecContext(ctx, query, acc.AccountName, acc.CurrencyName, acc.Balance, id)
	if err != nil {
		log.Printf("Error while executing update query on DB: %#+v\n", err)
		repoErr := GetAccountRepoErr(err)
		return nil, repoErr
	}

	// Check the returned result
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error while checking rows affected after update query: %#+v\n ", err)
		repoErr := GetAccountRepoErr(err)
		return nil, repoErr
	}
	if rows == 0 {
		return nil, ErrRecordNotExist
	}

	return &acc, nil
}

func (a *AccountImplPgsql) DeleteAccount(ctx context.Context, id int64) error {
	query := `Delete from account_info WHERE id = $1`

	res, err := a.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Error while executing Delete query on DB: %#+v\n", err)
		repoErr := GetAccountRepoErr(err)
		return repoErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error while getting rows affected after Delete query: %#+v\n", err)
		repoErr := GetAccountRepoErr(err)
		return repoErr
	}

	if rowsAffected == 0 {
		// Record does not exist, can be treated as an error as well depending on the design
		return nil
	}
	return err
}

func GetAccountRepoErr(err error) error {
	if err == nil {
		return nil
	}

	// Check for all possible errors that we can narrow down
	if errors.Is(err, sql.ErrNoRows) {
		return ErrRecordNotExist
	}

	var pgErr *pq.Error
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" { // "23505": "unique_violation"
			// Refer here for more details on error codes:
			// https://www.postgresql.org/docs/9.3/errcodes-appendix.html

			// You can further check by constraint name in order to give detailed errors to service
			if pgErr.Constraint == DB_CONST_CONSTRAINT_UNQ_ACC_NAME {
				return ErrDuplicateAccountName
			}
			// Return default unique constraint violation error
			return ErrDuplicateRecord
		}
	}

	// Return deault error as it is
	return err
}
