package main

import (
	"example/web-service/internal/routes"
	"github.com/gin-gonic/gin"
	"example/web-service/internal/db"
	"example/web-service/handler"
)
func main(){
	dbConn := db.NewPostgres()
	defer dbConn.Close()

	router := gin.Default()

	albumHandler := &handler.AlbumHandler{
		DB: dbConn,
	}

	routes.RegisterRoutes(router, albumHandler)

	router.Run(":8080")
}

