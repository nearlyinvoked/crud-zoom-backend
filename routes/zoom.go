package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *route) initZoom(bs *gin.RouterGroup) {

	user := bs.Group("/zoom")
	user.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
}
