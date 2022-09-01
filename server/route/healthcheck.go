package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "working...")
}
