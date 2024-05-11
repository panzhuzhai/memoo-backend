package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"memoo-backend/dto"
	"memoo-backend/serializer"
	"memoo-backend/service"
)

// @Summary web-unauthorized trending-tokens
// @Description  Tokens
// @Accept  json
// @Produce  json
// @Param  request body  dto.PageDto  true "TrendingTokens parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-unauthorized/trending-tokens [get]
func TrendingTokens(c *gin.Context) {
	var param dto.PageDto
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	paginator, err := service.TrendingTokens(param)
	serializer.WriteData2Front(c, paginator, err)
}

// @Summary web-unauthorized top-tokens
// @Description  Tokens
// @Accept  json
// @Produce  json
// @Param  request body  dto.PageDto  true "TopTokens parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-unauthorized/top-tokens [get]
func TopTokens(c *gin.Context) {
	var param dto.PageDto
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	paginator, err := service.TopTokens(param)
	serializer.WriteData2Front(c, paginator, err)
}
