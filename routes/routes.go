package routes

import (
	"example/web-service/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine){
	router.GET("/albums", handler.GetAlbums)
	router.POST("/albums", handler.PostAlbums)
	router.GET("/albums/:id", handler.GetAlbumByID)
}
