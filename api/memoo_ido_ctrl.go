package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"memoo-backend/dto"
	"memoo-backend/middleware/jwt"
	"memoo-backend/serializer"
	"memoo-backend/service"
)

// @Summary web-oriented ido-upcoming
// @Description  gameS
// @Accept  json
// @Produce  json
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/ido-upcoming [get]
func IdoUpcoming(c *gin.Context) {
	var param dto.PageDto
	address := jwt.GetAddress(c)
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	paginator, err := service.IdoUpcoming(param, address)
	serializer.WriteData2Front(c, paginator, err)
}

// @Summary web-oriented ido-active
// @Description  gameS
// @Accept  json
// @Produce  json
// @Param  request body  dto.PageDto  true "games parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/ido-active [get]
func IdoActive(c *gin.Context) {
	var param dto.PageDto
	address := jwt.GetAddress(c)
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	paginator, err := service.IdoActive(param, address)
	serializer.WriteData2Front(c, paginator, err)
}

// @Summary web-oriented ido-completed
// @Description  gameS
// @Accept  json
// @Produce  json
// @Param  request body  dto.PageDto  true "games parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/ido-completed [get]
func IdoCompleted(c *gin.Context) {
	var param dto.PageDto
	address := jwt.GetAddress(c)
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	paginator, err := service.IdoCompleted(param, address)
	serializer.WriteData2Front(c, paginator, err)
}

// @Summary web-oriented ido-upcoming-detail
// @Description  gameS
// @Accept  json
// @Produce  json
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/ido-upcoming-detail [get]
func IdoUpcomingDetail(c *gin.Context) {
	var param dto.PageDto
	address := jwt.GetAddress(c)
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	paginator, err := service.IdoUpcomingDetail(param, address)
	serializer.WriteData2Front(c, paginator, err)
}

// @Summary web-oriented ido-active-detail
// @Description  gameS
// @Accept  json
// @Produce  json
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/ido-active-detail [get]
func IdoActiveDetail(c *gin.Context) {
	var param dto.PageDto
	address := jwt.GetAddress(c)
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	paginator, err := service.IdoActiveDetail(param, address)
	serializer.WriteData2Front(c, paginator, err)
}

// @Summary web-oriented ido-launched-detail
// @Description  gameS
// @Accept  json
// @Produce  json
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/ido-launched-detail [get]
func IdoLaunchedDetail(c *gin.Context) {
	var param dto.PageDto
	address := jwt.GetAddress(c)
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, errors.New("args is err"))
		return
	}
	paginator, err := service.IdoLaunchedDetail(param, address)
	serializer.WriteData2Front(c, paginator, err)
}

// @Summary web-oriented ido-completed
// @Description  gameS
// @Accept  json
// @Produce  json
// @Param  request body  dto.PageDto  true "games parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/ido-launched-detail-top10 [get]
func IdoLaunchedDetailTop10List(c *gin.Context) {
	address := jwt.GetAddress(c)
	param := dto.PageDto{PageNumber: 1, PageSize: 10}
	paginator, err := service.IdoLaunchedDetailTop10List(param, address)
	serializer.WriteData2Front(c, paginator, err)
}
