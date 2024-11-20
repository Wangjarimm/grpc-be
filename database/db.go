// database/db.go
package database

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
    dsn := "root:@tcp(127.0.0.1:3306)/grpc_example" // Sesuaikan dengan konfigurasi MySQL Anda
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
        return nil, err
    }
    return db, nil
}
