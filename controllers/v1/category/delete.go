package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/middlewares"
	"github.com/huynhtrongtien/dove/pkg/http_response"
	"github.com/huynhtrongtien/dove/pkg/log"
)

func (h Handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	userID, _, _ := middlewares.ParseToken(c)
	log.For(c).Debug("[delete-category] start process", log.Field("user_id", userID))

	uuid := c.Param("category_uuid")
	data, err := h.Category.ReadByUUID(ctx, uuid)
	if err != nil {
		log.For(c).Debug("[delete-category] query info failed", log.Field("user_id", userID), log.Field("uuid", uuid), log.Err(err))
		http_response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	err = h.Category.Delete(ctx, data.ID)
	if err != nil {
		log.For(c).Debug("[delete-category] execute database failed", log.Field("user_id", userID), log.Field("id", data.ID), log.Err(err))
		http_response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	log.For(c).Debug("[delete-category] process success", log.Field("user_id", userID), log.Field("id", data.ID))
	c.JSON(http.StatusOK, nil)
}
