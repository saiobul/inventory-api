package product_test

import (
	"context"
	"testing"
	"time"

	"inventory-api/internal/db"
	"inventory-api/internal/product"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) Create(ctx context.Context, p *db.Product) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *mockRepo) Update(ctx context.Context, p *db.Product) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *mockRepo) GetByID(ctx context.Context, id int) (*db.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*db.Product), args.Error(1)
}

func (m *mockRepo) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type mockCache struct {
	mock.Mock
}

func (m *mockCache) Get(key string) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

func (m *mockCache) Set(key string, value string, ttl time.Duration) error {
	args := m.Called(key, value, ttl)
	return args.Error(0)
}

func (m *mockCache) Delete(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func TestAddProduct(t *testing.T) {
	repo := new(mockRepo)
	cache := new(mockCache)
	svc := product.NewService(repo, cache)

	p := &db.Product{ID: 1, Name: "New Product"}
	repo.On("Create", mock.Anything, p).Return(nil)

	err := svc.AddProduct(context.Background(), p)
	assert.NoError(t, err)
	repo.AssertCalled(t, "Create", mock.Anything, p)
}

func TestUpdateProduct(t *testing.T) {
	repo := new(mockRepo)
	cache := new(mockCache)
	svc := product.NewService(repo, cache)

	p := &db.Product{ID: 2, Name: "Updated Product"}
	repo.On("Update", mock.Anything, p).Return(nil)
	cache.On("Delete", "product:2").Return(nil)

	err := svc.UpdateProduct(context.Background(), p)
	assert.NoError(t, err)
	repo.AssertCalled(t, "Update", mock.Anything, p)
	cache.AssertCalled(t, "Delete", "product:2")
}

func TestDeleteProduct(t *testing.T) {
	repo := new(mockRepo)
	cache := new(mockCache)
	svc := product.NewService(repo, cache)

	repo.On("Delete", mock.Anything, 3).Return(nil)
	cache.On("Delete", "product:3").Return(nil)

	err := svc.DeleteProduct(context.Background(), 3)
	assert.NoError(t, err)
	repo.AssertCalled(t, "Delete", mock.Anything, 3)
	cache.AssertCalled(t, "Delete", "product:3")
}
