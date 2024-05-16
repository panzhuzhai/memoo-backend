package api

import (
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
// @Param  request body  service.TokenCreateOrUpdateDto  true "TokenNewOrEdit parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/token [post]
func TokenNewOrEdit(c *gin.Context) {
	address := jwt.GetAddress(c)
	param, tokenIconUrls, bannerUrls, otherLinks, pinnedTwitterLinks, errMsg, err := handleTokenNewOrEditArgs(c)
	if err != nil {
		serializer.WriteData2Front(c, nil, err, errMsg)
		return
	}
	resData, err := service.TokenNewOrEdit(param, address, tokenIconUrls, bannerUrls, otherLinks, pinnedTwitterLinks)
	serializer.WriteData2Front(c, resData, err, "")
}

func handleTokenNewOrEditArgs(c *gin.Context) (*service.TokenCreateOrUpdateDto, []string, []string, []service.LinksDto, []service.LinksDto, string, error) {
	err := c.Request.ParseMultipartForm(32 << 20) // 限制上传文件大小为32MB
	if err != nil {
		return nil, nil, nil, nil, nil, "Failed to parse form data", err
	}

	form := c.Request.MultipartForm
	tokenIconUrls, err := oss.BatchUploadFile(form.File["tokenIcon"])
	if err != nil {
		return nil, nil, nil, nil, nil, "Failed to UploadFile", err
	}

	bannerUrls, err := oss.BatchUploadFile(form.File["banner"])
	if err != nil {
		return nil, nil, nil, nil, nil, "Failed to UploadFile", err
	}

	var param *service.TokenCreateOrUpdateDto
	if err = c.ShouldBind(&param); err != nil {
		return nil, nil, nil, nil, nil, "args is err", err
	}
	otherLinks, err := buildLinks(param.OtherLinkStr)
	if err := c.ShouldBind(&param); err != nil {
		return nil, nil, nil, nil, nil, "args is err", err
	}
	pinnedTwitterLinks, err := buildLinks(param.PinnedTwitterLinkStr)
	if err := c.ShouldBind(&param); err != nil {
		return nil, nil, nil, nil, nil, "args is err", err
	}
	if err := c.ShouldBind(&param); err != nil {
		return nil, nil, nil, nil, nil, "args is err", err
	}
	return param, tokenIconUrls, bannerUrls, otherLinks, pinnedTwitterLinks, "", nil
}
