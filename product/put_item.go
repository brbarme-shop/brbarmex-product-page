package product

import (
	"context"
	"errors"
)

var (
	ErrProductPageAlreadyExist = errors.New("hey")
)

type ProductRepository interface {
	PutItem(ctx context.Context, product *ProductPage) error
	Exist(ctx context.Context, itemId string) (bool, error)
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
	Id          string     `json:"item_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Information string     `json:"information"`
	Sku         []SkuInput `json:"skus"`
}

type SkuInput struct {
	Code     string  `json:"code"`
	Detail   string  `json:"detail"`
	Quantity float64 `json:"qtd"`
	Price    float64 `json:"price"`
}

func PutProductPage(ctx context.Context, productPageInput *ProductPageInput, db ProductRepository) error {

	exist, err := db.Exist(ctx, productPageInput.Id)

	if err != nil {
		return err
	}

	if exist {
		return ErrProductPageAlreadyExist
	}

	return db.PutItem(ctx, buildModel(productPageInput))
}

func buildModel(dto *ProductPageInput) *ProductPage {

	var sku = make([]Sku, 0, len(dto.Sku))

	for i := range dto.Sku {
		sku = append(sku, Sku{
			Code:     dto.Sku[i].Code,
			Detail:   dto.Sku[i].Detail,
			Quantity: dto.Sku[i].Quantity,
			Price:    dto.Sku[i].Price,
		})
	}

	return &ProductPage{
		Id:          dto.Id,
		Title:       dto.Title,
		Description: dto.Description,
		Information: dto.Information,
		Skus:        sku,
	}
}
