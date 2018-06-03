/*
This program is a demonstration of how to use postgres stored procedures.
*/

package model

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	//	"time"

	"app/logs"
)

type DbPostgres struct {
	conf *DbConf
	conn *sql.DB
}

func (db *DbPostgres) GetConf() *DbConf {
	return db.conf
}

func (db *DbPostgres) SetConf(conf *DbConf) {
	db.conf = conf
}

func (db *DbPostgres) GetConn() *sql.DB {
	return db.conn
}

func (db *DbPostgres) SetConn(conn *sql.DB) {
	db.conn = conn
}

func (db *DbPostgres) Open(conf *DbConf) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		conf.Host,
		conf.Port,
		conf.UserName,
		conf.Password,
		conf.DbName,
		conf.SslMode)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		logs.Logger.Error("POSTGRES CONN FAILED. ERROR MSG: ", err, " CONFIGURATION: ", conf)
		return err
	}

	if err = conn.Ping(); err != nil {
		logs.Logger.Error("POSTGRES CONN PING FAILED. ERROR MSG: ", err, " CONFIGURATION: ", conf)
		return err
	}

	// Connection Successful!
	conf.Password = "" // for security, as conf is printed in logs
	logs.Logger.Info("POSTGRES CONN SUCCESSFUL. CONFIGURATION: ", conf)
	db.conf = conf
	db.conn = conn
	return nil
}

// Close connection to Postgres.
// Needed only when server is brought down.
// use: defer ClosePostgres(db)
func (db *DbPostgres) Close() error {
	if db.conn != nil {
		logs.Logger.Info("CLOSING POSTGRES CONN. SETTINGS: ", db.conn)
		return db.conn.Close()
	} else {
		logs.Logger.Error("POSTGRES CLOSING nil CONN. CONF: ", db.conf)
		return errors.New("CLOSING a nil connection.")
	}
}

func (db *DbPostgres) Name() string {
	if db.conf != nil {
		return fmt.Sprintf("postgres:%s", db.conf.DbName)
	} else {
		logs.Logger.Error("POSTGRES NO DBNAME.")
		return "postgres:??"
	}
}

func (db *DbPostgres) Create(query string, data ...interface{}) (uint64, error) {
	res, err := db.conn.Query(query, data)
	defer CloseResult(res, err, query, data)
	if err != nil {
		logs.Logger.Error(query, data)
		return 0, nil
	}

	// query should return the insert id
	id := ScanInsertId(res)
	return id, nil
}

func (db *DbPostgres) Read(query string, data ...interface{}) (*sql.Rows, error) {
	res, err := db.conn.Query(query, data...)
	//defer CloseResult(res, err, query, data) //caller's responsibility
	if err != nil {
		logs.Logger.Error(err, query, data)
		return nil, err
	}

	return res, nil
}

// Handles Update and Delete operations.
func (db *DbPostgres) Modify(query string, data ...interface{}) (bool, error) {
	res, err := db.conn.Query(query, data)
	defer CloseResult(res, err, query, data)
	if err != nil {
		logs.Logger.Error(query, data)
		return false, err
	}

	b := ScanBool(res)
	return b, nil
}
