package repository

import (
	"database/sql"

	"github.com/gustavoeyros/message-broker-golang/internal/entities"
)

type ProductRepositoryMySql struct {
	DB *sql.DB
}

func NewProductRepositoryMySql(db *sql.DB) *ProductRepositoryMySql {
	return &ProductRepositoryMySql{DB: db}
}

func (p *ProductRepositoryMySql) Create(product *entities.Product) error {
	_, err := p.DB.Exec("INSERT INTO products (id, name, price) values (?, ?, ?)", product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepositoryMySql) FindAll() ([]*entities.Product, error) {
	rows, err := p.DB.Query("SELECT id, name, price FROM Products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entities.Product

	for rows.Next() {
		var product entities.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)

	}
	return products, nil
}
