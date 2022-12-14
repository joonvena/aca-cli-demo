package main

import (
	"net/http"
	"os"
	"products/controllers"
	"products/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	ALLOW_ORIGINS = os.Getenv("ALLOW_ORIGINS")
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{ALLOW_ORIGINS},
		AllowMethods: []string{"*"},
	}))

	models.ConnectDatabase()

	r.GET("/products", controllers.ListProducts)
	r.POST("/products", controllers.AddProduct)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})
	r.Run()
}
