package domain

import "github.com/google/uuid"

// omitempty ใช้เพื่อไม่ให้แสดงฟิลด์ที่เป็นค่า null ใน JSON

// Cart สำหรับเก็บข้อมูลตะกร้าสินค้า
type Cart struct {
	BaseModel
	UserID     uuid.UUID  `json:"user_id"`
	User       User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CartItems  []CartItem `gorm:"foreignKey:CartID" json:"cart_items,omitempty"`
	TotalPrice float64    `gorm:"type:decimal(10,2)" json:"total_price"`
}

// CartItem สำหรับเก็บรายการสินค้าในตะกร้า
type CartItem struct {
	BaseModel
	CartID    uuid.UUID `json:"cart_id"`
	Cart      Cart      `gorm:"foreignKey:CartID" json:"cart,omitempty"`
	ProductID uuid.UUID `json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int       `gorm:"type:int" json:"quantity" validate:"required,min=1"`
	Price     float64   `gorm:"type:decimal(10,2)" json:"price"`
}
