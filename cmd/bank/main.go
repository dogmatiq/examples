package main

import (
	"context"
	"os"
	"time"

	"github.com/dogmatiq/dogma"
	"github.com/dogmatiq/example"
	"github.com/dogmatiq/example/messages/commands"
	"github.com/dogmatiq/testkit/engine"
)

func main() {
	app := &example.App{}

	en, err := engine.New(app)
	if err != nil {
		panic(err)
	}

	messages := []dogma.Message{
		commands.OpenAccountForNewCustomer{
			CustomerID:   "cust1",
			CustomerName: "Anna Smith",
			AccountID:    "acct1",
			AccountName:  "Anna Smith",
		},
		commands.OpenAccountForNewCustomer{
			CustomerID:   "cust2",
			CustomerName: "Bob Jones",
			AccountID:    "acct2",
			AccountName:  "Bob Jones",
		},
		commands.Deposit{
			TransactionID: "txn1",
			AccountID:     "acct1",
			Amount:        10000,
		},
		commands.Withdraw{
			TransactionID: "txn2",
			AccountID:     "acct1",
			Amount:        500,
			ScheduledDate: time.Now().Format("2006-01-02"),
		},
		commands.Transfer{
			TransactionID: "txn3",
			FromAccountID: "acct1",
			ToAccountID:   "acct2",
			Amount:        2500,
			ScheduledDate: time.Now().Format("2006-01-02"),
		},
		commands.Transfer{
			TransactionID: "txn4",
			FromAccountID: "acct1",
			ToAccountID:   "acct2",
			Amount:        500,
			ScheduledDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
		},
	}

	for _, m := range messages {
		err := en.Dispatch(
			context.Background(),
			m,
			// engine.WithObserver(
			// 	fact.ObserverFunc(func(f fact.Fact) {
			// 		dapper.Print(f)
			// 		fmt.Print("\n\n")
			// 	}),
			// ),
			engine.EnableProjections(true),
		)
		if err != nil {
			panic(err)
		}
	}

	if err := app.GenerateAccountCSV(os.Stdout); err != nil {
		panic(err)
	}
}
