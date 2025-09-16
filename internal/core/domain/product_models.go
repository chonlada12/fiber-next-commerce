package domain

import "github.com/google/uuid"

// omitempty ใช้เพื่อไม่ให้แสดงฟิลด์ที่เป็นค่า null ใน JSON

// Product สำหรับเก็บข้อมูลสินค้า
type Product struct {
	BaseModel
	Name        string         `gorm:"type:varchar(100)" json:"name" validate:"required"`
	Description string         `gorm:"type:text" json:"description"`
	Price       float64        `gorm:"type:decimal(10,2)" json:"price" validate:"required,min=0"`
	Stock       int            `gorm:"type:int" json:"stock" validate:"min=0"`
	Image       string         `gorm:"type:varchar(255)" json:"image"`
	Images      []ProductImage `gorm:"foreignKey:ProductID" json:"images,omitempty"`
	CategoryID  uuid.UUID      `json:"category_id" validate:"required"`
	Category    Category       `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	OrderItems  []OrderItem    `gorm:"foreignKey:ProductID" json:"order_items,omitempty"`
	CartItems   []CartItem     `gorm:"foreignKey:ProductID" json:"cart_items,omitempty"`
}

// ProductImage สำหรับเก็บรูปภาพของสินค้า
type ProductImage struct {
	BaseModel
	ProductID uuid.UUID `json:"product_id"`
	ImageURL  string    `gorm:"type:varchar(255)" json:"image_url" validate:"required"`
}
