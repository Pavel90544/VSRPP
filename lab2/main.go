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

    // Проверка подключения
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Подключение к базе данных успешно установлено")

    // Добавление данных
    result, err := db.Exec("INSERT INTO products (model, company, price) VALUES (?, ?, ?)",
        "iPhone X", "Apple", 72000)
    if err != nil {
        log.Fatal(err)
    }

    lastId, _ := result.LastInsertId()
    rowsAffected, _ := result.RowsAffected()

    fmt.Printf("ID последнего добавленного объекта: %d\n", lastId)
    fmt.Printf("Количество добавленных строк: %d\n", rowsAffected)
}
