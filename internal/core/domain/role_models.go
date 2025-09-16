package domain

// omitempty ใช้เพื่อไม่ให้แสดงฟิลด์ที่เป็นค่า null ใน JSON

// Role สำหรับเก็บข้อมูลสิทธิ์การใช้งาน
type Role struct {
	BaseModel
	Name        string       `gorm:"type:varchar(100);unique_index" json:"name" validate:"required"`
	Description string       `gorm:"type:text" json:"description"`
	Users       []User       `gorm:"foreignKey:RoleID" json:"users,omitempty"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}

// Permission สำหรับเก็บสิทธิ์การเข้าถึงระบบ
type Permission struct {
	BaseModel
	Name        string `gorm:"type:varchar(100);unique_index" json:"name" validate:"required"`
	Description string `gorm:"type:text" json:"description"`
	Roles       []Role `gorm:"many2many:role_permissions;" json:"roles,omitempty"`
}
