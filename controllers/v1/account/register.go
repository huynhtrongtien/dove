package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/apis"
	"github.com/huynhtrongtien/dove/entities"
	"github.com/huynhtrongtien/dove/pkg/crypto"
	"github.com/huynhtrongtien/dove/pkg/http_parser"
	"github.com/huynhtrongtien/dove/pkg/log"
)

func (h Handler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	req := &apis.Register{}

	err := http_parser.BindAndValid(c, req)
	if err != nil {
		log.For(c).Error("[register] invalid request", log.Err(err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	_, err = h.User.ReadByUsername(ctx, req.Username)
	if err == nil {
		log.For(c).Error("[login] query database failed", log.Field("username", req.Username), log.Err(err))
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
		log.For(c).Error("[register] hash password failed", log.Field("username", req.Username), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	token, err := h.User.Create(ctx, data)
	if err != nil {
		log.For(c).Error("[register] update datebase failed", log.Field("username", req.Username), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	log.For(c).Info("[register] process success", log.Field("username", req.Username), log.Field("user_id", data.ID))
	c.JSON(http.StatusOK, &apis.AuthenticateResponse{
		Token: token,
	})
}
