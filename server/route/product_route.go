package route

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/brbarme-shop/brbarmex-product-page/config"
	"github.com/brbarme-shop/brbarmex-product-page/postgresql"
	"github.com/brbarme-shop/brbarmex-product-page/product"
	"github.com/gin-gonic/gin"
)

var (
	cfg        = config.NewConfiguration()
	db         = postgresql.NewSqlDB(cfg)
	repository = postgresql.NewProductRepository(db)
)

func postProduct(c *gin.Context) {

	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	defer c.Request.Body.Close()

	var productPageInput *product.ProductPageInput
	err = json.Unmarshal(b, &productPageInput)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	err = product.PutNewProductPage(c.Request.Context(), productPageInput, repository)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, nil)
}
