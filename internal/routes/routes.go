package routes

import (
	"example/web-service/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, albumHandler *handler.AlbumHandler) {
	r.GET("/albums", albumHandler.GetAlbums)
	r.POST("/albums", albumHandler.CreateAlbum)
}
