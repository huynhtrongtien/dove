package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/apis"
	"github.com/huynhtrongtien/dove/pkg/http_parser"
	"github.com/huynhtrongtien/dove/pkg/http_response"
	"github.com/huynhtrongtien/dove/pkg/log"
)

func (h Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	req := &apis.AuthenticateRequest{}

	err := http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.For(c).Error("[login] invalid request", log.Err(err))
		http_response.Error(c, http.StatusBadRequest, err, nil)
		return
	}

	userID, token, err := h.User.Authenticate(ctx, req.Username, req.Password)
	if err != nil {
		log.For(c).Error("[login] authen token failed", log.Field("username", req.Username), log.Err(err))
		http_response.Error(c, http.StatusBadRequest, err, &http_response.Message{
			VI: "Username hoặc password không đúng",
			EN: "Username or password is incorrect",
		})
		return
	}

	log.For(c).Info("[login] process success", log.Field("user_id", userID))
	c.JSON(http.StatusOK, &apis.AuthenticateResponse{
		Token: token,
	})
}