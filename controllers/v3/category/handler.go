package category

import (
	"github.com/huynhtrongtien/dove/controllers/v1/category"
	"github.com/huynhtrongtien/dove/services"
)

type Handler struct {
	category.Handler
}

func NewHandler() *Handler {
	return &Handler{
		category.Handler{
			Category: services.NewCategory(),
		},
	}
}
