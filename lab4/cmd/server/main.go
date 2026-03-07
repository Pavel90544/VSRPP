package main

import (
    "log"
    "net/http"
    "github.com/Pavel90544/VSRPP/lab4/internal/repository"
    "github.com/Pavel90544/VSRPP/lab4/internal/router"
)

func main() {
    // Initialize database
    repo, err := repository.NewProductRepository("products.db")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer repo.Close()

    // Setup routes
    r := router.SetupRoutes(repo)

    // Start server
    port := ":8080"
    log.Printf("Server starting on port %s", port)
    log.Printf("Available endpoints:")
    log.Printf("  POST   /api/products - Create product")
    log.Printf("  GET    /api/products - Get all products")
    log.Printf("  GET    /api/products/{id} - Get product by ID")
    log.Printf("  PUT    /api/products/{id} - Update product")
    log.Printf("  DELETE /api/products/{id} - Delete product")

    if err := http.ListenAndServe(port, r); err != nil {
        log.Fatal("Server failed to start:", err)
    }
}
