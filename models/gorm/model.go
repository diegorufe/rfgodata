package modelsgorm

import "time"

// BaseModelGorm for all model
type BaseModelGorm struct {
	ID        uint      `gorm:"column:primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}
