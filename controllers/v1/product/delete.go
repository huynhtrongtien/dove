package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/middlewares"
	"github.com/huynhtrongtien/dove/pkg/http/response"
	"github.com/huynhtrongtien/dove/pkg/log"
)

func (h Handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	userID, _, _ := middlewares.ParseToken(c)
	log.For(c).Debug("[delete-product] start process", log.Field("user_id", userID))

	uuid := c.Param("product_uuid")
	data, err := h.Product.ReadByUUID(ctx, uuid)
	if err != nil {
		log.For(c).Debug("[delete-product] query info failed", log.Field("user_id", userID), log.Field("uuid", uuid), log.Err(err))
		response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	err = h.Product.Delete(ctx, data.ID)
	if err != nil {
		log.For(c).Debug("[delete-product] execute database failed", log.Field("user_id", userID), log.Field("id", data.ID), log.Err(err))
		response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	log.For(c).Debug("[delete-product] process success", log.Field("user_id", userID), log.Field("id", data.ID))
	c.JSON(http.StatusOK, nil)
}
