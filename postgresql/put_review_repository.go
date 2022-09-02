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

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}

	_sql := `INSERT INTO products
	(product_item_id, product_tile, product_description, product_information, product_datetime)
	VALUES(uuid_generate_v1(), $1, $2, $3, CURRENT_TIMESTAMP) RETURNING product_id;
	`

	var productId int64
	err = tx.QueryRowContext(ctx, _sql, product.Title, product.Description, product.Information).Scan(&productId)
	if err != nil {
		return nil
	}

	if productId <= 0 {
		return nil
	}

	_sql = `INSERT INTO products_sku (product_id, product_sku_code, product_sku_option, product_sku_option_value, product_sku_details, product_sku_qtd, product_sku_price) VALUES `

	vals := []interface{}{}
	for i := range product.Skus {
		_sql += "($1, $2, $3, $4, $5, $6, $7),"
		vals = append(vals, productId, product.Skus[i].Code, product.Skus[i].OptionType, product.Skus[i].OptionValue, product.Skus[i].Detail, product.Skus[i].Quantity, product.Skus[i].Price)
	}

	_sql = _sql[0 : len(_sql)-1]

	// stmt, err := db.PrepareContext(ctx, _sql)
	// if err != nil {
	// 	err = tx.Rollback()
	// 	return err
	// }

	sqlRows, err := tx.ExecContext(ctx, _sql, vals...)
	if err != nil {
		err = tx.Rollback()
		return err
	}

	rowsAffected, err := sqlRows.RowsAffected()
	if err != nil {
		err = tx.Rollback()
		return err
	}

	if rowsAffected <= 0 {
		err = tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func (r *repository) Exist(ctx context.Context, itemId string) (bool, error) {
	return false, nil
}

func NewProductRepository(db *sql.DB) product.ProductRepository {
	return &repository{db: db}
}
