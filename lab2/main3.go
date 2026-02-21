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

    // Обновление данных
    result, err := db.Exec("UPDATE products SET price = ? WHERE id = ?", 69000, 1)
    if err != nil {
        log.Fatal(err)
    }

    rowsAffected, _ := result.RowsAffected()
    fmt.Printf("Количество обновленных строк: %d\n", rowsAffected)
}
