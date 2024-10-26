package routes

import (
    "ecommerce-backend/handlers"
    "github.com/gin-gonic/gin"
)

// SetupRoutes initializes all routes
func SetupRoutes() *gin.Engine {
    router := gin.Default()

    // User routes
    router.POST("/register", handlers.Register)
    router.POST("/login", handlers.Login)

    // Product routes
    productRoutes := router.Group("/products")
    {
        productRoutes.POST("/", handlers.CreateProduct)
        productRoutes.GET("/", handlers.GetProducts)
        productRoutes.PUT("/:id", handlers.UpdateProduct)
        productRoutes.DELETE("/:id", handlers.DeleteProduct)
    }

    return router
}
