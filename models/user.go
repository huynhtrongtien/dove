package models

import (
	"context"
	"fmt"

	"github.com/huynhtrongtien/dove/clients"
	"github.com/huynhtrongtien/dove/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUser interface {
	Create(ctx context.Context, data *entities.User) (int64, error)
	First(ctx context.Context, filters map[string]any) (*entities.User, error)
	Update(ctx context.Context, data *entities.User) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filter map[string]any) ([]*entities.User, error)
}

type User struct {
}

func (User) Create(ctx context.Context, data *entities.User) (int64, error) {
	err := clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Create(data).Error
	})

	return data.ID, err
}

func (User) First(ctx context.Context, filters map[string]any) (*entities.User, error) {
	result := &entities.User{}

	db := clients.MySQLClient.WithContext(ctx).Table(result.TableName())
	for field, val := range filters {
		db = db.Where(fmt.Sprintf("%s = ?", field), val)
	}

	err := db.First(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (User) Update(ctx context.Context, data *entities.User) error {
	return clients.MySQLClient.Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Model(data).Omit(clause.Associations).Save(data).Where("id = ?", data.ID).Error
	})
}

func (User) Delete(ctx context.Context, id int64) error {
	db := clients.MySQLClient
	return db.WithContext(ctx).Delete(&entities.User{}, id).Error
}

func (User) List(ctx context.Context, filters map[string]any) ([]*entities.User, error) {
	result := []*entities.User{}

	db := clients.MySQLClient.WithContext(ctx)
	for field, val := range filters {
		db = db.Where(fmt.Sprintf("%s = ?", field), val)
	}

	err := db.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
