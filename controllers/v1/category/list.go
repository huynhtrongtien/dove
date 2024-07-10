package category

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
	log.For(c).Debug("[list-category] start process", log.Field("user_id", userID))

	data, err := h.Category.List(ctx)
	if err != nil {
		log.For(c).Debug("[list-category] query info failed", log.Field("user_id", userID), log.Err(err))
		http_response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	resp := &apis.ListCategoryResponse{}
	for _, val := range data {
		resp.Data = append(resp.Data, &apis.Category{
			UUID:     val.UUID,
			FullName: val.FullName,
			Code:     val.Code,
		})
	}

	sort.Slice(resp.Data, func(i, j int) bool {
		return strings.EqualFold(resp.Data[i].FullName, resp.Data[j].FullName)
	})

	log.For(c).Info("[list-category] process success", log.Field("user_id", userID), log.Field("resp", resp))
	c.JSON(http.StatusOK, resp)
}
