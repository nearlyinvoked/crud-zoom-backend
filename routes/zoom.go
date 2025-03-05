package routes

import (
	"github.com/gin-gonic/gin"

	"crud-zoom/controllers"
	"crud-zoom/services"
)

func (r *route) initZoom(app *gin.RouterGroup) {
	zoomSvc := services.NewZoomService(r.cfg, r.logger, r.repo)
	zoomCtrl := controllers.NewZoomController(r.cfg, r.logger, zoomSvc)

	zoom := app.Group("/zoom")
	zoom.GET("/", zoomCtrl.List)
	zoom.POST("/create", zoomCtrl.Create)
	zoom.PATCH("/update/:id", zoomCtrl.Update)
	zoom.DELETE("/delete/:id", zoomCtrl.Delete)
}
