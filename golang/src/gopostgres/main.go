package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "time"
)

type User struct {
    id          int
    username    string
    password    string
    nickname    string
    fullname    string
    mobile      string
    email       string
    role        string
    status      string
    createdOn   time.Time
    modifiedOn  time.Time
}

const (
    DB_HOST     = "127.0.0.1"
    DB_PORT     = 5432
    DB_USER     = "hop"
    DB_PASSWORD = "anshuisneo"
    DB_NAME     = "hop"
    DB_SSLMODE  = false;
)

// Connect to Postgres with the given arguments.
func ConnPostgres(host string, port int, usr string, password string, dbname string, sslmode bool) (*sql.DB, error) {
    ssl := "disable"
    if sslmode {
        ssl = "enable"
    }
    dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, usr, password, dbname, ssl)

    db, err := sql.Open("postgres", dbinfo)

    return db, err
}

// Close connection to Postgres.
// Needed only when server is brought down. Hence rarely needed.
// use: defer ClosePostgres(db)
func ClosePostgres(db *sql.DB) error {
    return db.Close()
}

// Returns the last indert id returned by the insert function.
func scanInsertId(res *sql.Rows) int {
    res.Next()
    var id int
    res.Scan(&id)
    return id
}

// Returns the boolean returned by the update/delete function.
func scanBool(res *sql.Rows) bool {
    res.Next()
    var b bool
    res.Scan(&b)
    return b
}

func main() {
    db, err := ConnPostgres(DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSLMODE)
    checkErr(err)
    defer ClosePostgres(db)

    fmt.Println("# Inserting values")

    //var lastInsertId int

    res, err := db.Query("select insertUser(username := $1, mobile := $2, email := $3)", "hero1", "+917645263720", "hero1@mail.com")
    checkErr(err)
    id := scanInsertId(res);
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

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}


