package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/apis"
	"github.com/huynhtrongtien/dove/entities"
	"github.com/huynhtrongtien/dove/middlewares"
	"github.com/huynhtrongtien/dove/pkg/http/request"
	"github.com/huynhtrongtien/dove/pkg/http/response"
	"github.com/huynhtrongtien/dove/pkg/log"
)

func (h Handler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	req := &apis.CreateProductRequest{}

	// get authen user
	userID, _, _ := middlewares.ParseToken(c)
	log.For(c).Debug("[create-product] start process", log.Field("user_id", userID))

	// parse JSON
	err := request.BindJSONAndValid(c, req)
	if err != nil {
		log.For(c).Debug("[create-product] invalid request", log.Field("user_id", userID), log.Err(err))
		response.Error(c, http.StatusBadRequest, err, nil)
		return
	}

	categoryUUID := c.Param("category_uuid")
	category, err := h.Category.ReadByUUID(ctx, categoryUUID)
	if err != nil {
		log.For(c).Debug("[create-product] query category info failed", log.Field("user_id", userID), log.Field("category_uuid", categoryUUID), log.Err(err))
		response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	data := &entities.Product{
		Base: entities.Base{
			CreatedBy: userID,
		},
		CategoryID: category.ID,
		FullName:   req.Name,
		Code:       req.Code,
	}

	_, err = h.Product.Create(ctx, data)
	if err != nil {
		log.For(c).Error("[create-product] insert data failed", log.Field("user_id", userID), log.Err(err))
		response.Error(c, http.StatusInternalServerError, err, nil)
		return
	}

	log.For(c).Info("[create-product] process success", log.Field("user_id", userID), log.Field("uuid", data.UUID))
	c.JSON(http.StatusOK, &apis.CreateResponse{
		UUID: data.UUID,
	})
}
