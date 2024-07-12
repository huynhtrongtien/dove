package category

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/apis"
	"github.com/huynhtrongtien/dove/middlewares"
	"github.com/huynhtrongtien/dove/pkg/http/request"
	"github.com/huynhtrongtien/dove/pkg/http/response"
	"github.com/huynhtrongtien/dove/pkg/log"
	"gorm.io/gorm"
)

func (h Handler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _, _ := middlewares.ParseToken(c)
	log.For(c).Debug("[update-category] start process", log.Field("user_id", userID))

	req := &apis.UpdateCategoryRequest{}
	err := request.BindAndValid(c, req)
	if err != nil {
		log.For(c).Debug("[update-category] invalid request", log.Field("user_id", userID), log.Err(err))
		response.Error(c, http.StatusBadRequest, err, nil)
		return
	}

	// get UUID form URL
	uuid := c.Param("category_uuid")
	data, err := h.Category.ReadByUUID(ctx, uuid)
	if err != nil {
		log.For(c).Debug("[update-category] query database failed", log.Field("user_id", userID), log.Field("uuid", uuid), log.Err(err))
		response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	// check name is duplicated
	if data.FullName != req.FullName {
		_, err = h.Category.ReadByName(ctx, req.FullName)
		if err == nil {
			log.For(c).Debug("[update-category] query database failed", log.Field("user_id", userID), log.Err(err))
			response.Error(c, http.StatusConflict, nil, &response.Message{
				VI: "Name is exist",
			})
			return

		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.For(c).Debug("[update-category] query database failed", log.Field("user_id", userID), log.Err(err))
			response.Error(c, http.StatusInternalServerError, err, nil)
			return
		}
	}

	data.UpdatedBy = userID
	data.FullName = req.FullName
	data.Code = req.Code
	err = h.Category.Update(ctx, data)
	if err != nil {
		log.For(c).Error("[update-category] update database failed", log.Field("user_id", userID), log.Field("id", data.ID), log.Err(err))
		response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	log.For(c).Info("[update-category] process success", log.Field("user_id", userID), log.Field("id", data.ID))
	c.JSON(http.StatusOK, nil)
}
