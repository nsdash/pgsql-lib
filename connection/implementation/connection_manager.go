package implementation

import (
	"database/sql"
	"fmt"
	dns2 "github.com/nsdash/pgsql-lib/connection/dns"
	"log"

	_ "github.com/lib/pq"
)

type ConnectionManager struct {
	openedConnection *sql.DB
}

var connectionManager *ConnectionManager

func GetConnectionManagerSingleton() ConnectionManager {
	if connectionManager == nil {
		connectionManager = &ConnectionManager{openedConnection: nil}
	}

	return *connectionManager
}

func (cm *ConnectionManager) GetConnection() *sql.DB {

	if cm.openedConnection != nil {
		return cm.openedConnection
	}

	cm.openedConnection = openConnection()

	return cm.openedConnection
}

func (cm *ConnectionManager) CloseConnection() {
	if cm.openedConnection == nil {
		return
	}

	err := cm.openedConnection.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func openConnection() *sql.DB {
	dnsData := dns2.NewManager().GetData()

	dns := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dnsData.Host,
		dnsData.Port,
		dnsData.User,
		dnsData.Password,
		dnsData.Database,
	)

	connection, err := sql.Open(dnsData.DriverName, dns)

	if err != nil {
		log.Fatal(err)
	}

	return connection
}
