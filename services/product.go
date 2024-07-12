package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/huynhtrongtien/dove/entities"
	"github.com/huynhtrongtien/dove/global"
	"github.com/huynhtrongtien/dove/models"
)

type IProduct interface {
	Create(ctx context.Context, data *entities.Product) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Product, error)
	ReadByUUID(ctx context.Context, uuid string) (*entities.Product, error)
	ReadFromDB(ctx context.Context, uuid string) (*entities.Product, error)
	Update(ctx context.Context, data *entities.Product) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, categoryID int64) ([]*entities.Product, error)
}

type Product struct {
	Model models.IProduct
	Cache models.ICachedProduct
}

func NewProduct() IProduct {
	return &Product{
		Model: models.Product{},
		Cache: models.CachedProduct{
			Prefix:     fmt.Sprintf("%s:%s", global.Environment(), global.ServiceName()),
			MaxRetry:   2,
			Expiration: time.Millisecond * 1 * 60 * 60 * 1000, // 1 hour
		},
	}
}

func (p *Product) Create(ctx context.Context, data *entities.Product) (int64, error) {
	data.UUID = uuid.NewString()
	result, err := p.Model.Create(ctx, data)
	if err != nil {
		return 0, err
	}

	// set cache to reuse in next time
	p.Cache.Set(ctx, data)
	return result, nil
}

func (p *Product) Read(ctx context.Context, id int64) (*entities.Product, error) {
	// try to get from cache
	result, err := p.Cache.Get(ctx, id)
	if err == nil {
		return result, nil
	}

	// query database
	result, err = p.Model.First(ctx, map[string]any{"id": id})
	if err != nil {
		return nil, err
	}

	// set cache to reuse in next time
	p.Cache.Set(ctx, result)

	return result, nil
}

func (p *Product) ReadByUUID(ctx context.Context, uuid string) (*entities.Product, error) {
	// try to get from cache
	result, err := p.Cache.GetByUUID(ctx, uuid)
	if err == nil {
		return result, nil
	}

	// query database
	result, err = p.Model.First(ctx, map[string]any{"uuid": uuid})
	if err != nil {
		return nil, err
	}

	// set cache to reuse in next time
	p.Cache.Set(ctx, result)

	return result, nil
}

func (p *Product) ReadFromDB(ctx context.Context, uuid string) (*entities.Product, error) {
	return p.Model.First(ctx, map[string]any{"uuid": uuid})
}

func (p *Product) Update(ctx context.Context, data *entities.Product) error {
	if err := p.Model.Update(ctx, data); err != nil {
		return err
	}

	// update cache
	if err := p.Cache.Set(ctx, data); err != nil {
		p.Delete(ctx, data.ID)
	}

	return nil
}

func (p *Product) Delete(ctx context.Context, id int64) error {
	err := p.Model.Delete(ctx, id)
	if err != nil {
		return err
	}

	// clear cache
	return p.Cache.Delete(ctx, id)
}

func (p *Product) List(ctx context.Context, categoryID int64) ([]*entities.Product, error) {
	filters := map[string]any{"category_id": categoryID}
	return p.Model.List(ctx, filters)
}
