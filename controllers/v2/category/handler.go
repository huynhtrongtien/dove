package category

import (
	"github.com/huynhtrongtien/dove/controllers/v1/category"
	"github.com/huynhtrongtien/dove/services"
)

type Handler struct {
	category.Handler
	Product services.IProduct
}

func NewHandler() *Handler {
	return &Handler{
		category.Handler{
			Category: services.NewCategory(),
		},
		services.NewProduct(),
	}
}
