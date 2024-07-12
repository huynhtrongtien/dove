package category

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/apis"
	"github.com/huynhtrongtien/dove/middlewares"
	"github.com/huynhtrongtien/dove/pkg/http/client"
	"github.com/huynhtrongtien/dove/pkg/http/response"
	"github.com/huynhtrongtien/dove/pkg/log"
)

func (h Handler) Read(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _, _ := middlewares.ParseToken(c)
	log.For(c).Debug("[read-category-v3] start process", log.Field("user_id", userID))

	client := client.NewDefaultHTTPClient()

	uuid := c.Param("category_uuid")

	// send request get category
	bytes, httpCode, err := client.GetWithTrace(ctx, fmt.Sprintf("http://localhost:3003/api/v3/categories/%s", uuid))
	if err != nil {
		log.For(c).Error("[read-category-v3] send request query category failed", log.Field("user_id", userID), log.Field("status_code", httpCode), log.Err(err))
		response.Error(c, httpCode, nil, nil)
		return
	}

	if httpCode != http.StatusOK {
		log.For(c).Error("[read-category-v3] send request query category failed", log.Field("user_id", userID), log.Field("status_code", httpCode), log.Err(err))
		response.Error(c, httpCode, nil, nil)
		return
	}

	catData := &apis.Category{}
	if err = json.Unmarshal(bytes, catData); err != nil {
		log.For(c).Error("[get-user] unmarshal json category failed", log.Field("user_id", userID), log.Err(err))
		response.Error(c, http.StatusInternalServerError, nil, nil)
		return
	}

	// send request get list product
	bytes, httpCode, err = client.GetWithTrace(ctx, fmt.Sprintf("http://localhost:3004/api/v3/categories/%s/products", uuid))
	if err != nil {
		log.For(c).Error("[read-category-v3] send request query list product failed", log.Field("user_id", userID), log.Field("status_code", httpCode), log.Err(err))
		response.Error(c, httpCode, nil, nil)
		return
	}

	if httpCode != http.StatusOK {
		log.For(c).Error("[read-category-v3] send request query list product failed", log.Field("user_id", userID), log.Field("status_code", httpCode), log.Err(err))
		response.Error(c, httpCode, nil, nil)
		return
	}

	listProductData := &apis.ListProductResponse{}
	if err = json.Unmarshal(bytes, listProductData); err != nil {
		log.For(c).Error("[read-category-v3] unmarshal json list product failed", log.Field("user_id", userID), log.Err(err))
		response.Error(c, http.StatusInternalServerError, nil, nil)
		return
	}

	resp := &apis.Category{
		UUID:     catData.UUID,
		FullName: catData.FullName,
		Code:     catData.Code,
	}
	for _, val := range listProductData.Data {
		resp.Products = append(resp.Products, &apis.Product{
			UUID: val.UUID,
			Name: val.Name,
			Code: val.Code,
		})
	}

	// simulate access redis to has redis and database trace
	if len(resp.Products) > 0 {
		h.Product.ReadByUUID(ctx, resp.Products[0].UUID)
		h.Product.ReadFromDB(ctx, resp.Products[0].UUID)
	}

	log.For(c).Debug("[read-category-v3] process success", log.Field("user_id", userID), log.Field("uuid", uuid))
	c.JSON(http.StatusOK, resp)
}
