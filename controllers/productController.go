package controllers

import (
	"ecommerce-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GET /products
func GetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		if err := db.Preload("Category").Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

// POST /products (Admin only - BasicAuth)
func CreateProduct(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var product models.Product
        if err := c.ShouldBindJSON(&product); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := db.Create(&product).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // 🔑 reload dengan preload category
        db.Preload("Category").First(&product, product.ID)

        c.JSON(http.StatusCreated, product)
    }
}

// PUT /products/:id (Admin only - BasicAuth)
func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")

        var product models.Product
        if err := db.First(&product, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
            return
        }

        if err := c.ShouldBindJSON(&product); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        db.Save(&product)

        // 🔑 reload dengan preload category
        db.Preload("Category").First(&product, product.ID)

        c.JSON(http.StatusOK, product)
    }
}

// DELETE /products/:id (Admin only - BasicAuth)
func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := db.Delete(&models.Product{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
	}
}
