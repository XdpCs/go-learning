package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func goodsHandler(c *gin.Context) {
	c.String(http.StatusOK, "goods:xdp")
}

func checkoutHandler(c *gin.Context) {
	c.String(http.StatusOK, "checkout:fyy")
}
