package httpnumbersfunctions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OrderNumbers(c *gin.Context) {
	c.String(http.StatusOK, "Vamos a ordenar los numeros")
}
