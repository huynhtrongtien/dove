package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/apis"
	"github.com/huynhtrongtien/dove/entities"
	"github.com/huynhtrongtien/dove/pkg/crypto"
	"github.com/huynhtrongtien/dove/pkg/http_parser"
)

func (h Handler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	req := &apis.Register{}

	err := http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	_, err = h.User.ReadByUsername(ctx, req.Username)
	if err == nil {
		c.JSON(http.StatusConflict, &apis.Error{
			Message: &apis.ErrorMessage{
				VI: "User name đã tồn tại",
				EN: "User name is exist",
			},
		})
		return
	}

	data := &entities.User{
		DisplayName: req.DisplayName,
		Username:    req.Username,
	}

	data.PasswordHash, err = crypto.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	token, err := h.User.Create(ctx, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &apis.AuthenticateResponse{
		Token: token,
	})
}
