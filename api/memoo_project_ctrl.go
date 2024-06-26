package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"memoo-backend/middleware/jwt"
	"memoo-backend/oss"
	"memoo-backend/serializer"
	"memoo-backend/service"
)

func ProjectNewOrEdit(c *gin.Context) {
	address := jwt.GetAddress(c)
	param, bannerUrls, errMsg, err := handleProjectNewOrEditArgs(c)
	if err != nil {
		serializer.WriteData2Front(c, nil, err, errMsg)
		return
	}
	resData, err := service.ProjectNewOrEdit(param, address, bannerUrls)
	serializer.WriteData2Front(c, resData, err, "")
}

func handleProjectNewOrEditArgs(c *gin.Context) (*service.ProjectCreateOrUpdateDto, []string, string, error) {
	err := c.Request.ParseMultipartForm(32 << 20) // 限制上传文件大小为32MB
	var param *service.ProjectCreateOrUpdateDto
	if err := c.ShouldBind(&param); err != nil {
		return nil, nil, "args is err", err
	}
	bannerUrls, err := oss.BatchUploadFile(param.Banners)
	if err != nil {
		return nil, nil, "Failed to UploadFile", err
	}
	return param, bannerUrls, "", nil
}

func buildLinks(linkStr string) ([]service.LinksDto, error) {
	var links []service.LinksDto
	if linkStr == "" || len(linkStr) == 0 {
		return nil, nil
	}
	err := json.Unmarshal([]byte(linkStr), &links)
	return links, err
}
