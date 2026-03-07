package repository

import (
    "database/sql"
    "github.com/Pavel90544/VSRPP/lab4/internal/models"
    _ "github.com/mattn/go-sqlite3"
)

type ProductRepository struct {
    db *sql.DB
}

func NewProductRepository(dbPath string) (*ProductRepository, error) {
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

    return &ProductRepository{db: db}, nil
}

func (r *ProductRepository) Close() error {
    return r.db.Close()
}

func (r *ProductRepository) Create(product *models.Product) error {
    result, err := r.db.Exec(
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

func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
    var product models.Product
    err := r.db.QueryRow(
        "SELECT id, model, company, price FROM products WHERE id = ?", id,
    ).Scan(&product.ID, &product.Model, &product.Company, &product.Price)

    if err != nil {
        return nil, err
    }

    return &product, nil
}

func (r *ProductRepository) GetAll() ([]*models.Product, error) {
    rows, err := r.db.Query("SELECT id, model, company, price FROM products")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []*models.Product
    for rows.Next() {
        var product models.Product
        err := rows.Scan(&product.ID, &product.Model, &product.Company, &product.Price)
        if err != nil {
            return nil, err
        }
        products = append(products, &product)
    }

    return products, nil
}

func (r *ProductRepository) Update(product *models.Product) error {
    _, err := r.db.Exec(
        "UPDATE products SET model = ?, company = ?, price = ? WHERE id = ?",
        product.Model, product.Company, product.Price, product.ID,
    )
    return err
}

func (r *ProductRepository) Delete(id int) error {
    _, err := r.db.Exec("DELETE FROM products WHERE id = ?", id)
    return err
}
