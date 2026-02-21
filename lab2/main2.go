package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

type Product struct {
    ID      int
    Model   string
    Company string
    Price   int
}

func main() {
    // Открытие подключения к базе данных
    db, err := sql.Open("sqlite3", "store.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Получение всех данных
    rows, err := db.Query("SELECT * FROM products")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    products := []Product{}

    for rows.Next() {
        p := Product{}
        err := rows.Scan(&p.ID, &p.Model, &p.Company, &p.Price)
        if err != nil {
            log.Println(err)
            continue
        }
        products = append(products, p)
    }

    // Проверка на ошибки после итерации
    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }

    // Вывод результатов
    fmt.Println("Список продуктов:")
    for _, p := range products {
        fmt.Printf("ID: %d, Model: %s, Company: %s, Price: %d\n", 
            p.ID, p.Model, p.Company, p.Price)
    }

    // Получение одной строки
    fmt.Println("\nПолучение продукта с ID=1:")
    row := db.QueryRow("SELECT * FROM products WHERE id = ?", 1)
    prod := Product{}
    err = row.Scan(&prod.ID, &prod.Model, &prod.Company, &prod.Price)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("Продукт не найден")
        } else {
            log.Fatal(err)
        }
    } else {
        fmt.Printf("ID: %d, Model: %s, Company: %s, Price: %d\n", 
            prod.ID, prod.Model, prod.Company, prod.Price)
    }
}
