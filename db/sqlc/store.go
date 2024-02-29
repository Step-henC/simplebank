package db

import (
	"context"
	"database/sql"
	"fmt"
)

// store provides all functions to execute queries and transactions
// queries struct does not operate on multiple tables at once
// for every bank transaction we have to operate on all three tables
// this way is "composition" and is preferred way in golang to extend store functionality vs "inheritance"
type Store struct {
	*Queries //all queries are available to us now
	db       *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// executes a function within db transaction
// do not want outer packages to call it
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {

	tx, err := store.db.BeginTx(ctx, nil) //nil usually is sql.TxOption to set isolation level

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction err: %v, rollback err: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}

// all the fields needed to create transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// result of transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// creates transfer record, adds account entries and update accounts' balances within single db tx
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {

	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(context.Background(), CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount, //money is moving from this account
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount, //money is moving TO this account
		})

		if err != nil {
			return err
		}

		//get account -> update its balance :need proper locking mechanism

		if arg.FromAccountID < arg.ToAccountID { //perform tx in same order ea time to avoid db deadlock

			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)

		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}
		return nil
	})

	return result, err
}

// refactor function to avoid long code of swapping order of accounts trying to avoid db deadlock
func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64, //first account id
	amount1 int64, // money transferred out
	accountID2 int64, // second account id
	amount2 int64, //money transferred in

) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return //acts the same as adding accounts cool syntax feature of go
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})

	return
}
