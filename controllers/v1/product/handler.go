package product

import "github.com/huynhtrongtien/dove/services"

type Handler struct {
	Category services.ICategory
	Product  services.IProduct
}

func NewHandler() *Handler {
	return &Handler{
		Category: services.NewCategory(),
		Product:  services.NewProduct(),
	}
}
