package models

import (
	"context"
	"fmt"

	"github.com/huynhtrongtien/dove/clients"
	"github.com/huynhtrongtien/dove/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ICategory interface {
	Create(ctx context.Context, data *entities.Category) (int64, error)
	First(ctx context.Context, filters map[string]any) (*entities.Category, error)
	Update(ctx context.Context, data *entities.Category) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filter map[string]any) ([]*entities.Category, error)
}

type Category struct {
}

func (Category) Create(ctx context.Context, data *entities.Category) (int64, error) {
	err := clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Create(data).Error
	})

	return data.ID, err
}

func (Category) First(ctx context.Context, filters map[string]any) (*entities.Category, error) {
	result := &entities.Category{}

	db := clients.MySQLClient.WithContext(ctx).Table(result.TableName())
	for field, val := range filters {
		db = db.Where(fmt.Sprintf("%s = ?", field), val)
	}

	err := db.Preload("Products").First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Category) Update(ctx context.Context, data *entities.Category) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(data).Omit(clause.Associations).Save(data).Where("id = ?", data.ID).Error
	})
}

func (Category) Delete(ctx context.Context, id int64) error {
	db := clients.MySQLClient
	return db.WithContext(ctx).Delete(&entities.Category{}, id).Error
}

func (Category) List(ctx context.Context, filters map[string]any) ([]*entities.Category, error) {
	result := []*entities.Category{}

	db := clients.MySQLClient.WithContext(ctx)
	for field, val := range filters {
		db = db.Where(fmt.Sprintf("%s = ?", field), val)
	}

	err := db.Preload("Products").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
