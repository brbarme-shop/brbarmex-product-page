package product

import "context"

type ProductRepository interface {
	PutItem(ctx context.Context, product *ProductPage) error
}

type ProductPage struct {
	Id          string
	Title       string
	Description string
	Information string
	Skus        []Sku
}

type Sku struct {
	Code     string
	Detail   string
	Quantity float64
	Price    float64
}

type ProductPageInput struct {
}

func PutNewProductPage(ctx context.Context, productPageInput *ProductPageInput, db ProductRepository) error {
	return nil
}
