package pgsql

import (
	"database/sql"
	_ "github.com/lib/pq"
	internal_sql "github.com/nsdash/pgsql-lib/connection/implementation"
)

type SqlManager struct {
	connection *sql.DB
}

func NewSqlManager() SqlManager {
	connection := internal_sql.GetConnectionManagerSingleton().GetConnection()

	return SqlManager{connection: connection}
}

func (s SqlManager) CountGt(query string, countToCompare int) bool {
	var count int

	err := s.connection.QueryRow(query).Scan(&count)

	if err != nil {
		panic(err)
	}

	return count > countToCompare
}

func (s SqlManager) Exec(query string) {
	_, err := s.connection.Exec(query)

	if err != nil {
		panic(err)
	}
}

func (s SqlManager) Query(query string) *sql.Rows {
	rows, err := s.connection.Query(query)

	if err != nil {
		panic(err)
	}

	return rows
}

func (s SqlManager) QueryRow(query string) *sql.Row {
	row := s.connection.QueryRow(query)

	return row
}

func (s SqlManager) ExecTransaction(query string, transaction *sql.Tx) sql.Result {
	result, err := transaction.Exec(query)

	if err != nil {
		panic(err)
	}

	return result
}

func (s SqlManager) Transaction(callable func(transaction *sql.Tx)) {
	transaction, err := s.connection.Begin()

	if err != nil {
		panic("Transaction failed: " + err.Error())
	}

	defer func() {
		if err := recover(); err != nil {
			err := transaction.Rollback()

			if err != nil {
				panic("Failed rollback transaction: " + err.Error())
			}

			panic("Transaction failed")
		}
	}()

	callable(transaction)

	err = transaction.Commit()

	if err != nil {
		panic("Transaction failed: " + err.Error())
	}
}
