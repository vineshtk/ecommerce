package handlers

import (
    "ecommerce-backend/config"
    "ecommerce-backend/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

// Create a new product
func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    query := `INSERT INTO products (name, description, price, stock) 
              VALUES ($1, $2, $3, $4) RETURNING id`
    err := config.DB.QueryRow(query, product.Name, product.Description, product.Price, product.Stock).Scan(&product.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }

    c.JSON(http.StatusCreated, product)
}

// Get all products
func GetProducts(c *gin.Context) {
    rows, err := config.DB.Query("SELECT * FROM products")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
        return
    }
    defer rows.Close()

    var products []models.Product
    for rows.Next() {
        var product models.Product
        if err := rows.Scan(&product.ID, &product.Name, &product.Description, 
                            &product.Price, &product.Stock, 
                            &product.CreatedAt, &product.UpdatedAt); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan product"})
            return
        }
        products = append(products, product)
    }

    c.JSON(http.StatusOK, products)
}

// Update a product by ID
func UpdateProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    query := `UPDATE products SET name=$1, description=$2, price=$3, stock=$4, updated_at=NOW() 
              WHERE id=$5`
    _, err := config.DB.Exec(query, product.Name, product.Description, product.Price, product.Stock, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// Delete a product by ID
func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    query := "DELETE FROM products WHERE id=$1"
    _, err := config.DB.Exec(query, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}


func test(){

    
}
