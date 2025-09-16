package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// omitempty ใช้เพื่อไม่ให้แสดงฟิลด์ที่เป็นค่า null ใน JSON

// BaseModel เป็นโครงสร้างพื้นฐานสำหรับทุกโมเดล
type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
