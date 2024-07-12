package product

import (
	"github.com/huynhtrongtien/dove/controllers/v1/product"
	"github.com/huynhtrongtien/dove/services"
)

type Handler struct {
	product.Handler
}

func NewHandler() *Handler {
	return &Handler{
		product.Handler{
			Category: services.NewCategory(),
			Product:  services.NewProduct(),
		},
	}
}
