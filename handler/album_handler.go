package handler

import (
	"example/web-service/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context){
	c.JSON(http.StatusOK ,models.Albums)
}

func PostAlbums(c *gin.Context){
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.Albums = append(models.Albums ,newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context){
	id := c.Param("id")

	for _ ,album := range models.Albums{
		if album.ID==id{
			c.JSON(http.StatusOK, album)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}