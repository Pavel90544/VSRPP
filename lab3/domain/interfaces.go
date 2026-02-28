package domain

// Product represents a product in the database
type Product struct {
    ID      int
    Model   string
    Company string
    Price   int
}

// DatabaseInterface defines methods for database operations
type DatabaseInterface interface {
    GetProduct(id int) (*Product, error)
    GetAllProducts() ([]*Product, error)
    CreateProduct(product *Product) error
    UpdateProduct(product *Product) error
    DeleteProduct(id int) error
    Close() error
}
