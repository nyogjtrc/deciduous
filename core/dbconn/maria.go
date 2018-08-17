package dbconn

import "database/sql"

var (
	rConnection *sql.DB
	wConnection *sql.DB
)

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

func DBRead() *sql.DB {
	return rConnection
}

func DBWrite() *sql.DB {
	return wConnection
}
