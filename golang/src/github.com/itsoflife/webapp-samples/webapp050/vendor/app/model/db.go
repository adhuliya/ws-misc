package model

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// The global datastore interface objects.
// One datastore for each source used in code.
var dataStore1 DataStore

// contains connection information
type DbConf struct {
	user     string
	password string
	dbname   string
	host     string
	port     string
	ssl      string
}

type DataStore interface {
	Open() error  // open the current datastore
	Close() error // safely close the connection
	Name()        // name of current datastore
	Create(query string, data ...interface{}) (uint64, error)
	Read(query string, data ...interface{}) (*sql.Rows, error)
	// for both Update and Delete operations
	Modify(query string, data ...interface{}) (bool, error)
}

// UTILITY FUNCTIONS
// Returns the last insert id returned by the insert function.
func ScanInsertId(res *sql.Rows) uint64 {
	res.Next()
	var id uint64
	res.Scan(&id)
	return id
}

// Returns the boolean returned by the update/delete function.
func ScanBool(res *sql.Rows) bool {
	res.Next()
	var b bool
	res.Scan(&b)
	return b
}

// Close the resultset to avoid memory leaks.
// `err` is the sql error returned on query execution,
// if `err` in not nil, then there is no result.
// `query`, `data` are used for logging purpose.
func CloseResult(res *sql.Rows, err error, query string, data interface{}) {
	if err == nil && res != nil {
		if e := res.Close(); e != nil {
			logs.Logger.Error(query, data)
		}
	}
}
