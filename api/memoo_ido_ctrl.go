package api

import (
	"github.com/gin-gonic/gin"
	"memoo-backend/dto"
	"memoo-backend/middleware/jwt"
	"memoo-backend/serializer"
	"memoo-backend/service"
)

// @Summary web-oriented ido-upcoming
// @Description  ido
// @Accept  json
// @Produce  json
// @Param  request body  dto.PageDto  true "IdoUpcoming parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/ido-upcoming [get]
func IdoUpcoming(c *gin.Context) {
	var param dto.PageDto
	address := jwt.GetAddress(c)
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, err, "args is err")
		return
	}
	paginator, err := service.IdoUpcoming(param, address)
	serializer.WriteData2Front(c, paginator, err, "")
}

// @Summary web-oriented ido-active
// @Description  ido
// @Accept  json
// @Produce  json
// @Param  request body  dto.PageDto  true "IdoActive parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/ido-active [get]
func IdoActive(c *gin.Context) {
	var param dto.PageDto
	address := jwt.GetAddress(c)
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, err, "args is err")
		return
	}
	paginator, err := service.IdoActive(param, address)
	serializer.WriteData2Front(c, paginator, err, "")
}

// @Summary web-oriented ido-completed
// @Description  ido
// @Accept  json
// @Produce  json
// @Param  request body  dto.PageDto  true "IdoCompleted parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/ido-completed [get]
func IdoCompleted(c *gin.Context) {
	var param dto.PageDto
	address := jwt.GetAddress(c)
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, err, "args is err")
		return
	}
	paginator, err := service.IdoCompleted(param, address)
	serializer.WriteData2Front(c, paginator, err, "")
}

// @Summary web-oriented ido-upcoming-detail
// @Description  ido
// @Accept  json
// @Produce  json
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/ido-upcoming-detail [get]
func IdoUpcomingDetail(c *gin.Context) {
	address := jwt.GetAddress(c)
	resp, err := service.IdoUpcomingDetail(address)
	serializer.WriteData2Front(c, resp, err, "")
}

// @Summary web-oriented ido-active-detail
// @Description  ido
// @Accept  json
// @Produce  json
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/ido-active-detail [get]
func IdoActiveDetail(c *gin.Context) {
	address := jwt.GetAddress(c)
	resp, err := service.IdoActiveDetail(address)
	serializer.WriteData2Front(c, resp, err, "")
}

// @Summary web-oriented ido-launched-detail
// @Description  ido
// @Accept  json
// @Produce  json
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/ido-launched-detail [get]
func IdoLaunchedDetail(c *gin.Context) {
	address := jwt.GetAddress(c)
	resp, err := service.IdoLaunchedDetail(address)
	serializer.WriteData2Front(c, resp, err, "")
}

// @Summary web-oriented ido-completed
// @Description  ido
// @Accept  json
// @Produce  json
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/ido-launched-detail-top10 [get]
func IdoLaunchedDetailTop10List(c *gin.Context) {
	address := jwt.GetAddress(c)
	param := dto.PageDto{PageNumber: 1, PageSize: 10}
	paginator, err := service.IdoLaunchedDetailTop10List(param, address)
	serializer.WriteData2Front(c, paginator, err, "")
}
