package controllers

import (
	"ecommerce-backend/models"
	"net/http"
	"strconv"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// POST /orders (JWT - User only)
func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("user_id").(uint)
		fmt.Println("DEBUG userID:", userID)

		var req struct {
			Items []struct {
				ProductID uint `json:"product_id"`
				Quantity  int  `json:"quantity"`
			} `json:"items"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var total float64
		var orderItems []models.OrderItem

		// hitung total + validasi produk
		for _, item := range req.Items {
			var product models.Product
			if err := db.First(&product, item.ProductID).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
				return
			}

			if product.Stock < item.Quantity {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Stock not enough"})
				return
			}

			subtotal := product.Price * float64(item.Quantity)
			total += subtotal

			orderItems = append(orderItems, models.OrderItem{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Subtotal:  subtotal,
			})

			// kurangi stock
			product.Stock -= item.Quantity
			db.Save(&product)
		}

		// simpan order dulu
		order := models.Order{
			UserID: userID,
			Total:  total,
			Status: "pending",
		}

		if err := db.Create(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// simpan order items dengan OrderID
		for i := range orderItems {
			orderItems[i].OrderID = order.ID
		}
		if err := db.Create(&orderItems).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 🔑 Preload biar response lengkap
		db.Preload("User").
			Preload("Items.Product.Category").
			First(&order, order.ID)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Order created successfully",
			"order":   order,
		})
	}
}



// GET /orders (JWT - lihat semua order user login)
func GetOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("user_id").(uint)

		var orders []models.Order
		if err := db.
			Preload("Items.Product.Category"). // ⬅️ load item + product + category aja
			Where("user_id = ?", userID).
			Find(&orders).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, orders)
	}
}

// GET /orders/:id (JWT - detail order user login)
func GetOrderByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("user_id").(uint)
		id, _ := strconv.Atoi(c.Param("id"))

		var order models.Order
		if err := db.
			Preload("Items.Product.Category").
			Where("id = ? AND user_id = ?", id, userID).
			First(&order).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}

		c.JSON(http.StatusOK, order)
	}
}

// PUT /orders/:id/status (Admin - JWT)
func UpdateOrderStatus(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var order models.Order
		// Cari order + preload items (supaya nanti bisa return lengkap)
		if err := db.Preload("Items.Product.Category").First(&order, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}

		// Bind request JSON
		var req struct {
			Status string `json:"status"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update status
		order.Status = req.Status
		if err := db.Save(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Reload order dengan preload lengkap
		db.Preload("Items.Product.Category").
			First(&order, order.ID)

		// Return JSON lengkap
		c.JSON(http.StatusOK, gin.H{
			"message": "Order status updated successfully",
			"order":   order,
		})
	}
}


// GET /admin/orders (JWT - Admin only)
func GetAllOrders(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.MustGet("role").(string)
        if role != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Only admin can view all orders"})
            return
        }

        var orders []models.Order
        if err := db.Preload("User").
            Preload("Items.Product.Category").
            Find(&orders).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, orders)
    }
}

// GET /admin/orders/:id (JWT - Admin only)
func GetOrderByIDAdmin(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.MustGet("role").(string)
        if role != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Only admin can view order detail"})
            return
        }

        id, _ := strconv.Atoi(c.Param("id"))

        var order models.Order
        if err := db.Preload("User").
            Preload("Items.Product.Category").
            First(&order, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
            return
        }

        c.JSON(http.StatusOK, order)
    }
}

