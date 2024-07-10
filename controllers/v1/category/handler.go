package category

import "github.com/huynhtrongtien/dove/services"

type Handler struct {
	Category services.ICategory
}

func NewHandler() *Handler {
	return &Handler{
		Category: services.NewCategory(),
	}
}
