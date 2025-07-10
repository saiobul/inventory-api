package product

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"inventory-api/internal/db"
	"inventory-api/internal/utils"

	"go.uber.org/zap"
)

type Service interface {
	AddProduct(ctx context.Context, p *db.Product) error
	UpdateProduct(ctx context.Context, p *db.Product) error
	GetProduct(ctx context.Context, id int) (*db.Product, error)
	DeleteProduct(ctx context.Context, id int) error
}

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string, ttl time.Duration) error
	Delete(key string) error
}

type service struct {
	repo  Repository
	cache Cache
}

func NewService(repo Repository, cache Cache) Service {
	return &service{repo: repo, cache: cache}
}

func (s *service) AddProduct(ctx context.Context, p *db.Product) error {
	logger := zap.L().With(
		zap.String("request_id", utils.GetRequestID(ctx)),
		zap.String("user_id", utils.GetUserID(ctx)),
	)
	logger.Info("Adding product")

	select {
	case <-ctx.Done():
		logger.Warn("AddProduct cancelled or timed out")
		return ctx.Err()
	default:
	}

	return s.repo.Create(ctx, p)
}

func (s *service) UpdateProduct(ctx context.Context, p *db.Product) error {
	logger := zap.L().With(
		zap.String("request_id", utils.GetRequestID(ctx)),
		zap.String("user_id", utils.GetUserID(ctx)),
		zap.Int("product_id", p.ID),
	)
	logger.Info("Updating product")

	select {
	case <-ctx.Done():
		logger.Warn("UpdateProduct cancelled or timed out")
		return ctx.Err()
	default:
	}

	err := s.repo.Update(ctx, p)
	if err == nil {
		s.cache.Delete(fmt.Sprintf("product:%d", p.ID))
	}
	return err
}

func (s *service) GetProduct(ctx context.Context, id int) (*db.Product, error) {
	logger := zap.L().With(
		zap.String("request_id", utils.GetRequestID(ctx)),
		zap.String("user_id", utils.GetUserID(ctx)),
		zap.Int("product_id", id),
	)
	logger.Info("Fetching product")

	select {
	case <-ctx.Done():
		logger.Warn("GetProduct cancelled or timed out")
		return nil, ctx.Err()
	default:
	}

	key := fmt.Sprintf("product:%d", id)

	cached, err := s.cache.Get(key)
	if err == nil {
		var p db.Product
		if json.Unmarshal([]byte(cached), &p) == nil {
			logger.Info("Product found in cache")
			return &p, nil
		}
	}

	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if data, err := json.Marshal(product); err == nil {
		s.cache.Set(key, string(data), 10*time.Minute)
	}

	return product, nil
}

func (s *service) DeleteProduct(ctx context.Context, id int) error {
	logger := zap.L().With(
		zap.String("request_id", utils.GetRequestID(ctx)),
		zap.String("user_id", utils.GetUserID(ctx)),
		zap.Int("product_id", id),
	)
	logger.Info("Deleting product")

	select {
	case <-ctx.Done():
		logger.Warn("DeleteProduct cancelled or timed out")
		return ctx.Err()
	default:
	}

	err := s.repo.Delete(ctx, id)
	if err == nil {
		s.cache.Delete(fmt.Sprintf("product:%d", id))
	}
	return err
}
