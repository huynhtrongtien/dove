package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/huynhtrongtien/dove/entities"
	"github.com/huynhtrongtien/dove/models"
)

type IProduct interface {
	Create(ctx context.Context, data *entities.Product) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Product, error)
	ReadByUUID(ctx context.Context, uuid string) (*entities.Product, error)
	Update(ctx context.Context, data *entities.Product) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, categoryID int64) ([]*entities.Product, error)
}

type Product struct {
	Model models.IProduct
}

func NewProduct() IProduct {
	return &Product{
		Model: models.Product{},
	}
}

func (p *Product) Create(ctx context.Context, data *entities.Product) (int64, error) {
	data.UUID = uuid.NewString()
	return p.Model.Create(ctx, data)
}

func (p *Product) Read(ctx context.Context, id int64) (*entities.Product, error) {
	return p.Model.First(ctx, map[string]any{"id": id})
}

func (p *Product) ReadByUUID(ctx context.Context, uuid string) (*entities.Product, error) {
	return p.Model.First(ctx, map[string]any{"uuid": uuid})
}

func (p *Product) Update(ctx context.Context, data *entities.Product) error {
	return p.Model.Update(ctx, data)
}

func (p *Product) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *Product) List(ctx context.Context, categoryID int64) ([]*entities.Product, error) {
	filters := map[string]any{"category_id": categoryID}
	return p.Model.List(ctx, filters)
}
