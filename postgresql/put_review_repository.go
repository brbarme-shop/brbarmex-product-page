package postgresql

import (
	"context"
	"database/sql"

	"github.com/brbarme-shop/brbarmex-product-page/product"
	_ "github.com/lib/pq"
)

type repository struct {
	db *sql.DB
}

func (r *repository) PutItem(ctx context.Context, product *product.ProductPage) error {
	return nil
}

func NewProductRepository(db *sql.DB) product.ProductRepository {
	return &repository{db: db}
}
