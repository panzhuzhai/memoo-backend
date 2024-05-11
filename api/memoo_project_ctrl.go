package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"memoo-backend/middleware/jwt"
	"memoo-backend/oss"
	"memoo-backend/serializer"
	"memoo-backend/service"
)

// @Summary special game
// @Description  game
// @Accept  json
// @Produce  json
// @Param  request body  dto.GameCreateDto  true "game parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/user [post]
func ProjectNewOrEdit(c *gin.Context) {
	address := jwt.GetAddress(c)
	err := c.Request.ParseMultipartForm(32 << 20) // 限制上传文件大小为32MB
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("Failed to parse form data"))
		return
	}

	form := c.Request.MultipartForm
	bannerUrls, err := oss.BatchUploadFile(form.File["profileBanner"])
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("Failed to UploadFile"))
		return
	}

	var param *service.ProjectCreateOrUpdateDto
	if err := c.ShouldBind(&param); err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	resData, err := service.ProjectNewOrEdit(param, address, bannerUrls)
	serializer.WriteData2Front(c, resData, err)
}
