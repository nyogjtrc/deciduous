package conn

import "database/sql"

var (
	dbRead  *sql.DB
	dbWrite *sql.DB
)

// SetDBRead will replace db read
func SetDBRead(db *sql.DB) {
	dbRead = db
}

// SetDBWrite will replace db write
func SetDBWrite(db *sql.DB) {
	dbRead = db
}

// DBRead return read connection
func DBRead() *sql.DB {
	return dbRead
}

// DBWrite return write connection
func DBWrite() *sql.DB {
	return dbWrite
}

// DBReadOpen DB connection
func DBReadOpen(DSN string) (*sql.DB, error) {
	var err error

	db, err := sql.Open("mysql", DSN)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// DBWriteOpen DB connection
func DBWriteOpen(DSN string) (*sql.DB, error) {
	var err error

	db, err := sql.Open("mysql", DSN)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
