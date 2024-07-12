package apis

type Category struct {
	UUID     string     `json:"uuid,omitempty"`
	FullName string     `json:"fullname,omitempty"`
	Code     string     `json:"code,omitempty"`
	Products []*Product `json:"products,omitempty"`
}

type CreateCategoryRequest struct {
	FullName string `json:"fullname" binding:"required" valid:"MaxSize(100)"`
	Code     string `json:"code,omitempty"`
}

type UpdateCategoryRequest struct {
	FullName string `json:"fullname,omitempty" binding:"required" valid:"MaxSize(100)"`
	Code     string `json:"code,omitempty"`
}

type ListCategoryResponse struct {
	Data []*Category `json:"data,omitempty"`
}
