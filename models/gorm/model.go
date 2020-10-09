package modelsgorm

import "time"

// BaseModelGorm for all model
type BaseModelGorm struct {
	ID        uint      `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

// Tabler Interface indicates table name for table
type Tabler interface {
	TableName() string
}
