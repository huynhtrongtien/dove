package apis

type Product struct {
	UUID     string    `json:"uuid,omitempty" binding:"required"`
	Name     string    `json:"fullname,omitempty" binding:"required"`
	Code     string    `json:"code,omitempty"`
	Category *Category `json:"category,omitempty"`
}

type CreateProductRequest struct {
	Name string `json:"fullname" binding:"required"`
	Code string `json:"code,omitempty"`
}

type UpdateProductRequest struct {
	Name string `json:"fullname" binding:"required"`
	Code string `json:"code,omitempty"`
}

type ListProductResponse struct {
	Data []*Product `json:"data,omitempty"`
}
