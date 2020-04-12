package model

type BaseModel struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT;not null"`
	CreatedAt int64 `gorm:"cloumn:created_at"`
	UpdatedAt int64 `gorm:"cloumn:updated_at"`
}
