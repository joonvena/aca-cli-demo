package controllers

import (
	"net/http"
	"products/models"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	var input models.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{Name: input.Name, Description: input.Description, Image: input.Image, Price: input.Price}
	models.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

func ListProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	c.JSON(http.StatusOK, products)
}
