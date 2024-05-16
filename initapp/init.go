package initapp

import (
	"github.com/gin-gonic/gin"
	"memoo-backend/localconfig"
	"memoo-backend/logger"
	"memoo-backend/middleware/database"
	"memoo-backend/oss"
)

func Init() {

	// 读取翻译文件
	if err := localconfig.LoadLocales("localconfig/locales/zh-cn.yaml"); err != nil {
		//util.Log().Panic("翻译文件加载失败", err)
	}

	localconfig.LoadAppConfig()
	gin.SetMode(localconfig.AppConfig.Extension.GinMode)

	logger.InitLogger(gin.Mode())
	localconfig.InitBtcNetwork(gin.Mode())

	database.InitDatabase()
	oss.InitAws()
}
