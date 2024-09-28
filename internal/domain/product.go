package domain

import "time"

type Product struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	ShopID    uint64    `json:"shop_id"`
	Name      string    `gorm:"size:255" json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
