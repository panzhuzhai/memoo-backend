package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"memoo-backend/dto"
	"memoo-backend/middleware/jwt"
	"memoo-backend/oss"
	"memoo-backend/serializer"
	"memoo-backend/service"
)

// @Summary web-web-oriented user-view-others
// @Description  gameS
// @Accept  json
// @Produce  json
// @Param  request body  dto.PageDto  true "games parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/user-view-others [get]
func UserViewOthers(c *gin.Context) {
	address := jwt.GetAddress(c)
	resp, err := service.UserViewOthers(address)
	serializer.WriteData2Front(c, resp, err)
}

// @Summary web-web-oriented user-view-others-list
// @Description  gameS
// @Accept  json
// @Produce  json
// @Param  request body  dto.PageDto  true "games parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-web-oriented/user-view-others-list [get]
func UserViewOthersList(c *gin.Context) {
	var param dto.PageDto
	address := jwt.GetAddress(c)
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	paginator, err := service.UserViewOthersList(param, address)
	serializer.WriteData2Front(c, paginator, err)
}

// @Summary special game
// @Description  game
// @Accept  json
// @Produce  json
// @Param  request body  dto.GameCreateDto  true "game parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/user [post]
func UserNewOrEdit(c *gin.Context) {
	address := jwt.GetAddress(c)
	// 解析表单数据
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
	profileImages, err := oss.BatchUploadFile(form.File["profileImage"])
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("Failed to UploadFile"))
		return
	}

	// 绑定JSON数据
	var param *service.UserCreateOrUpdateDto
	if err := c.ShouldBind(&param); err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	resData, err := service.UserNewOrEdit(param, address, profileImages, bannerUrls)
	serializer.WriteData2Front(c, resData, err)
}

// @Summary special game
// @Description  game
// @Accept  json
// @Produce  json
// @Param  request body  dto.GameCreateDto  true "game parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/UserNew [post]
func UserNew(c *gin.Context) {
	//address := jwt.GetAddress(c)
	var param *service.UserCreateOrUpdateDto
	err := c.BindJSON(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	//resData, err := service.UserNewOrEdit(param, address)
	//serializer.WriteData2Front(c, resData, err)
}
