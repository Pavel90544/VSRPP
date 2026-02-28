package app

import (
	"database/sql"
	"go-sqlite-example/domain"

	_ "github.com/mattn/go-sqlite3"
)

// ProductDB implements DatabaseInterface
type ProductDB struct {
	db *sql.DB
}

// NewProductDB creates a new ProductDB instance
func NewProductDB(dbPath string) (*ProductDB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Create table if not exists
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS products(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        model TEXT,
        company TEXT,
        price INTEGER
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return &ProductDB{db: db}, nil
}

// GetProduct retrieves a product by ID
func (p *ProductDB) GetProduct(id int) (*domain.Product, error) {
	var product domain.Product
	err := p.db.QueryRow("SELECT id, model, company, price FROM products WHERE id = ?", id).
		Scan(&product.ID, &product.Model, &product.Company, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetAllProducts retrieves all products
func (p *ProductDB) GetAllProducts() ([]*domain.Product, error) {
	rows, err := p.db.Query("SELECT id, model, company, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.ID, &product.Model, &product.Company, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

// CreateProduct adds a new product
func (p *ProductDB) CreateProduct(product *domain.Product) error {
	result, err := p.db.Exec(
		"INSERT INTO products (model, company, price) VALUES (?, ?, ?)",
		product.Model, product.Company, product.Price,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.ID = int(id)
	return nil
}

// UpdateProduct modifies an existing product
func (p *ProductDB) UpdateProduct(product *domain.Product) error {
	_, err := p.db.Exec(
		"UPDATE products SET model = ?, company = ?, price = ? WHERE id = ?",
		product.Model, product.Company, product.Price, product.ID,
	)
	return err
}

// DeleteProduct removes a product
func (p *ProductDB) DeleteProduct(id int) error {
	_, err := p.db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}

// Close closes the database connection
func (p *ProductDB) Close() error {
	return p.db.Close()
}
