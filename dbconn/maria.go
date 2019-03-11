package dbconn

import "database/sql"

var (
	rConnection *sql.DB
	wConnection *sql.DB
)

// OpenRead DB connection
func OpenRead(DSN string) {
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

// OpenWrite DB connection
func OpenWrite(DSN string) {
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
