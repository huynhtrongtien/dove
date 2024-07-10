package product

import (
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/apis"
	"github.com/huynhtrongtien/dove/middlewares"
	"github.com/huynhtrongtien/dove/pkg/http_response"
	"github.com/huynhtrongtien/dove/pkg/log"
)

func (h Handler) List(c *gin.Context) {
	ctx := c.Request.Context()

	userID, _, _ := middlewares.ParseToken(c)
	log.For(c).Debug("[list-product] start process", log.Field("user_id", userID))

	categoryUUID := c.Param("category_uuid")
	category, err := h.Category.ReadByUUID(ctx, categoryUUID)
	if err != nil {
		log.For(c).Debug("[list-product] query group info failed", log.Field("user_id", userID), log.Field("category_uuid", categoryUUID), log.Err(err))
		http_response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	data, err := h.Product.List(ctx, category.ID)
	if err != nil {
		log.For(c).Debug("[list-product] query info failed", log.Field("user_id", userID), log.Err(err))
		http_response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	resp := &apis.ListProductResponse{}
	for _, val := range data {
		resp.Data = append(resp.Data, &apis.Product{
			UUID: val.UUID,
			Name: val.FullName,
			Code: val.Code,
		})
	}

	sort.Slice(resp.Data, func(i, j int) bool {
		return strings.EqualFold(resp.Data[i].Name, resp.Data[j].Name)
	})

	log.For(c).Info("[list-product] process success", log.Field("user_id", userID), log.Field("resp", resp))
	c.JSON(http.StatusOK, resp)
}
