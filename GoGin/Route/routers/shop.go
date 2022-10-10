package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadShop(e *gin.Engine) {
	e.GET("/hello", helloHandler)
	e.GET("/goods", goodsHandler)
	e.GET("/checkout", checkoutHandler)
}

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "hello,fyy")
}

func goodsHandler(c *gin.Context) {
	c.String(http.StatusOK, "goods:xdp")
}

func checkoutHandler(c *gin.Context) {
	c.String(http.StatusOK, "checkout:fyy")
}
