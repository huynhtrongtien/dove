package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/apis"
	"github.com/huynhtrongtien/dove/middlewares"
	"github.com/huynhtrongtien/dove/pkg/http_parser"
	"github.com/huynhtrongtien/dove/pkg/http_response"
	"github.com/huynhtrongtien/dove/pkg/log"
)

func (h Handler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _, _ := middlewares.ParseToken(c)
	log.For(c).Debug("[update-product] start process", log.Field("user_id", userID))

	req := &apis.UpdateProductRequest{}
	err := http_parser.BindAndValid(c, req)
	if err != nil {
		log.For(c).Debug("[update-product] invalid request", log.Field("user_id", userID), log.Err(err))
		http_response.Error(c, http.StatusBadRequest, err, nil)
		return
	}

	// get UUID form URL
	uuid := c.Param("product_uuid")
	data, err := h.Product.ReadByUUID(ctx, uuid)
	if err != nil {
		log.For(c).Debug("[update-product] query database failed", log.Field("user_id", userID), log.Field("uuid", uuid), log.Err(err))
		http_response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	data.UpdatedBy = userID
	data.FullName = req.Name
	data.Code = req.Code
	err = h.Product.Update(ctx, data)
	if err != nil {
		log.For(c).Error("[update-product] update database failed", log.Field("user_id", userID), log.Field("id", data.ID), log.Err(err))
		http_response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	log.For(c).Info("[update-product] process success", log.Field("user_id", userID), log.Field("id", data.ID))
	c.JSON(http.StatusOK, nil)
}
