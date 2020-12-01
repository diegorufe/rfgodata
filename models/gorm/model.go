package modelsgorm

import "database/sql"

// BaseModelGorm for all model
type BaseModelGorm struct {
	Id        uint         `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt sql.NullTime `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `gorm:"column:updatedAt" json:"updatedAt"`
}

// Tabler Interface indicates table name for table
type Tabler interface {
	TableName() string
}
