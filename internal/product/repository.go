package product

import (
	"context"
	"database/sql"

	"inventory-api/internal/db"
)

type Repository interface {
	Create(ctx context.Context, p *db.Product) error
	Update(ctx context.Context, p *db.Product) error
	GetByID(ctx context.Context, id int) (*db.Product, error)
	Delete(ctx context.Context, id int) error
}

type repository struct {
	conn *sql.DB
}

func NewRepository(conn *sql.DB) Repository {
	return &repository{conn: conn}
}

func (r *repository) Create(ctx context.Context, p *db.Product) error {
	query := "INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id"
	return r.conn.QueryRowContext(ctx, query, p.Name, p.Description, p.Price).Scan(&p.ID)
}

func (r *repository) Update(ctx context.Context, p *db.Product) error {
	query := "UPDATE products SET name=$1, description=$2, price=$3 WHERE id=$4"
	_, err := r.conn.ExecContext(ctx, query, p.Name, p.Description, p.Price, p.ID)
	return err
}

func (r *repository) GetByID(ctx context.Context, id int) (*db.Product, error) {
	query := "SELECT id, name, description, price FROM products WHERE id=$1"
	row := r.conn.QueryRowContext(ctx, query, id)

	var p db.Product
	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM products WHERE id=$1"
	_, err := r.conn.ExecContext(ctx, query, id)
	return err
}
