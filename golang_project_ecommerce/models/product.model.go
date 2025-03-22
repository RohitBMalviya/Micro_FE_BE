// Package models ...
package models

// Product model struct
type Product struct {
	BaseModel
	ProductName        string `gorm:"type:varchar;not null" json:"product_name"`
	ProductPrice       int    `gorm:"type:int(6);not null" json:"product_price"`
	ProductDiscount    int    `gorm:"type:int(3);not null" json:"product_discount"`
	ProductDescription string `gorm:"type:varchar;not null" json:"product_description"`
	ProductImage       string `gorm:"type:varchar;not null" json:"product_image"`
	ProductWebsite     string `gorm:"type:varchar;not null" json:"product_website"`
	ProductStockCount  int    `gorm:"type:int(4);not null" json:"product_stock_count"`
}
