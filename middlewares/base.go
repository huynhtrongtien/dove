package middlewares

import "github.com/huynhtrongtien/dove/models"

var userModel *models.User

func InitMiddlewares() {
	userModel = &models.User{}
}
