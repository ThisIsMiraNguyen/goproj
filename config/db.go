package config

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func Connect() {
    dsn := "root:@Meimei52Hz_Arya@@tcp(localhost:3306)/todo_app"
    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
    if err = DB.Ping(); err != nil {
        log.Fatal("Failed to ping database: ", err)
    }
    fmt.Println("Connected to MySQL database!")
}