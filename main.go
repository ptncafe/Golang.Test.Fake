package main

import (
	"github.com/gin-gonic/gin"
	"ptncafe.test/controllers"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/api")
	{
		client.GET("/store/:id", controllers.GetStoreById)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":12345")
}
