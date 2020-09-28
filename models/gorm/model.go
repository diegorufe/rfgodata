package modelsgorm

import "time"

// BaseModelGorm for all model
type BaseModelGorm struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}
