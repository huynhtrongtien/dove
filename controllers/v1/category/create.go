package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/apis"
	"github.com/huynhtrongtien/dove/entities"
	"github.com/huynhtrongtien/dove/middlewares"
	"github.com/huynhtrongtien/dove/pkg/http_parser"
	"github.com/huynhtrongtien/dove/pkg/http_response"
	"github.com/huynhtrongtien/dove/pkg/log"
)

func (h Handler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _, _ := middlewares.ParseToken(c)
	log.For(c).Debug("[create-category] start process", log.Field("user_id", userID))

	req := &apis.CreateCategoryRequest{}
	err := http_parser.BindJSONAndValid(c, req)
	if err != nil {
		log.For(c).Error("[create-category] invalid request", log.Field("user_id", userID), log.Err(err))
		http_response.Error(c, http.StatusBadRequest, err, nil)
		return
	}

	data := &entities.Category{
		FullName: req.FullName,
		Code:     req.Code,
	}

	_, err = h.Category.Create(ctx, data)
	if err != nil {
		log.For(c).Error("[create-category] query category info failed", log.Field("user_id", userID), log.Err(err))
		http_response.Error(c, http.StatusBadRequest, err, nil)
		return
	}

	log.For(c).Info("[create-category] process success", log.Field("user_id", userID), log.Field("uuid", data.UUID))
	c.JSON(http.StatusOK, &apis.CreateResponse{
		UUID: data.UUID,
	})
}
