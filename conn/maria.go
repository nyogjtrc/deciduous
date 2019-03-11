package conn

import "database/sql"

var (
	rConnection *sql.DB
	wConnection *sql.DB
)

// DBOpenRead DB connection
func DBOpenRead(DSN string) {
	var err error

	rConnection, err = sql.Open("mysql", DSN)
	if err != nil {
		panic(err.Error())
	}
	err = rConnection.Ping()
	if err != nil {
		panic(err.Error())
	}
}

// DBOpenWrite DB connection
func DBOpenWrite(DSN string) {
	var err error

	wConnection, err = sql.Open("mysql", DSN)
	if err != nil {
		panic(err.Error())
	}
	err = wConnection.Ping()
	if err != nil {
		panic(err.Error())
	}
}

// DBRead return read connection
func DBRead() *sql.DB {
	return rConnection
}

// DBWrite return write connection
func DBWrite() *sql.DB {
	return wConnection
}
