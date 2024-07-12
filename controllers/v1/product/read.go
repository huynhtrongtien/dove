package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/apis"
	"github.com/huynhtrongtien/dove/middlewares"
	"github.com/huynhtrongtien/dove/pkg/http/response"
	"github.com/huynhtrongtien/dove/pkg/log"
)

func (h Handler) Read(c *gin.Context) {
	ctx := c.Request.Context()

	userID, _, _ := middlewares.ParseToken(c)
	log.For(c).Debug("[read-product] start process", log.Field("user_id", userID))

	uuid := c.Param("product_uuid")
	data, err := h.Product.ReadByUUID(ctx, uuid)
	if err != nil {
		log.For(c).Debug("[read-product] query database failed", log.Field("user_id", userID), log.Field("uuid", uuid))
		response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	log.For(c).Info("[read-product] process success", log.Field("user_id", userID), log.Field("resp", data))
	c.JSON(http.StatusOK, &apis.Product{
		UUID: data.UUID,
		Name: data.FullName,
		Code: data.Code,
	})
}
