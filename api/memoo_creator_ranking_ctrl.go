package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"memoo-backend/dto"
	"memoo-backend/serializer"
	"memoo-backend/service"
)

// @Summary web-unauthorized trending-creators
// @Description  creators
// @Accept  json
// @Produce  json
// @Param  request body  dto.PageDto  true "TrendingCreators parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-unauthorized/trending-creators [get]
func TrendingCreators(c *gin.Context) {
	var param dto.PageDto
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	paginator, err := service.TrendingCreators(param)
	serializer.WriteData2Front(c, paginator, err)
}

// @Summary web-unauthorized top-creators
// @Description  creators
// @Accept  json
// @Produce  json
// @Param  request body  dto.PageDto  true "TopCreators  parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-unauthorized/top-creators [get]
func TopCreators(c *gin.Context) {
	var param dto.PageDto
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	paginator, err := service.TopCreators(param)
	serializer.WriteData2Front(c, paginator, err)
}
