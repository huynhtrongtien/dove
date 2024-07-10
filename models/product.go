package models

import (
	"context"
	"fmt"

	"github.com/huynhtrongtien/dove/clients"
	"github.com/huynhtrongtien/dove/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IProduct interface {
	Create(ctx context.Context, data *entities.Product) (int64, error)
	First(ctx context.Context, filters map[string]any) (*entities.Product, error)
	Update(ctx context.Context, data *entities.Product) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filter map[string]any) ([]*entities.Product, error)
}

type Product struct {
}

func (Product) Create(ctx context.Context, data *entities.Product) (int64, error) {
	err := clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Create(data).Error
	})

	return data.ID, err
}

func (Product) First(ctx context.Context, filters map[string]any) (*entities.Product, error) {
	result := &entities.Product{}

	db := clients.MySQLClient.WithContext(ctx).Table(result.TableName())
	for field, val := range filters {
		db = db.Where(fmt.Sprintf("%s = ?", field), val)
	}

	err := db.Preload("Category").First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (Product) Update(ctx context.Context, data *entities.Product) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(data).Omit(clause.Associations).Save(data).Where("id = ?", data.ID).Error
	})
}

func (Product) Delete(ctx context.Context, id int64) error {
	db := clients.MySQLClient
	return db.WithContext(ctx).Delete(&entities.Product{}, id).Error
}

func (Product) List(ctx context.Context, filters map[string]any) ([]*entities.Product, error) {
	result := []*entities.Product{}

	db := clients.MySQLClient.WithContext(ctx)
	for field, val := range filters {
		db = db.Where(fmt.Sprintf("%s = ?", field), val)
	}

	err := db.Preload("Category").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
