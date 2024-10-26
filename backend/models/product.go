package models

import "time"

// Product represents the product entity.
type Product struct {
    ID          int       `json:"id"`
    Name        string    `json:"name" binding:"required"`
    Description string    `json:"description"`
    Price       float64   `json:"price" binding:"required"`
    Stock       int       `json:"stock" binding:"required"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
