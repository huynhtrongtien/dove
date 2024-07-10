package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/apis"
	"github.com/huynhtrongtien/dove/middlewares"
	"github.com/huynhtrongtien/dove/pkg/log"
)

func (h Handler) Me(c *gin.Context) {
	ctx := c.Request.Context()

	userID, _, _ := middlewares.ParseToken(c)

	data, err := h.User.Read(ctx, userID)
	if err != nil {
		log.For(c).Error("[me] invalid request", log.Field("user_id", userID), log.Err(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	log.For(c).Info("[me] process success", log.Field("user_id", userID))
	c.JSON(http.StatusOK, &apis.SelfProfile{
		UUID:        data.UUID,
		DisplayName: data.DisplayName,
		Username:    data.Username,
	})
}
