package product_test

import (
	"context"
	"inventory-api/internal/db"
	"inventory-api/internal/product"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	dbConn, mock, _ := sqlmock.New()
	defer dbConn.Close()

	repo := product.NewRepository(dbConn)

	rows := sqlmock.NewRows([]string{"id", "name", "description", "price"}).AddRow(1, "Test", "Desc", 9.99)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, description, price FROM products WHERE id=$1")).
		WithArgs(1).WillReturnRows(rows)

	result, err := repo.GetByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "Test", result.Name)
}

func TestUpdate(t *testing.T) {
	dbConn, mock, _ := sqlmock.New()
	defer dbConn.Close()

	repo := product.NewRepository(dbConn)

	mock.ExpectExec(regexp.QuoteMeta("UPDATE products SET name=$1, description=$2, price=$3 WHERE id=$4")).
		WithArgs("Updated", "New Desc", 20.0, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	p := &db.Product{ID: 1, Name: "Updated", Description: "New Desc", Price: 20.0}
	err := repo.Update(context.Background(), p)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	dbConn, mock, _ := sqlmock.New()
	defer dbConn.Close()

	repo := product.NewRepository(dbConn)

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM products WHERE id=$1")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(context.Background(), 1)
	assert.NoError(t, err)
}
