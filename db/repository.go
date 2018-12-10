package db

import (
	"context"
	"stock/schema"
)

type Repository interface {
	//OK
	CreateStock(ctx context.Context, stock *schema.Stock) error

	//OK
	ModifyStock(ctx context.Context, id string, quantity int) error

	//OK
	GetStock(ctx context.Context, id string) (*schema.Stock, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func CreateStock(ctx context.Context, stock *schema.Stock) error {
	return impl.CreateStock(ctx, stock)
}

func ModifyStock(ctx context.Context, id string, quantity int) error {
	return impl.ModifyStock(ctx, id, quantity)
}

func GetStock(ctx context.Context, id string) (*schema.Stock, error) {
	return impl.GetStock(ctx, id)
}
