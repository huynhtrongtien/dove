package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/huynhtrongtien/dove/entities"
	"github.com/huynhtrongtien/dove/models"
)

type ICategory interface {
	Create(ctx context.Context, data *entities.Category) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Category, error)
	ReadByUUID(ctx context.Context, uuid string) (*entities.Category, error)
	ReadByName(ctx context.Context, name string) (*entities.Category, error)
	Update(ctx context.Context, data *entities.Category) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*entities.Category, error)
}

type Category struct {
	Model models.ICategory
}

func NewCategory() ICategory {
	return &Category{
		Model: models.Category{},
	}
}

func (p *Category) Create(ctx context.Context, data *entities.Category) (int64, error) {
	data.UUID = uuid.NewString()
	return p.Model.Create(ctx, data)
}

func (p *Category) Read(ctx context.Context, id int64) (*entities.Category, error) {
	return p.Model.First(ctx, map[string]any{"id": id})
}

func (p *Category) ReadByUUID(ctx context.Context, uuid string) (*entities.Category, error) {
	return p.Model.First(ctx, map[string]any{"uuid": uuid})
}

func (p *Category) ReadByName(ctx context.Context, name string) (*entities.Category, error) {
	return p.Model.First(ctx, map[string]any{"fullname": name})
}

func (p *Category) Update(ctx context.Context, data *entities.Category) error {
	return p.Model.Update(ctx, data)
}

func (p *Category) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *Category) List(ctx context.Context) ([]*entities.Category, error) {
	return p.Model.List(ctx, nil)
}
