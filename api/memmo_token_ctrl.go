package api

import (
	"github.com/gin-gonic/gin"
	"memoo-backend/middleware/jwt"
	"memoo-backend/oss"
	"memoo-backend/serializer"
	"memoo-backend/service"
	"mime/multipart"
)

// @Summary web-oriented token
// @Description  token
// @Accept  json
// @Produce  json
// @Param  request body  service.TokenListReqDto  true "TokenList parameters"
// @Success 200 {object} service.TokenListResp
// @Router /api/v1/web-oriented/token [get]
func TokenList(c *gin.Context) {
	var param service.TokenListReqDto
	address := jwt.GetAddress(c)
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, err, "args is err")
		return
	}
	paginator, err := service.TokenList(param, address)
	serializer.WriteData2Front(c, paginator, err, "")
}

// @Summary web-oriented token
// @Description  token
// @Accept  json
// @Produce  json
// @Param  request body  service.TokenDetailReqDto  true "TokenList parameters"
// @Success 200 {object} service.TokenDetailResp
// @Router /api/v1/web-oriented/token [get]
func TokenDetail(c *gin.Context) {
	var param service.TokenDetailReqDto
	address := jwt.GetAddress(c)
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, err, "args is err")
		return
	}
	data, err := service.TokenDetail(param, address)
	serializer.WriteData2Front(c, data, err, "")
}

// UploadFile godoc
// @Summary web-oriented token
// @Description  token
// @Accept multipart/form-data
// @Produce  json
// @Param request body service.TokenCreateOrUpdateDto true "Request body"
// @Param tokenIcon formData file true "File to upload"
// @Param banners formData file true "Files to upload"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/token [post]
func TokenNewOrEdit(c *gin.Context) {
	address := jwt.GetAddress(c)
	param, tokenIconUrls, bannerUrls, otherLinks, pinnedTwitterLinks, errMsg, err := handleTokenNewOrEditArgs(c)
	if err != nil {
		serializer.WriteData2Front(c, nil, err, errMsg)
		return
	}
	resData, err := service.TokenNewOrEdit(param, address, tokenIconUrls, bannerUrls, otherLinks, pinnedTwitterLinks, service.DRAFT)
	serializer.WriteData2Front(c, resData, err, "")
}

// UploadFile godoc
// @Summary web-oriented token
// @Description  token
// @Accept multipart/form-data
// @Produce  json
// @Param request body service.TokenCreateOrUpdateDto true "Request body"
// @Param tokenIcon formData file true "File to upload"
// @Param banners formData file true "Files to upload"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/token-confirm [post]
func TokenConfirm(c *gin.Context) {
	address := jwt.GetAddress(c)
	param, tokenIconUrls, bannerUrls, otherLinks, pinnedTwitterLinks, errMsg, err := handleTokenNewOrEditArgs(c)
	if err != nil {
		serializer.WriteData2Front(c, nil, err, errMsg)
		return
	}
	resData, err := service.TokenConfirm(param, address, tokenIconUrls, bannerUrls, otherLinks, pinnedTwitterLinks, service.QUEUE)
	serializer.WriteData2Front(c, resData, err, "")
}

func handleTokenNewOrEditArgs(c *gin.Context) (*service.TokenCreateOrUpdateDto, []string, []string, []service.LinksDto, []service.LinksDto, string, error) {
	var param *service.TokenCreateOrUpdateDto
	if err := c.ShouldBind(&param); err != nil {
		return nil, nil, nil, nil, nil, "args is err", err
	}

	tokenIconUrls, err := oss.BatchUploadFile([]*multipart.FileHeader{param.TokenIcon})
	if err != nil {
		return nil, nil, nil, nil, nil, "Failed to UploadFile", err
	}

	bannerUrls, err := oss.BatchUploadFile(param.Banners)
	if err != nil {
		return nil, nil, nil, nil, nil, "Failed to UploadFile", err
	}

	otherLinks, err := buildLinks(param.OtherLinkStr)
	if err != nil {
		return nil, nil, nil, nil, nil, "args is err", err
	}
	pinnedTwitterLinks, err := buildLinks(param.PinnedTwitterLinkStr)
	if err != nil {
		return nil, nil, nil, nil, nil, "args is err", err
	}
	return param, tokenIconUrls, bannerUrls, otherLinks, pinnedTwitterLinks, "", nil
}
