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
