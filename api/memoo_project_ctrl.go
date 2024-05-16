package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"memoo-backend/middleware/jwt"
	"memoo-backend/oss"
	"memoo-backend/serializer"
	"memoo-backend/service"
)

// @Summary web-oriented project
// @Description  project
// @Accept  json
// @Produce  json
// @Param  request body  service.ProjectCreateOrUpdateDto  true "ProjectNewOrEdit parameters"
// @Success 200 {object} serializer.Response
// @Router /api/v1/web-oriented/project [post]
func ProjectNewOrEdit(c *gin.Context) {
	address := jwt.GetAddress(c)
	param, bannerUrls, otherLinks, pinnedTwitterLinks, errMsg, err := handleProjectNewOrEditArgs(c)
	if err != nil {
		serializer.WriteData2Front(c, nil, err, errMsg)
		return
	}
	resData, err := service.ProjectNewOrEdit(param, address, bannerUrls, otherLinks, pinnedTwitterLinks)
	serializer.WriteData2Front(c, resData, err, "")
}

func handleProjectNewOrEditArgs(c *gin.Context) (*service.ProjectCreateOrUpdateDto, []string, []service.LinksDto, []service.LinksDto, string, error) {
	err := c.Request.ParseMultipartForm(32 << 20) // 限制上传文件大小为32MB
	if err != nil {
		return nil, nil, nil, nil, "Failed to parse form data", err
	}
	form := c.Request.MultipartForm
	bannerUrls, err := oss.BatchUploadFile(form.File["banner"])
	if err != nil {
		return nil, nil, nil, nil, "Failed to UploadFile", err
	}
	var param *service.ProjectCreateOrUpdateDto
	if err := c.ShouldBind(&param); err != nil {
		return nil, nil, nil, nil, "args is err", err
	}
	otherLinks, err := buildLinks(param.OtherLinkStr)
	if err := c.ShouldBind(&param); err != nil {
		return nil, nil, nil, nil, "args is err", err
	}
	pinnedTwitterLinks, err := buildLinks(param.PinnedTwitterLinkStr)
	if err := c.ShouldBind(&param); err != nil {
		return nil, nil, nil, nil, "args is err", err
	}
	if err := c.ShouldBind(&param); err != nil {
		return nil, nil, nil, nil, "args is err", err
	}
	return param, bannerUrls, otherLinks, pinnedTwitterLinks, "", nil
}

func buildLinks(linkStr string) ([]service.LinksDto, error) {
	var links []service.LinksDto
	if linkStr == "" || len(linkStr) == 0 {
		return nil, nil
	}
	err := json.Unmarshal([]byte(linkStr), &links)
	return links, err
}
