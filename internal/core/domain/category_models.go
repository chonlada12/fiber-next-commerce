package domain

// Category สำหรับเก็บข้อมูลหมวดหมู่สินค้า
type Category struct {
	BaseModel
	Name        string    `gorm:"type:varchar(100);unique_index" json:"name" validate:"required"`
	Description string    `gorm:"type:text" json:"description"`
	Image       string    `gorm:"type:varchar(255)" json:"image"`
	Products    []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}
