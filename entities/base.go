package entities

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	CreatedBy int64          `gorm:"column:created_by;omitempty" redis:"-"`
	Creator   *User          `gorm:"references:created_by;foreignKey:id;omitempty" redis:"-"`
	CreatedAt time.Time      `gorm:"column:created_at;omitempty" redis:"-"`
	UpdatedBy int64          `gorm:"column:updated_by;omitempty" redis:"-"`
	Updater   *User          `gorm:"references:updated_by;foreignKey:id;omitempty" redis:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at" redis:"-"`
	DeletedAt gorm.DeletedAt `swaggertype:"string" gorm:"index;column:deleted_at" redis:"-"`
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
