package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func ConnectDB() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   "root",
        Passwd: "my-secret-pw",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "todo",
    }
    // Get a database handle.
    var err error
    DB, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := DB.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}
