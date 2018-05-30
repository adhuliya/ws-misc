/*
This program is a demonstration of how to use postgres stored procedures.
*/

package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"

	"app/logs"
)

type DbPostgres struct {
	dbConf DbConf
	conn   *sql.DB
}

func (db *DbPostgres) Open(conf DbConf) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.host,
		conf.port,
		conf.user,
		conf.password,
		conf.dbname,
		conf.ssl)

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
	conf.password = "" // for security, as conf is printed in logs
	logs.Logger.Error("POSTGRES CONN SUCCESSFUL. CONFIGURATION: ", conf)
	db.dbConf = conf
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
		return fmt.Sprintf("postgres:%s", db.conf.dbname)
	} else {
		logs.Logger.Error("POSTGRES NO DBNAME.")
		return "postgres:??"
	}
}

func (db *DbPostgres) Create(query string, data ...interface{}) (int64, error) {
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
	res, err := db.conn.Query(query, data)
	//defer CloseResult(res, err, query, data) //caller's responsibility
	if err != nil {
		logs.Logger.Error(query, data)
		return nil, err
	}
	defer CloseResult(res)

	return res, nil
}

// Handles Update and Delete operations.
func (db *DbPostgres) Modify(query string, data ...interface{}) (bool, error) {
	res, err := db.conn.Query(query, data)
	defer CloseResult(res, err, query, data)
	if err != nil {
		logs.Logger.Error(query, data)
		return nil, err
	}
	defer CloseResult(res)

	b := ScanBool(res)
	return b, nil
}

func main() {
	db, err := ConnPostgres(DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSLMODE)
	checkErr(err)
	defer ClosePostgres(db)

	fmt.Println("# Inserting values")

	//var lastInsertId int

	res, err := db.Query("select insertUser(username := $1, mobile := $2, email := $3)", "hero1", "+917645263720", "hero1@mail.com")
	checkErr(err)
	id := scanInsertId(res)
	fmt.Println("Insert ID:", id)

	res, err = db.Query("select updateUser(id := $1, username := $2)", id, "hero2")
	checkErr(err)
	b := scanBool(res)
	fmt.Println("Updated:", b)

	time.Sleep(1 * time.Second)

	res, err = db.Query("select * from selectUser()")
	checkErr(err)
	var user = User{}
	count := 0
	for res.Next() {
		count++
		err = res.Scan(
			&user.id,
			&user.username,
			&user.password,
			&user.nickname,
			&user.fullname,
			&user.mobile,
			&user.email,
			&user.role,
			&user.status,
			&user.createdOn,
			&user.modifiedOn,
		)
		checkErr(err)
		fmt.Println(user)
	}
	fmt.Println("Row Count:", count)

	res, err = db.Query("select deleteUser(id := $1)", id)
	checkErr(err)
	b = scanBool(res)
	fmt.Println("Deleted:", b)

	fmt.Println("Success.")
}
