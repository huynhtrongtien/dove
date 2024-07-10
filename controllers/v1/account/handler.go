package account

import "github.com/huynhtrongtien/dove/services"

type Handler struct {
	User services.IUser
}

func NewHandler() *Handler {
	return &Handler{
		User: services.NewUser(),
	}
}
