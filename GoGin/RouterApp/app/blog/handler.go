package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func postHandler(c *gin.Context) {
	c.String(http.StatusOK, "post:fyy")
}

func commentHandler(c *gin.Context) {
	c.String(http.StatusOK, "comment:xdp")
}
