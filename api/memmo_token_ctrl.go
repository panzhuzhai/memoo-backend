package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"memoo-backend/middleware/jwt"
	"memoo-backend/oss"
	"memoo-backend/serializer"
	"memoo-backend/service"
)

// @Summary web-oriented token
// @Description  token
// @Accept  json
// @Produce  json
// @Param  request body  service.TokenListReqDto  true "TokenList parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/token [get]
func TokenList(c *gin.Context) {
	var param service.TokenListReqDto
	address := jwt.GetAddress(c)
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	paginator, err := service.TokenList(param, address)
	serializer.WriteData2Front(c, paginator, err)
}

// @Summary web-oriented token
// @Description  token
// @Accept  json
// @Produce  json
// @Param  request body  service.TokenCreateOrUpdateDto  true "TokenNewOrEdit parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/token [post]
func TokenNewOrEdit(c *gin.Context) {
	address := jwt.GetAddress(c)
	err := c.Request.ParseMultipartForm(32 << 20) // 限制上传文件大小为32MB
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("Failed to parse form data"))
		return
	}

	form := c.Request.MultipartForm
	tokenIconUrls, err := oss.BatchUploadFile(form.File["tokenIcon"])
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("Failed to UploadFile"))
		return
	}

	bannerUrls, err := oss.BatchUploadFile(form.File["profileBanner"])
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("Failed to UploadFile"))
		return
	}

	var param *service.TokenCreateOrUpdateDto
	if err := c.ShouldBind(&param); err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	resData, err := service.TokenNewOrEdit(param, address, tokenIconUrls, bannerUrls)
	serializer.WriteData2Front(c, resData, err)
}
