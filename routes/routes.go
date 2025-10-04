package routes

import (
	"ecommerce-backend/controllers"
	"ecommerce-backend/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")

	// 🔑 Auth
	api.POST("/users/register", controllers.Register(db))
	api.POST("/users/login", controllers.Login(db))

	// 🌍 Public endpoints
	api.GET("/products", controllers.GetProducts(db))
	api.GET("/categories", controllers.GetCategories(db))

	// 👤 Protected with JWT (customer)
	api.Use(middlewares.JWTAuthMiddleware())
	{
		api.GET("/users/me", func(c *gin.Context) {
			userID := c.MustGet("user_id").(uint)
			role := c.MustGet("role").(string)

			c.JSON(200, gin.H{
				"user_id": userID,
				"role":    role,
			})
		})

		// Orders (Customer only)
		api.POST("/orders", controllers.CreateOrder(db))
		api.GET("/orders", controllers.GetOrders(db))
		api.GET("/orders/:id", controllers.GetOrderByID(db))
	}

	// 👑 Admin routes (JWT + role check)
	admin := api.Group("/admin")
	admin.Use(middlewares.JWTAuthMiddleware()) // pakai JWT
	{
		admin.GET("/orders", controllers.GetAllOrders(db))
		admin.GET("/orders/:id", controllers.GetOrderByIDAdmin(db))
		admin.PUT("/orders/:id/status", controllers.UpdateOrderStatus(db))

		admin.POST("/products", controllers.CreateProduct(db))
		admin.PUT("/products/:id", controllers.UpdateProduct(db))
		admin.DELETE("/products/:id", controllers.DeleteProduct(db))

		admin.POST("/categories", controllers.CreateCategory(db))
		admin.PUT("/categories/:id", controllers.UpdateCategory(db))
		admin.DELETE("/categories/:id", controllers.DeleteCategory(db))
	}
}


