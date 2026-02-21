package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Открытие подключения к базе данных
    db, err := sql.Open("sqlite3", "store.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Удаление данных
    result, err := db.Exec("DELETE FROM products WHERE id = ?", 1)
    if err != nil {
        log.Fatal(err)
    }

    rowsAffected, _ := result.RowsAffected()
    fmt.Printf("Количество удаленных строк: %d\n", rowsAffected)
}
