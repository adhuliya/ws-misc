package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "time"
)

const (
    DB_USER     = "hop"
    DB_PASSWORD = "anshuisneo"
    DB_NAME     = "hop"
)

func main() {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
    DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    fmt.Println("# Inserting values")

    var lastInsertId int
    err = db.QueryRow("INSERT INTO public.user(username, password, role, createdOn) VALUES($1,$2,$3,$4) returning id;", "astaxie", "anshu", "user2", "2012-12-09T10:10:20").Scan(&lastInsertId)
    checkErr(err)
    fmt.Println("last inserted id =", lastInsertId)

    fmt.Println("# Updating")
    stmt, err := db.Prepare("update public.user set username=$1 where id=$2")
    checkErr(err)

    res, err := stmt.Exec("astaxieupdate", lastInsertId)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect, "rows changed")

    fmt.Println("# Querying")
    rows, err := db.Query("SELECT id, username, password, createdOn FROM public.user")
    checkErr(err)

    for rows.Next() {
        var id int
        var username string
        var password string
        var created time.Time
        err = rows.Scan(&id, &username, &password, &created)
        checkErr(err)
        fmt.Println("id | username | password | createdOn ")
        fmt.Printf("%3v | %8v | %6v | %6v\n", id, username, department, created)
    }

    fmt.Println("# Deleting")
    stmt, err = db.Prepare("delete from public.user where id=$1")
    checkErr(err)

    res, err = stmt.Exec(lastInsertId)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect, "rows changed")
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
