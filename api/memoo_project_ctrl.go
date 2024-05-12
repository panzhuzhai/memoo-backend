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
	err := c.Request.ParseMultipartForm(32 << 20) // 限制上传文件大小为32MB
	if err != nil {
		serializer.WriteData2Front(c, nil, err, "Failed to parse form data")
		return
	}

	form := c.Request.MultipartForm
	bannerUrls, err := oss.BatchUploadFile(form.File["profileBanner"])
	if err != nil {
		serializer.WriteData2Front(c, nil, err, "Failed to UploadFile")
		return
	}

	var param *service.ProjectCreateOrUpdateDto
	if err := c.ShouldBind(&param); err != nil {
		serializer.WriteData2Front(c, nil, err, "args is err")
		return
	}
	otherLinkStr := param.OtherLinkStr
	var otherLinks []service.OtherLinksDto
	if otherLinkStr != "" && len(otherLinkStr) != 0 {
		err = json.Unmarshal([]byte(otherLinkStr), &otherLinks)
		if err != nil {
			serializer.WriteData2Front(c, nil, err, "args is err")
			return
		}
		param.OtherLinks = otherLinks
	}
	resData, err := service.ProjectNewOrEdit(param, address, bannerUrls)
	serializer.WriteData2Front(c, resData, err, "")
}
