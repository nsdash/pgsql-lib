package contracts

import "database/sql"

type ConnectionManager interface {
	GetConnection() *sql.DB
	CloseConnection()
}
