package entities

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	UUID      string         `gorm:"column:uuid;omitempty"`
	CreatedBy int64          `gorm:"column:created_by;omitempty"`
	Creator   *User          `gorm:"references:created_by;foreignKey:id;omitempty"`
	CreatedAt time.Time      `gorm:"column:created_at;omitempty"`
	UpdatedBy int64          `gorm:"column:updated_by;omitempty"`
	Updater   *User          `gorm:"references:updated_by;foreignKey:id;omitempty"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `swaggertype:"string" gorm:"index;column:deleted_at"`
}

func (b *Base) GetCreator() *User {
	if b == nil {
		return nil
	}

	return b.Creator
}

func (b *Base) GetUpdater() *User {
	if b == nil {
		return nil
	}

	return b.Updater
}
