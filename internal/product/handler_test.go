package product_test

import (
	"bytes"
	"context"
	"encoding/json"
	"inventory-api/internal/db"
	"inventory-api/internal/product"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

func (m *mockService) AddProduct(ctx context.Context, p *db.Product) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *mockService) UpdateProduct(ctx context.Context, p *db.Product) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *mockService) GetProduct(ctx context.Context, id int) (*db.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*db.Product), args.Error(1)
}

func (m *mockService) DeleteProduct(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateProduct(t *testing.T) {
	svc := new(mockService)
	handler := product.NewHandler(svc)

	prod := db.Product{ID: 1, Name: "Test"}
	svc.On("AddProduct", mock.Anything, &prod).Return(nil)

	body, _ := json.Marshal(prod)
	req := httptest.NewRequest("POST", "/products", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler.CreateProduct(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetProductTimeout(t *testing.T) {
	svc := new(mockService)
	handler := product.NewHandler(svc)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	time.Sleep(2 * time.Nanosecond)

	req := httptest.NewRequest("GET", "/products/1", nil).WithContext(ctx)
	rr := httptest.NewRecorder()

	handler.GetProduct(rr, req)

	assert.Equal(t, http.StatusRequestTimeout, rr.Code)
}

func TestUpdateProductInvalidID(t *testing.T) {
	svc := new(mockService)
	handler := product.NewHandler(svc)

	req := httptest.NewRequest("PUT", "/products/abc", nil)
	rr := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{"id": "abc"})

	handler.UpdateProduct(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
