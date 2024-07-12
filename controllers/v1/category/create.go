package category

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/apis"
	"github.com/huynhtrongtien/dove/entities"
	"github.com/huynhtrongtien/dove/middlewares"
	"github.com/huynhtrongtien/dove/pkg/http/request"
	"github.com/huynhtrongtien/dove/pkg/http/response"
	"github.com/huynhtrongtien/dove/pkg/log"
	"gorm.io/gorm"
)

func (h Handler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _, _ := middlewares.ParseToken(c)
	log.For(c).Debug("[create-category] start process", log.Field("user_id", userID))

	req := &apis.CreateCategoryRequest{}
	err := request.BindJSONAndValid(c, req)
	if err != nil {
		log.For(c).Error("[create-category] invalid request", log.Field("user_id", userID), log.Err(err))
		response.Error(c, http.StatusBadRequest, err, nil)
		return
	}

	// check name exist
	existData, err := h.Category.ReadByName(ctx, req.FullName)
	if err == nil {
		log.For(c).Debug("[update-category] category name is exist", log.Field("user_id", userID), log.Field("uuid", existData.UUID))
		c.JSON(http.StatusCreated, &apis.CreateResponse{
			UUID: existData.UUID,
		})
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) == false {
		log.For(c).Debug("[update-category] query database failed", log.Field("user_id", userID), log.Err(err))
		response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	data := &entities.Category{
		Base: entities.Base{
			CreatedBy: userID,
		},
		FullName: req.FullName,
		Code:     req.Code,
	}

	_, err = h.Category.Create(ctx, data)
	if err != nil {
		log.For(c).Error("[create-category] query category info failed", log.Field("user_id", userID), log.Err(err))
		response.Error(c, http.StatusBadRequest, err, nil)
		return
	}

	log.For(c).Info("[create-category] process success", log.Field("user_id", userID), log.Field("uuid", data.UUID))
	c.JSON(http.StatusOK, &apis.CreateResponse{
		UUID: data.UUID,
	})
}
