package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/apis"
	"github.com/huynhtrongtien/dove/middlewares"
	"github.com/huynhtrongtien/dove/pkg/http_parser"
	"github.com/huynhtrongtien/dove/pkg/http_response"
	"github.com/huynhtrongtien/dove/pkg/log"
)

func (h Handler) SelfUpdate(c *gin.Context) {
	ctx := c.Request.Context()

	userID, _, _ := middlewares.ParseToken(c)
	log.For(c).Debug("[self-update-profile] start process", log.Field("user_id", userID))

	req := &apis.SelfUpdateProfileRequest{}
	err := http_parser.BindAndValid(c, req)
	if err != nil {
		log.For(c).Error("[self-update-profile] invalid request", log.Field("user_id", userID), log.Err(err))
		http_response.Error(c, http.StatusBadRequest, err, nil)
		return
	}

	data, err := h.User.Read(ctx, userID)
	if err != nil {
		log.For(c).Error("[self-update-profile] query database failed", log.Field("user_id", userID), log.Err(err))
		http_response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	// update data
	data.DisplayName = req.DisplayName

	// run service
	err = h.User.Update(ctx, data)
	if err != nil {
		log.For(c).Error("[self-update-profile] update database failed", log.Field("user_id", userID), log.Err(err))
		http_response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	log.For(c).Info("[self-update-profile] process success", log.Field("user_id", userID))
	c.JSON(http.StatusOK, nil)
}
