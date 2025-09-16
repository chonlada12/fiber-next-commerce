package domain

import (
	"time"

	"github.com/google/uuid"
)

// omitempty ใช้เพื่อไม่ให้แสดงฟิลด์ที่เป็นค่า null ใน JSON

// User สำหรับเก็บข้อมูลผู้ใช้งาน
type User struct {
	BaseModel
	Email            string    `gorm:"type:varchar(100);unique_index" json:"email" validate:"required,email"`
	Password         string    `gorm:"type:varchar(100)" json:"-" validate:"required,min=6"`
	FirstName        string    `gorm:"type:varchar(100)" json:"first_name" validate:"required"`
	LastName         string    `gorm:"type:varchar(100)" json:"last_name" validate:"required"`
	Avatar           string    `gorm:"type:varchar(255)" json:"avatar"`
	Phone            string    `gorm:"type:varchar(20)" json:"phone"`
	Address          string    `gorm:"type:text" json:"address"`
	Active           bool      `gorm:"default:true" json:"active"`
	RoleID           uuid.UUID `json:"role_id" validate:"required"`
	Role             Role      `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	Orders           []Order   `gorm:"foreignKey:UserID" json:"orders,omitempty"`
	WishList         []Product `gorm:"many2many:user_wishlist;" json:"wishlist,omitempty"`
	RefreshToken     string    `gorm:"type:text" json:"-"`
	ResetToken       string    `gorm:"type:text" json:"-"`
	ResetTokenExpiry time.Time `json:"-"`
}
