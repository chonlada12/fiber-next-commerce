package domain

import "github.com/google/uuid"

// omitempty ใช้เพื่อไม่ให้แสดงฟิลด์ที่เป็นค่า null ใน JSON

// Transaction สำหรับเก็บข้อมูลธุรกรรมการชำระเงิน
type Transaction struct {
	BaseModel
	OrderID       uuid.UUID `json:"order_id"`
	Order         Order     `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	Amount        float64   `gorm:"type:decimal(10,2)" json:"amount"`
	PaymentMethod string    `gorm:"type:varchar(50)" json:"payment_method"`
	Status        string    `gorm:"type:varchar(50);default:'pending'" json:"status"`
	TransactionID string    `gorm:"type:varchar(100)" json:"transaction_id"`
	PaymentData   string    `gorm:"type:text" json:"payment_data"`
}
