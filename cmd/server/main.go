package main

import (
	"example/web-service/internal/routes"
	"github.com/gin-gonic/gin"
	"example/web-service/internal/db"
)
func main(){
	dbConn := db.NewPostgres()
	defer dbConn.Close()

	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":8080")
}

