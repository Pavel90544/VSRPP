package main

import (
    "fmt"
    "log"
    "go-sqlite-example/app"
    "go-sqlite-example/domain"
)

func main() {
    fmt.Println("Лабораторная работа №2: Тестирование")
    fmt.Println("=====================================")
    
    // Create database instance
    db, err := app.NewProductDB("test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Create a product
    product := &domain.Product{
        Model:   "iPhone 14",
        Company: "Apple",
        Price:   120000,
    }
    
    err = db.CreateProduct(product)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created product with ID: %d\n", product.ID)
    
    // Get all products
    products, err := db.GetAllProducts()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("\nAll products:")
    for _, p := range products {
        fmt.Printf("ID: %d, Model: %s, Company: %s, Price: %d\n",
            p.ID, p.Model, p.Company, p.Price)
    }
    
    fmt.Println("\nТесты можно запустить командой: go test ./app -v")
    fmt.Println("Моки сгенерированы в ./mocks/domain_mocks.go")
}
