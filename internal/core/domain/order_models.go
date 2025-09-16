package domain

import "github.com/google/uuid"

// omitempty ใช้เพื่อไม่ให้แสดงฟิลด์ที่เป็นค่า null ใน JSON

// Order สำหรับเก็บข้อมูลการสั่งซื้อ
type Order struct {
	BaseModel
	UserID          uuid.UUID     `json:"user_id"`
	User            User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	OrderItems      []OrderItem   `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
	TotalPrice      float64       `gorm:"type:decimal(10,2)" json:"total_price"`
	Status          string        `gorm:"type:varchar(50);default:'pending'" json:"status"`
	PaymentMethod   string        `gorm:"type:varchar(50)" json:"payment_method"`
	PaymentStatus   string        `gorm:"type:varchar(50);default:'pending'" json:"payment_status"`
	ShippingMethod  string        `gorm:"type:varchar(50)" json:"shipping_method"`
	ShippingStatus  string        `gorm:"type:varchar(50);default:'pending'" json:"shipping_status"`
	ShippingAddress string        `gorm:"type:text" json:"shipping_address"`
	TrackingNumber  string        `gorm:"type:varchar(100)" json:"tracking_number"`
	Notes           string        `gorm:"type:text" json:"notes"`
	Transactions    []Transaction `gorm:"foreignKey:OrderID" json:"transactions,omitempty"`
}

// OrderItem สำหรับเก็บรายการสินค้าในคำสั่งซื้อ
type OrderItem struct {
	BaseModel
	OrderID   uuid.UUID `json:"order_id"`
	Order     Order     `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	ProductID uuid.UUID `json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int       `gorm:"type:int" json:"quantity" validate:"required,min=1"`
	Price     float64   `gorm:"type:decimal(10,2)" json:"price"`
}
