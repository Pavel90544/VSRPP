package app

import (
    "database/sql"
    "testing"
    "go-sqlite-example/domain"
    "go-sqlite-example/mocks"
    "go.uber.org/mock/gomock"
    _ "github.com/mattn/go-sqlite3"
)

func TestNewProductDB(t *testing.T) {
    // Test with valid database path
    db, err := NewProductDB(":memory:")
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if db == nil {
        t.Error("Expected db to be non-nil")
    }
    db.Close()
}

func TestCreateAndGetProduct(t *testing.T) {
    db, err := NewProductDB(":memory:")
    if err != nil {
        t.Fatalf("Failed to create db: %v", err)
    }
    defer db.Close()
    
    // Create product
    product := &domain.Product{
        Model:   "iPhone 13",
        Company: "Apple",
        Price:   99000,
    }
    
    err = db.CreateProduct(product)
    if err != nil {
        t.Errorf("Failed to create product: %v", err)
    }
    
    if product.ID == 0 {
        t.Error("Expected product ID to be set")
    }
    
    // Get product
    retrieved, err := db.GetProduct(product.ID)
    if err != nil {
        t.Errorf("Failed to get product: %v", err)
    }
    
    if retrieved.Model != product.Model {
        t.Errorf("Expected model %s, got %s", product.Model, retrieved.Model)
    }
}

func TestGetAllProducts(t *testing.T) {
    db, err := NewProductDB(":memory:")
    if err != nil {
        t.Fatalf("Failed to create db: %v", err)
    }
    defer db.Close()
    
    // Add some products
    products := []*domain.Product{
        {Model: "iPhone 13", Company: "Apple", Price: 99000},
        {Model: "Galaxy S21", Company: "Samsung", Price: 75000},
    }
    
    for _, p := range products {
        err := db.CreateProduct(p)
        if err != nil {
            t.Errorf("Failed to create product: %v", err)
        }
    }
    
    // Get all products
    all, err := db.GetAllProducts()
    if err != nil {
        t.Errorf("Failed to get all products: %v", err)
    }
    
    if len(all) != 2 {
        t.Errorf("Expected 2 products, got %d", len(all))
    }
}

func TestUpdateProduct(t *testing.T) {
    db, err := NewProductDB(":memory:")
    if err != nil {
        t.Fatalf("Failed to create db: %v", err)
    }
    defer db.Close()
    
    // Create product
    product := &domain.Product{
        Model:   "iPhone 13",
        Company: "Apple",
        Price:   99000,
    }
    
    err = db.CreateProduct(product)
    if err != nil {
        t.Errorf("Failed to create product: %v", err)
    }
    
    // Update product
    product.Price = 89000
    err = db.UpdateProduct(product)
    if err != nil {
        t.Errorf("Failed to update product: %v", err)
    }
    
    // Verify update
    updated, err := db.GetProduct(product.ID)
    if err != nil {
        t.Errorf("Failed to get updated product: %v", err)
    }
    
    if updated.Price != 89000 {
        t.Errorf("Expected price 89000, got %d", updated.Price)
    }
}

func TestDeleteProduct(t *testing.T) {
    db, err := NewProductDB(":memory:")
    if err != nil {
        t.Fatalf("Failed to create db: %v", err)
    }
    defer db.Close()
    
    // Create product
    product := &domain.Product{
        Model:   "iPhone 13",
        Company: "Apple",
        Price:   99000,
    }
    
    err = db.CreateProduct(product)
    if err != nil {
        t.Errorf("Failed to create product: %v", err)
    }
    
    // Delete product
    err = db.DeleteProduct(product.ID)
    if err != nil {
        t.Errorf("Failed to delete product: %v", err)
    }
    
    // Verify deletion
    _, err = db.GetProduct(product.ID)
    if err != sql.ErrNoRows {
        t.Errorf("Expected sql.ErrNoRows, got %v", err)
    }
}

// Test with mocks
func TestWithMock(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    
    mockDB := mocks.NewMockDatabaseInterface(ctrl)
    
    // Setup expectations
    expectedProduct := &domain.Product{ID: 1, Model: "Test", Company: "TestCo", Price: 1000}
    mockDB.EXPECT().GetProduct(1).Return(expectedProduct, nil)
    
    // Test code that uses the mock
    product, err := mockDB.GetProduct(1)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    
    if product.Model != "Test" {
        t.Errorf("Expected 'Test', got %s", product.Model)
    }
}
