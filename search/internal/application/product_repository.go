package application

import (
	"context"

	"eda-in-golang/search/internal/models"
)

type ProductRepository interface {
	Find(ctx context.Context, productID string) (*models.Product, error)
}

type ProductCacheRepository interface {
	Add(ctx context.Context, productID, storeID, name string, price float64) error
	Rebrand(ctx context.Context, productID, name string) error
	Remove(ctx context.Context, productID string) error
	ProductRepository
}
